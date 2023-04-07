package main

import (
	"log"
	"os"
)

// configuration locations
var (
	home = func() string {
		res, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Couldn't get user's home dir: ", err)
		}
		return res
	}()

	Tloc      = home + "/.local/share/Trash/"
	TInfoLoc  = Tloc + "info/"
	TFilesLoc = Tloc + "files/"
	RLoc      = "/tmp/boBrowser/"
	MetaLoc   = RLoc + ".bin_organizer12"
)
