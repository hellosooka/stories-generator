package fileparser

import (
	"fmt"
	"strings"

	"github.com/hellosooka/stories-generator/src/constants"
)

func ParseTemplatePaths(rootDir string) ([]string, error) {
	files, err := ParseFilesPath(rootDir)
	if err != nil {
		return nil, err
	}
	return FilterByExtend(files, constants.TEMPLATE_EXTEND), nil
}

func ClearFileExtend(filename string) string {
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
