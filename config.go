package main

import "os"

// configuration locations
var (
	home = func() string {
		res, _ := os.UserHomeDir()
		return res
	}()

	Tloc      = home + "/.local/share/Trash/"
	TInfoLoc  = Tloc + "info/"
	TFilesLoc = Tloc + "files/"
	RLoc      = "/tmp/boBrowser"
)
