package utils

import (
	"github.com/hellosooka/stories-generator/src/files"
	"log"
	"text/template"
)

var temp *template.Template

func GetTemplates(templatesPath string) *template.Template {
	temps, err := fileparser.ParseTemplatePaths(templatesPath)
	if err != nil {
		log.Fatal(err)
	}
	return template.Must(template.ParseFiles(temps...))
}
