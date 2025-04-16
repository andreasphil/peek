package lib

import (
	"os"
	"path/filepath"
)

func readFile(filename string) ([]byte, error) {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	return content, nil
}
