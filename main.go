package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/yuin/goldmark"
	"html/template"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		fmt.Println("panic?")
		panic(e)
	}
}

func main() {
	recreateBuildFolder()
	buildFrontPage()
	buildBlogPosts()
}
func recreateBuildFolder() {
	err := os.RemoveAll("build")
	check(err)
	os.Mkdir("build", 0755)
}

func buildFrontPage() {
	t, err := getTemplate("index", "templates/index.template.html")
	check(err)
	var buf bytes.Buffer
	postNames := getPostNames()
	t.Execute(&buf, template.HTML(generateLinkList(postNames)))
	os.WriteFile("build/index.html", buf.Bytes(), 0755)
	fmt.Println("build done")
}

func generateLinkList(postNames []string) (res string) {
	res += "<ul>\n"
	for _, name := range postNames {
		res += fmt.Sprintf("<li><a href='/%s.html'> %s </a></li>\n", name, name)
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

func buildBlogPosts() {
	dir, err := os.ReadDir("posts")
	tmplate, err := getTemplate("post", "templates/post.template.html")
	check(err)
	for _, blogPost := range dir {
		data, err := os.ReadFile("posts/" + blogPost.Name())
		check(err)
		htmlString := mdToHtml(data).String()
		var buf bytes.Buffer
		tmplate.Execute(&buf, template.HTML(htmlString))
		os.WriteFile("build/"+blogPost.Name()+".html", buf.Bytes(), 0755)
	}
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
