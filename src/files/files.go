package fileparser

import (
	"fmt"
	"strings"

	"github.com/hellosooka/stories-generator/src/constants"
)

type StoryItem struct {
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

func GetStoryItems(files []string) []StoryItem {
	var result []StoryItem
	for i := 0; i < len(files); i++ {
		result = append(result, pathToStoryItem(files[i]))
	}
	return result
}

func pathToStoryItem(path string) StoryItem {
	parsPath := strings.Split(path, "/")

	return StoryItem{
		Filename:     clearFileExtend(parsPath[len(parsPath)-1]),
		FullFilename: parsPath[len(parsPath)-1],
		Section:      parsPath[0],
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
