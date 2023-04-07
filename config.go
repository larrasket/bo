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
	// If you have nested trash dirs like me, you might get incomplete trash
	// paths in your trash info files. For example I get files like the
	// following:
	// [Trash Info]
	// Path=me/temp/20230323_035707.jpg
	// DeletionDate=2023-04-07T10:48:41
	// the "me" directory is inside "/mnt/disk". So I will put "/mnt/disk" as a
	// value. Safely ignore it if you don't suffer from this issue in your
	// setup.

	ParentLoc = "/mnt/disk/"
)
