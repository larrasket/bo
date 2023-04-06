package main

import (
	"os"
	"path/filepath"
	"strings"
)

func setup() (err error) {
	err = removeDirectoryIfExists(RLoc)
	if err != nil {
		return
	}

	os.Mkdir(RLoc, 0755)

	lst, err := ListInfo()
	if err != nil {
		return
	}

	if len(lst) == 0 {
		return
	}

	cwd, err := os.Getwd()
	cwd = "/home/ghd/"
	if err != nil {
		return
	}

	for _, value := range lst {
		p, _ := Parse(value)
		if !strings.Contains(p, cwd) {
			continue
		}

		p = strings.TrimPrefix(p, cwd)
		fp := strings.LastIndex(p, "/")
		name := strings.TrimSuffix(value.Name(), ".trashinfo")
		tfloc := filepath.Join(TFilesLoc, name)
		if fp == -1 {
			os.Symlink(tfloc, RLoc)
			continue
		}

		err = os.MkdirAll(filepath.Join(RLoc, p[:fp]), 0755)
		if err != nil {
			_ = removeDirectoryIfExists(RLoc)
			return
		}

		os.Symlink(tfloc, filepath.Join(filepath.Join(RLoc, p[:fp]), name))
	}
	return
}
