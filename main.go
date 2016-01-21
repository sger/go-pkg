package main

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"log"
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
		str := colorstring.Color(fmt.Sprintf("[red]Name:[blue]\t\t%s \n[red]Package:[blue]\t%s \n[red]Author:[blue]\t\t%s \n[red]ProjectURL:[blue]\t%s", element.Name, element.Package, element.Author, element.ProjectURL))
		colorstring.Println(str)
		fmt.Println(" ")
	}

}
