package main

import (
	"log"

	"github.com/coldinke/leveldb-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
