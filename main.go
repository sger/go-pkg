package main

import (
	"log"
	"github.com/mitchellh/colorstring"
	"fmt"
)

func main() {

	resp, err := search("log")

	if err != nil {
		log.Fatal(err)
	}

	for _, element := range resp.Results {
		str := colorstring.Color(fmt.Sprintf("[red] Name [blue]%s", element.Name))
		colorstring.Println(str)
	}
}