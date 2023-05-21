package vue_parser

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/hellosooka/stories-generator/src/constants"
	fileparser "github.com/hellosooka/stories-generator/src/files"
	"github.com/hellosooka/stories-generator/src/utils"
)

var temp *template.Template

const vueTemplate string = constants.VUE_TEMPLATE_FILENAME + constants.TEMPLATE_EXTEND

func CreateVueStories(path string, templatePath string) {
	temp = utils.GetFilteredTemplates(templatePath, vueTemplate)

	files := fileparser.GetStoryItems(parseVueFilesPath(path))

	for i := 0; i < len(files); i++ {
		f, err := os.Create(fmt.Sprintf("%s%s", files[i].Directory, createStoryFilename(files[i].Filename)))
		if err != nil {
			log.Fatal(err)
		}
		temp.Execute(f, files[i])
	}
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
