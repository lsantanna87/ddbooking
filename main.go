package main

import (
	"fmt"

	"github.com/lsantanna87/ddbooking/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		fmt.Printf("error when executing cmd... %+v", err)
	}
}
