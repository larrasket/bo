package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if len(os.Args) > 1 && os.Args[1] == "r" {
		// TODO restore
		return
	}
	err := setup()
	if err != nil {
		log.Fatal(err)
	}
}
