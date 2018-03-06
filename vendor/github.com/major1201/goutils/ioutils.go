package goutils

import (
	"os"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func IsFile(path string) bool {
	fi, err := os.Stat(path)
	if err == nil && fi.Mode().IsRegular() == true {
		return true
	}
	return false
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err == nil && fi.Mode().IsDir() == true {
		return true
	}
	return false
}
