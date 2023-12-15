package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error repo not be create : %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error validation repo : %v", err)
	}
	return nil
}

// Check file extension
func IsImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".bmp"
}
