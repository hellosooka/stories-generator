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
	clearInter := strings.ReplaceAll(inter, " ", "")
	splittedInter := strings.Split(clearInter, "\n")
	result := []string{}
	for i := 0; i < len(splittedInter); i++ {
		line := splittedInter[i]
		if strings.Contains(line, ":") {
			splittedLine := strings.Split(line, ":")
			result = append(result, fmt.Sprintf("\n    %s:%s", strings.ReplaceAll(splittedLine[0], "?", ""), "{}"))
		}
	}
	return fmt.Sprintf("{%s}", strings.Join(result, ","))
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
	if len(f.FindStringSubmatch(file)) > 2 {
		return f.FindStringSubmatch(file)[1]
	}
	return ""
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
