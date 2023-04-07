package main

import (
	"io/ioutil"
	"os"
	"regexp"
)

func removeDirectoryIfExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}
	return os.RemoveAll(path)
}

func listDir(dir string) ([]os.FileInfo, error) {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func extractPath(dat string) string {
	re := regexp.MustCompile(`Path=([\S]+)`)
	matches := re.FindStringSubmatch(string(dat))
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}
