package fileparser

import (
	"os"
	"path/filepath"
	"strings"
)

func ParseTemplatePaths(rootDir string) ([]string, error) {
	files, err := ParseFilesPath(rootDir)
	if err != nil {
		return nil, err
	}
	return FilterByExtend(files, ".tmpl"), nil
}

func ParseFilesPath(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func FilterByExtend(paths []string, extend string) []string {
	var result []string
	for i := 0; i < len(paths); i++ {
		if strings.Contains(paths[i], extend) {
			result = append(result, paths[i])
		}
	}
	return result
}
