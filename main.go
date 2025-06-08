package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/yuin/goldmark"
)

func check(e error) {
	if e != nil {
		fmt.Println("panic?")
		panic(e)
	}
}

var wg sync.WaitGroup
var posts []*Post

func main() {
	timeIt("Everything took", func() {
		timeIt("build folder", recreateBuildFolder)
		wg.Add(3)
		go timeItDone("stylesheets", addStyleSheets)
		go timeItDone("frontpage", buildFrontPage)
		go timeItDone("blogpost", buildBlogPosts)
		wg.Wait()
	})
}
func recreateBuildFolder() {
	err := os.RemoveAll("build")
	check(err)
	os.Mkdir("build", 0755)
}

func timeIt(name string, fn func()) {
	start := time.Now()
	fn()
	fmt.Printf("%s took %s\n", name, time.Since(start))
}

func timeItDone(name string, fn func()) {
	start := time.Now()
	fn()
	wg.Done()
	fmt.Printf("%s took %s\n", name, time.Since(start))
}
func addStyleSheets() {
	dat, err := os.ReadFile("templates/styles.css")
	check(err)
	err = os.WriteFile("build/styles.css", dat, 0755)
	check(err)
}

func buildFrontPage() {
	t, err := getTemplate("index", "templates/index.template.html")
	check(err)
	var buf bytes.Buffer
	postNames := getPostNames()
	t.Execute(&buf, template.HTML(generateLinkList(postNames)))
	os.WriteFile("build/index.html", buf.Bytes(), 0755)
}

func buildBlogPosts() {
	dir, err := os.ReadDir("posts")
	tmplate, err := getTemplate("post", "templates/post.template.html")
	check(err)
	var wgBg sync.WaitGroup
	for _, blogPost := range dir {
		wgBg.Add(1)
		go func(bspost os.DirEntry) {
			defer wgBg.Done()
			nameWithoutExt := blogPost.Name()[:len(blogPost.Name())-3]
			data, err := os.ReadFile("posts/" + blogPost.Name())
			check(err)
			htmlBuf := mdToHtml(data).Bytes()
			metaData, err := parseMetadata(data)
			check(err)
			var buf bytes.Buffer
			tmplate.Execute(&buf, template.HTML(htmlBuf))
			post := Post{
				MetaData:   metaData,
				postBuffer: &htmlBuf,
			}
			if !post.MetaData.Draft {
				os.WriteFile("build/"+nameWithoutExt+".html", buf.Bytes(), 0755)
			}
		}(blogPost)
	}
	wgBg.Wait()
}

type Post struct {
	MetaData   *PostMetadata
	postBuffer *[]byte
}

type PostMetadata struct {
	Title       string
	Tags        []string
	Draft       bool
	Description string
	Date        string // for now
}

func parseMetadata(post []byte) (*PostMetadata, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(post)))
	var lines []string
	var metaData PostMetadata
	for scanner.Scan() {
		line := scanner.Text()
		if line == "+++" {
			for scanner.Scan() {
				line := scanner.Text()
				if line == "+++" {
					break
				}
				lines = append(lines, line)
				// use regex to get each side
				re := regexp.MustCompile(`(\w+) ?=? ?(.*)`)
				match := re.FindStringSubmatch(line)
				key := match[1]
				value := match[2]
				switch key {
				case "title":
					metaData.Title = strings.Trim(value, `" `)
				case "tags":
					trimmed := strings.Trim(value, "[]")
					parts := strings.Split(trimmed, ",")
					for i := range parts {
						parts[i] = strings.Trim(parts[i], `" `)
					}
					metaData.Tags = parts
				case "draft":
					myBool := true
					if value == "false" {
						myBool = false
					}
					if value != "true" && value != "false" {
						fmt.Printf("Something else? %s\n", value)
					}
					metaData.Draft = myBool
				case "description":
					metaData.Description = strings.Trim(value, `" `)
				case "date":
					metaData.Date = strings.Trim(value, `" `)

				}
				// fmt.Printf("Key: %s, Value: %s\n", key, value)
			}
			break
		}
	}
	return &metaData, nil
}

func generateLinkList(postNames []string) (res string) {
	res += "<ul>\n"
	for _, name := range postNames {
		name = name[:len(name)-3] // Remove the .md extension
		escapedName := template.HTMLEscapeString(name)
		res += fmt.Sprintf("<li><a href='/%s.html'> %s </a></li>\n", escapedName, escapedName)
	}
	res += "</ul>"
	return
}

func getPostNames() (postNames []string) {
	dir, _ := os.ReadDir("posts")
	for _, post := range dir {
		postNames = append(postNames, post.Name())
	}
	return
}

func mdToHtml(markdownBuffer []byte) *bytes.Buffer {
	var buf bytes.Buffer
	md := goldmark.New()
	err := md.Convert(markdownBuffer, &buf)
	check(err)
	return &buf
}

// func getTemplatePath(name string) string {
// 	return "templates/" + name + ".template.html"
// }

func getTemplateName(path string) (name *string, err error) {
	re := regexp.MustCompile(`(?:^|/)(\w+)\.template\.`)
	match := re.FindStringSubmatch(path)
	if match != nil && len(match) > 1 {
		return &match[1], nil
	}
	return nil, errors.New("No name found")
}

func getTemplate(name string, path string) (*template.Template, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("Could not read file")
	}
	return template.New(name).Parse(string(dat))
}
