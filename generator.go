package main

import (
	"flag"
	vue_parser "github.com/hellosooka/stories-generator/src/vue"
)

func main() {
	template := flag.String("t", "./templates/", "")
	fileTree := flag.String("d", "./testFileTree/", "")
	section := flag.String("s", "story", "")
	isVue := flag.Bool("vue", false, "")

	flag.Parse()

	if *isVue {
		vue_parser.CreateVueStories(*fileTree, *template, *section)
	}

}

// Vue/utils
