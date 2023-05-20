package fileparser

import (
	"strings"

	"github.com/hellosooka/stories-generator/src/constants"
)

type StoryItem struct {
	Filename     string
	FullFilename string
	Section      string
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
		path:         path,
	}
}

func clearFileExtend(filename string) string {
	return strings.Split(filename, ".")[0]
}
