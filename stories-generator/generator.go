package main

import (
	"flag"

	vue_parser "github.com/hellosooka/stories-generator/src/vue"
)

func main() {
	template := flag.String("t", "./templates/", "")
	fileTree := flag.String("d", "./testFileTree/", "")

	flag.Parse()

	vue_parser.CreateVueStories(*fileTree, *template)

}

// Vue/utils
