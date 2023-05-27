package stories

import (
	"strings"

	fileparser "github.com/hellosooka/stories-generator/src/files"
)

type StoryItem struct {
	Filename     string
	FullFilename string
	Section      string
	Directory    string
	Props        string
	path         string
}

func GetStoryItems(files []string, section string) []StoryItem {
	var result []StoryItem
	for i := 0; i < len(files); i++ {
		result = append(result, pathToStoryItem(files[i], section))
	}
	return result
}

func pathToStoryItem(path string, section string) StoryItem {
	parsPath := strings.Split(path, "/")

	return StoryItem{
		Filename:     fileparser.ClearFileExtend(parsPath[len(parsPath)-1]),
		FullFilename: parsPath[len(parsPath)-1],
		Section:      section,
		Directory:    fileparser.ClearLastInPath(path),
		path:         path,
	}
}
