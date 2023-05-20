package main

import vue_parser "github.com/hellosooka/stories-generator/src/vue"

type StoryItem struct {
	Filename string
	Section  string
}

func main() {
	vue_parser.CreateVueStories("./testFileTree/", "./templates/")
}

// Vue/utils
