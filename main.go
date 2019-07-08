package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var ws bool
var an bool

func main() {
	var file string

	flag.StringVar(&file, "f", "", "file to analyse")
	flag.BoolVar(&ws, "ws", false, "frequency with white spaces")
	flag.BoolVar(&an, "a", true, "analyse file")
	flag.Parse()
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Nelze přečíst soubor")
	}
	if an {
		analyse(raw)
	}
}
