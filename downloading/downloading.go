package downloading

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	cachePath      = "downloading/cache/"
	cacheExtension = ".txt"

	remoteURLBase   = "https://raw.githubusercontent.com/tldr-pages/tldr/master/pages/"
	remoteDir       = "common/"
	remoteExtension = ".md"
)

// GetCommandDesc - takes a command "cmd" and returns the tldr
// description of it, via https://github.com/tldr-pages/tldr/tree/master/pages
// returns ptr to string because it might be rather large
func GetCommandDesc(cmd string) *string {
	if desc := getCachedDesc(cmd); desc != nil {
		return desc
	}

	if desc := getRemoteDesc(cmd); desc != nil {
		return desc
	}
	return nil
}

// getCachedDesc - attempts to get cached tldr description
// saved in downloading/cache/,
func getCachedDesc(cmd string) *string {
	desc, err := ioutil.ReadFile(cachePath + cmd + cacheExtension)
	if err != nil {
		return nil
	}
	result := string(desc)
	return &result
}

// getRemoteDesc - attempts to download description of "cmd"
// from https://github.com/tldr-pages/tldr/tree/master/pages,
// TODO: add support for specific OSs, starting with the one the user is
// currently using
func getRemoteDesc(cmd string) *string {
	desc, err := tryDownload(remoteURLBase + remoteDir + cmd + remoteExtension)
	if err != nil {
		return nil
	}

	err = tryWriteToFile(desc, cmd)
	if err != nil {
		fmt.Println("failed to cache description")
	}

	return desc
}

// tryDownload - attempts to download file at "url",
func tryDownload(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := string(content)
	return &result, nil
}

// tryWriteToFile - attempts to write file with path "cache/{fName}.txt"
// with content "content"
func tryWriteToFile(content *string, fName string) error {
	f, err := os.Create(cachePath + fName + cacheExtension)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.WriteString(f, *content)
	if err != nil {
		return err
	}

	return f.Sync()
}
