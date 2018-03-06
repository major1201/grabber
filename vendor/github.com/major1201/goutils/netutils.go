package goutils

import (
	"os"
	"net/http"
	"io"
)

func Download(url string, dest string) (err error) {
	out, fileError := os.Create(dest)
	defer out.Close()
	if fileError != nil {
		err = fileError
		return
	}
	resp, httpError := http.Get(url)
	defer resp.Body.Close()
	if httpError != nil {
		err = httpError
		return
	}
	_, copyError := io.Copy(out, resp.Body)
	if copyError != nil {
		err = copyError
		return
	}
	return err
}
