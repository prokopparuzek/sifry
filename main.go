package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var file string

	flag.StringVar(&file, "f", "", "file to analyse")
	flag.Parse()
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Nelze přečíst soubor")
	}
	analyse(raw)
}
