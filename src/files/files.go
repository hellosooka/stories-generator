package fileparser

import (
	"fmt"
	"strings"

	"github.com/hellosooka/stories-generator/src/constants"
)

type storyItem struct {
	Filename     string
	FullFilename string
	Section      string
	Directory    string
	path         string
}

func ParseTemplatePaths(rootDir string) ([]string, error) {
	files, err := ParseFilesPath(rootDir)
	if err != nil {
		return nil, err
	}
	return FilterByExtend(files, constants.TEMPLATE_EXTEND), nil
}

func GetStoryItems(files []string, section string) []storyItem {
	var result []storyItem
	for i := 0; i < len(files); i++ {
		result = append(result, pathToStoryItem(files[i], section))
	}
	return result
}

func pathToStoryItem(path string, section string) storyItem {
	parsPath := strings.Split(path, "/")

	return storyItem{
		Filename:     clearFileExtend(parsPath[len(parsPath)-1]),
		FullFilename: parsPath[len(parsPath)-1],
		Section:      section,
		Directory:    clearLastInPath(path),
		path:         path,
	}
}

func clearFileExtend(filename string) string {
	return strings.Split(filename, ".")[0]
}

func clearLastInPath(path string) string {
	parsPath := strings.Split(path, "/")
	var result string
	for i := 0; i < len(parsPath)-1; i++ {
		result += fmt.Sprintf("%s/", parsPath[i])
	}
	return result
}
