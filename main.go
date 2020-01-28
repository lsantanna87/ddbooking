package main

import (
	"log"

	"github.com/lsantanna87/ddbooking/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Panic("Error when invoking execute command... paniking")
	}
}
