package vue_parser

import (
	"fmt"
	"github.com/hellosooka/stories-generator/src/constants"
	fileparser "github.com/hellosooka/stories-generator/src/files"
	"github.com/hellosooka/stories-generator/src/utils"
	"log"
	"os"
	"sync"
	"text/template"
)

var temp *template.Template

const vueTemplate string = constants.VUE_TEMPLATE_FILENAME + constants.TEMPLATE_EXTEND

func CreateVueStories(path string, templatePath string, section string) {
	var wg sync.WaitGroup
	temp = utils.GetFilteredTemplates(templatePath, vueTemplate)
	files := fileparser.GetStoryItems(parseVueFilesPath(path), section)

	wg.Add(len(files))
	for i := 0; i < len(files); i++ {
		go func(i int) {
			createStory(files[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func createStory(story fileparser.StoryItem) {
	f, err := os.Create(fmt.Sprintf("%s%s", story.Directory, createStoryFilename(story.Filename)))
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(f, story)
}

func createStoryFilename(filename string) string {
	return fmt.Sprintf("%s%s.ts", filename, constants.STORY_POSTFIX)
}

func parseVueFilesPath(path string) []string {
	files, err := fileparser.ParseFilesPath(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileparser.FilterByExtend(files, constants.VUE_EXTEND)
}
