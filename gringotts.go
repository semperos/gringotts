package gringotts

import (
	"io/ioutil"
	"log"
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

func DownloadFile(url string, localFile string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to download '%s' with error: %v\n", url, err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Download of '%s' returned non-2xx status in response: %v\n", resp.Request.URL, resp)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read file downloaded from '%s' with error: %v\n", url, err)
		} else {
			err := ioutil.WriteFile(localFile, body, 0755)
			if err != nil {
				log.Fatalf("Failed to write downloaded file to local path '%s' with error: %v\n", localFile, err)
			}
		}
	}
	return localFile
}
