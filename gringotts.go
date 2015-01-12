package gringotts

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DoesFileExist(pathname string) bool {
	_, err := os.Stat(pathname)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else if os.IsNotExist(err) {
			return false
		} else {
			log.Fatalf("Unrecoverable error has occurred: %v.\n", err)
		}
	}
	return true
}

func DownloadFile(url string, localFile string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	} else if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Download of '%s' returned non-2xx status in response: %v\n", resp.Request.URL, resp))
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		} else {
			err := ioutil.WriteFile(localFile, body, 0755)
			if err != nil {
				return "", err
			}
		}
	}
	return localFile, nil
}

func DownloadFileOrFail(url string, localFile string) string {
	ret, err := DownloadFile(url, localFile)
	if err != nil {
		log.Fatalf("Failed to download %s to %s", url, localFile)
	}
	return ret
}
