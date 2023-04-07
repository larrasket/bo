package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func setup() error {
	err := removeDirectoryIfExists(RLoc)
	if err != nil {
		return fmt.Errorf("The temp dir %s is used or can't be deleted; %w",
			RLoc, err)
	}

	err = os.Mkdir(RLoc, 0755)
	if err != nil {
		return fmt.Errorf("Couldn't create the temp dir %s; %w", RLoc, err)
	}

	lst, err := listDir(TInfoLoc)
	if err != nil {
		return fmt.Errorf("Couldn't read trash content from %s; %w", TInfoLoc,
			err)
	}

	if len(lst) == 0 {
		return nil
	}

	cwd, err := os.Getwd()
	// cwd = "/home/ghd/" // for testing
	if err != nil {
		return fmt.Errorf("Couldn't read cwd; %w", err)
	}

	for _, f := range lst {
		p, _ := parse(f)
		if !strings.Contains(p, cwd) {
			continue
		}

		p = strings.TrimPrefix(p, cwd)
		fp := strings.LastIndex(p, "/")
		name := strings.TrimSuffix(f.Name(), ".trashinfo")
		tfloc := filepath.Join(TFilesLoc, name)
		if fp == -1 {
			os.Symlink(tfloc, RLoc)
			continue
		}

		crt := filepath.Join(RLoc, p[:fp])
		err = os.MkdirAll(crt, 0755)
		if err != nil {
			_ = removeDirectoryIfExists(RLoc)
			return fmt.Errorf("Couldn't create path %s; %w", crt, err)
		}
		os.Symlink(tfloc, filepath.Join(filepath.Join(RLoc, p[:fp]), name))
	}
	f, err := os.Create(MetaLoc)
	if err != nil {
		return fmt.Errorf("Couldn't create metadata file %s; %w", MetaLoc, err)
	}
	defer f.Close()
	_, err = f.WriteString(cwd)
	if err != nil {
		return fmt.Errorf("Couldn't write to the metadata file %s; %w", MetaLoc,
			err)
	}
	return nil
}
