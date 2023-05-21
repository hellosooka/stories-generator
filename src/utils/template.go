package utils

import (
	"log"
	"strings"
	"text/template"

	"github.com/hellosooka/stories-generator/src/files"
)

var temp *template.Template

func getTemplates(templatesPath string) []string {
	temps, err := fileparser.ParseTemplatePaths(templatesPath)
	if err != nil {
		log.Fatal(err)
	}
	return temps
}

func GetFilteredTemplates(templatesPath string, filter string) *template.Template {
	var result []string
	temps := getTemplates(templatesPath)

	for i := 0; i < len(temps); i++ {
		if strings.Contains(temps[i], filter) {
			result = append(result, temps[i])
		}
	}

	return template.Must(template.ParseFiles(result...))
}
