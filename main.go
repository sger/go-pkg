package main

import (
	"log"
	"github.com/mitchellh/colorstring"
	"fmt"
)

func main() {
	colorstring.Println("[blue]Hello [red]World!")

	resp, err := search("log")

	if err != nil {
		log.Fatal(err)
	}

	for _, element := range resp.Results {
		fmt.Println("%s", element.Name)
	}
}