package downloading

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	cachePath      = "cache/"
	cacheExtension = ".txt"

	remoteURLBase   = "raw.githubusercontent.com/tldr-pages/tldr/master/pages/"
	remoteExtension = ".md"
)

// GetCommandDesc - takes a command "cmd" and returns the tldr
// description of it, via https://github.com/tldr-pages/tldr/tree/master/pages
// returns error when cmd is not recognized
func GetCommandDesc(cmd string) (string, error) {
	if desc, ok := getCachedDesc(cmd); ok {
		return desc, nil
	}

	if desc, ok := getRemoteDesc(cmd); ok {
		return desc, nil
	}

	return "", fmt.Errorf("command '%s' not recognized", cmd)
}

// getCachedDesc - attempts to get cached tldr description
// saved in downloading/cache/,
// follows comma ok idiom
func getCachedDesc(cmd string) (string, bool) {
	desc, err := ioutil.ReadFile(cachePath + cmd + cacheExtension)
	if err != nil {
		return "", false
	}
	return string(desc), true
}

// getRemoteDesc - attempts to download description of "cmd"
// from https://github.com/tldr-pages/tldr/tree/master/pages,
// follows comma ok idiom
// TODO: add support for specific OSs, starting with the one the user is
// currently using
func getRemoteDesc(cmd string) (string, bool) {
	desc, err := tryDownload(remoteURLBase + cmd + remoteExtension)
	if err != nil {
		return "", false
	}

	if err = tryWriteToFile(desc, cmd); err != nil {
		return "", false
	}

	return desc, true
}

// tryDownload - attempts to download file at "url",
// follows comma error idiom
func tryDownload(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// tryWriteToFile - attempts to write file with path "cache/{fName}.txt"
// with content "content"
func tryWriteToFile(content string, fName string) error {
	f, err := os.Create(cachePath + fName + cacheExtension)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.WriteString(f, content)
	if err != nil {
		return err
	}

	return f.Sync()
}
