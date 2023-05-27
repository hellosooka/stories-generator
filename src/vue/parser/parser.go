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
	f := regexp.MustCompile(`{([\s\S]+?)}`)
	if f.MatchString(inter) {
		str := strings.ReplaceAll(strings.TrimSpace(f.FindStringSubmatch(inter)[1]), ",", "")
		splittedStr := strings.Split(str, "\n")

		for i := 0; i < len(splittedStr); i++ {
			s := strings.Split(splittedStr[i], ":")
			splittedStr[i] = fmt.Sprintf(`%s:'%s'`, s[0], strings.TrimSpace(s[1]))
		}

		return fmt.Sprintf("{%s}", strings.Join(splittedStr, ","))
	}
	return ""
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
	return "", errors.New("empty")
}

func findInterface(file string, name string) string {
	f := regexp.MustCompile(fmt.Sprintf(`interface %s([\s\S]+?){([\s\S]+?)}`, name))
	return fmt.Sprintf(`{%s}`, f.FindStringSubmatch(file)[2])
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
