package fileparser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ParseFiles(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
