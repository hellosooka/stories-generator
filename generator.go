package main

import (
	"flag"
	vueStoryGenerator "github.com/hellosooka/stories-generator/src/vue/story"
)

func main() {
	template := flag.String("t", "./templates/", "")
	fileTree := flag.String("d", "./testFileTree/", "")
	section := flag.String("s", "story", "")
	isVue := flag.Bool("vue", false, "")

	flag.Parse()

	if *isVue {
		vueStoryGenerator.CreateVueStories(*fileTree, *template, *section)
	}

}

// Vue/utils
