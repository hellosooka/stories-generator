package vueStoryGenerator

import (
	"fmt"
	"log"
	"os"
	"sync"
	"text/template"

	"github.com/hellosooka/stories-generator/src/constants"
	"github.com/hellosooka/stories-generator/src/stories"
	"github.com/hellosooka/stories-generator/src/utils"
	"github.com/hellosooka/stories-generator/src/vue/parser"
)

var temp *template.Template

const vueTemplate string = constants.VUE_TEMPLATE_FILENAME + constants.TEMPLATE_EXTEND

func CreateVueStories(path string, templatePath string, section string) {
	var wg sync.WaitGroup
	temp = utils.GetFilteredTemplates(templatePath, vueTemplate)
	files := stories.GetStoryItems(vueParser.ParseVueFilesPath(path), section)

	wg.Add(len(files))
	for i := 0; i < len(files); i++ {
		go func(i int) {
			createStory(files[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func createStory(story stories.StoryItem) {
	workDirectory := fmt.Sprintf("%s%s", story.Directory, createStoryFilename(story.Filename))
	_, err := os.Open(workDirectory)
	if err != nil {
		f, err := os.Create(workDirectory)
		defer f.Close()

		story.Props = vueParser.GetProps(story)
		if err != nil {
			log.Fatal(err)
		}
		temp.Execute(f, story)
	}
}

func createStoryFilename(filename string) string {
	return fmt.Sprintf("%s%s.ts", filename, constants.STORY_POSTFIX)
}
