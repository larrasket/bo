package main

import (
	"os"
)

func parse(f os.FileInfo) (string, error) {
	p := TInfoLoc + f.Name()
	dat, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return extractPath(string(dat)), nil

}
