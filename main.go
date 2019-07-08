package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var ws bool
var an bool
var std bool
var ad bool
var fl bool
var fr bool

func main() {
	var file string

	flag.StringVar(&file, "f", "", "file to analyse")
	flag.BoolVar(&ws, "ws", false, "frequency with white spaces")
	flag.BoolVar(&an, "a", true, "analyse file")
	flag.BoolVar(&std, "std", false, "standart data")
	flag.BoolVar(&ad, "ad", false, "only alpha-digit; also standart")
	flag.BoolVar(&fl, "fl", false, "display Flesh index")
	flag.BoolVar(&fr, "fr", true, "display frequency")
	flag.Parse()
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Nelze přečíst soubor")
	}
	if an {
		analyse(raw)
	}
}
