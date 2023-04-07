package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// TODO I think there should be an "all or nothing" policy in implementing
// restore, so everything is just set back to the same of how it was before.
func restore(uri []string) error {
	meta, err := os.ReadFile(MetaLoc)
	if err != nil {
		return fmt.Errorf("Couldn't read metadata file from%s; %w", MetaLoc, err)
	}
	mdir := string(meta)
	for _, u := range uri {
		tu := strings.TrimPrefix(u, RLoc)
		fileInfo, err := os.Lstat(u)
		if err != nil {
			log.Printf("Error getting file %s info; %v.. skipping\n", u, err)
			continue
		}
		if fileInfo.Mode()&os.ModeDir != 0 && fileInfo.Mode()&os.ModeSymlink == 0 {
			fs, err := listDir(u)
			if err != nil {
				log.Printf("Error getting directory %s info; %v.. skipping\n",
					u, err)
				continue
			}
			err = restore(func(arr []os.FileInfo) []string {
				var res []string
				for _, f := range arr {
					res = append(res, filepath.Join(u, f.Name()))
				}
				return res
			}(fs))
			if err != nil {
				log.Printf("Couldn't restore directory %s; %v.. skipping\n", u,
					err)
				continue
			}

			err = os.RemoveAll(u)
			if err != nil {
				log.Printf("Couldn't remove dummy dir %s; %v.. skipping\n", u, err)
			}
			continue
		}

		if fileInfo.Mode()&os.ModeSymlink == 0 {
			log.Printf("file %s is not a link nor a dir of links; ..skipping\n", u)
		}
		rUri := filepath.Join(mdir, tu)
		err = cliRestore(rUri)
		if err != nil {
			log.Printf("Couldn't restore %s; %v.. skipping\n", rUri, err)
			continue
		}
		err = os.Remove(u)
		if err != nil {
			log.Printf("Couldn't remove link %s; %v.. skipping\n", rUri, err)
		}
	}
	return nil
}

func cliRestore(path string) error {
	// TODO should be refactored to use only one process

	cmd := exec.Command("trash-restore", path)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Error creating stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Error starting command: %w", err)
	}

	scanner := bufio.NewScanner(stdout)
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	_ = cmd.Process.Kill()

	cmd = exec.Command("trash-restore", path)
	cmd.Stdin = bytes.NewReader([]byte("0\n"))
	if len(output) == 2 {
		cmd.Run()
		return nil
	}
	if len(output) == 3 {
		fp := output[0][strings.Index(output[0], "/"):]
		sp := output[1][strings.Index(output[1], "/"):]
		if fp == sp {
			cmd.Run()
			return nil
		}
	}
	return fmt.Errorf(
		"Unexpected output from trash-restore, expected 2 lines output but got: %d. Full output: \n %s",
		len(output), strings.Join(output, "\n"))
}
