package vueParser

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/hellosooka/stories-generator/src/constants"
	fileparser "github.com/hellosooka/stories-generator/src/files"
	story "github.com/hellosooka/stories-generator/src/stories"
)

func GetProps(story story.StoryItem) string {
	inter, _ := parseProps(readVueFile(story))
	return formatProps(inter)
}

func formatProps(inter string) string {
	clearInter := strings.ReplaceAll(inter, ",", "")
	commedInter := strings.Replace(strings.Join(strings.Split(clearInter, "\n"), " ,"), ",", "", 1)
	splittedInter := strings.Split(commedInter, " ")
	for i := 0; i < len(splittedInter); i++ {
		reg := regexp.MustCompile("}|{|:|,")
		if !reg.MatchString(splittedInter[i]) && len(splittedInter[i]) > 0 {
			splittedInter[i] = fmt.Sprintf("'%s'", splittedInter[i])
		}
	}

	return strings.Join(splittedInter, " ")
}

func parseProps(file string) (string, error) {
	f := regexp.MustCompile(fmt.Sprintf("%s%s", constants.VUE_PROPS_SETUP, `<([\s\S]+?)>`))
	if f.MatchString(file) {
		str := f.FindStringSubmatch(file)[1]
		if strings.Contains(str, "{") {
			return str, nil
		}
		return findInterface(file, str), nil
	}
	return "{}", errors.New("empty")
}

func findInterface(file string, name string) string {
	f := regexp.MustCompile(fmt.Sprintf(`interface %s\s*({([^{}]|{[^{}]*})*})`, name))
	return f.FindStringSubmatch(file)[1]
}

func readVueFile(story story.StoryItem) string {
	file, err := os.ReadFile(fmt.Sprintf("%s%s%s", story.Directory, story.Filename, constants.VUE_EXTEND))
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func ParseVueFilesPath(path string) []string {
	files, err := fileparser.ParseFilesPath(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileparser.FilterByExtend(files, constants.VUE_EXTEND)
}
