package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func DataReader(dirName string) (map[string]string, error) {
	files := make(map[string]string)
	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			parentDir := filepath.Dir(path)
			parts := strings.Split(parentDir, "/")

			if len(parts) > 0 {
				parentDir = parts[len(parts)-1]
			} else {
				return errors.New("Unable to get dir name")
			}

			files[string(fileData)] = parentDir

			fmt.Println("File:", path)
			fmt.Println(string(fileData))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
