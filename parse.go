package main

import (
	"io/ioutil"
	"os"
	"regexp"
)

func Parse(f os.FileInfo) (string, error) {
	p := TInfoLoc + f.Name()
	dat, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return extractPath(string(dat)), nil

}

func ListInfo() ([]os.FileInfo, error) {
	entries, err := ioutil.ReadDir(TInfoLoc)
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
