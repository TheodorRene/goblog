package main

import (
	"testing"
	"reflect"
)

func TestGenerateLinkList(t *testing.T) {
	postNames := []string{"post1.md", "post2.md", "post3.md"}
	expected := "<ul>\n<li><a href='/post1.html'> post1 </a></li>\n<li><a href='/post2.html'> post2 </a></li>\n<li><a href='/post3.html'> post3 </a></li>\n</ul>"
	result := generateLinkList(postNames)
	if result != expected {
		t.Errorf("generateLinkList() = %v, want %v", result, expected)
	}
}

func TestParseMetaData(t *testing.T) {
	metaData := `+++
title = "Config_rant"
date = "2023-05-15T10:11:12+02:00"
author = ""
authorTwitter = "" #do not include @
cover = ""
tags = ["", ""]
keywords = ["", "hei"]
description = "hei"
showFullContent = false
readingTime = false
draft = true
+++
disregard = this
`
	metaDataAsBytes := []byte(metaData)
	expected := PostMetadata {
		Title: "Config_rant",
		Tags: []string{"", ""},
		Draft: true,
		Description: "hei",
		Date: "2023-05-15T10:11:12+02:00",
	}
	result, err := parseMetadata(metaDataAsBytes)
	if err != nil {
		t.Errorf("parseMetadata failed")
	}
	if !reflect.DeepEqual(*result, expected){
		t.Errorf("generateLinkList() = %v, want %v", result, expected)
	}

}


