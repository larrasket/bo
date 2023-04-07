package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if len(os.Args) > 2 && os.Args[1] == "r" {
		err := restore(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err := setup()
	if err != nil {
		log.Fatal(err)
	}
}
