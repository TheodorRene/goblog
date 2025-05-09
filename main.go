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

	"github.com/yuin/goldmark"
)

func check(e error) {
	if e != nil {
		fmt.Println("panic?")
		panic(e)
	}
}

func main() {
	recreateBuildFolder()
	addStyleSheets()
	fmt.Println("Stylesheet(s) added")
	buildFrontPage()
	fmt.Println("Frontpage added")
	buildBlogPosts()
	fmt.Println("Build done")
}
func recreateBuildFolder() {
	err := os.RemoveAll("build")
	check(err)
	os.Mkdir("build", 0755)
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
	for _, blogPost := range dir {

		nameWithoutExt := blogPost.Name()[:len(blogPost.Name())-3]
		data, err := os.ReadFile("posts/" + blogPost.Name())
		check(err)
		htmlString := mdToHtml(data).String()
		var buf bytes.Buffer
		tmplate.Execute(&buf, template.HTML(htmlString))
		os.WriteFile("build/"+nameWithoutExt+".html", buf.Bytes(), 0755)
	}
	fmt.Println("Posts added")
}

type PostMetadata struct {
	title       string
	tags        []string
	draft       bool
	description string
	date        string // for now
}

func parseMetadata(post []byte) *PostMetadata {
	scanner := bufio.NewScanner(strings.NewReader(string(post)))
	var lines []string
	var metaData PostMetadata
	for scanner.Scan() {
		line := scanner.Text()
		if line == "+++" {
			for scanner.Scan() {
				if line == "+++" {
					break
				}
				lines = append(lines, line)
				// use regex to get each side
			}
			break
		}
	}


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
