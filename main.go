package main

import (
	"log"
	"github.com/mitchellh/colorstring"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: go-pkg sql\n")
		os.Exit(1)
	}

	resp, err := search(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	for _, element := range resp.Results {
		str := colorstring.Color(fmt.Sprintf("[red] Name\n [blue]%s", element.Name))
		colorstring.Println(str)
	}
}