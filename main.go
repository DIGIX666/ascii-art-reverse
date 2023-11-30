package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		reverse(args)
	} else {
		fmt.Println("Too many arguments")
	}
}
