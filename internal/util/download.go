package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(href, fileName string) error {
	response, err := http.Get(href)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		bs, _ := io.ReadAll(response.Body)
		return fmt.Errorf("got %d from: %s\n\t%s", response.StatusCode, href, string(bs))
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
