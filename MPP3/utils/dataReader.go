package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Language       string
	Text           string
	CharProportion []float64
}

func DataReader(dirName string) ([]File, error) {
	files := make([]File, 0)
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

			newFile := File{
				Language:       parentDir,
				Text:           string(fileData),
				CharProportion: countAllLetters(string(fileData)),
			}

			files = append(files, newFile)

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

func countAllLetters(text string) []float64 {
	letterProportions := make([]float64, 26)
	letterProportions[25] = -1

	for _, char := range strings.ToLower(text) {
		if 'a' <= char && char <= 'z' {
			letterProportions['a'-char]++
		}
	}

	return letterProportions
}
