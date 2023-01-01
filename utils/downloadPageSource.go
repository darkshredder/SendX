package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func DownloadPageSource(uri string, id string) bool {

	resp, err := http.Get(uri)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	fmt.Printf("\nDownloaded %d\n", resp.Body)

	bodyVal, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	fmt.Printf("\nRead %s\n", bodyVal)
	// Write the response body to a file
	err = ioutil.WriteFile("./downloads/"+id+".html", bodyVal, 0644)
	if err != nil {
		return false
	}

	return true

}

func DownloadPageSourceJob(uri string, id string) bool {

	resp := DownloadPageSource(uri, id)
	return resp

}
