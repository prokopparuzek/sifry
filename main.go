package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"prokop/sifry/analyza"
	"strings"
)

var ws bool
var an bool
var std bool
var ad bool
var fl bool
var fr bool
var st bool
var fw bool
var re bool
var what string
var combien int64
var lenght int64

func main() {
	var file string
	var raw []byte

	flag.StringVar(&file, "f", "", "file to analyse")
	flag.BoolVar(&ws, "ws", false, "frequency with white spaces")
	flag.BoolVar(&an, "a", false, "analyse file")
	flag.BoolVar(&std, "std", false, "standart data")
	flag.BoolVar(&ad, "ad", false, "only alpha-digit")
	flag.BoolVar(&fl, "fl", false, "display Flesh index")
	flag.BoolVar(&fr, "fr", true, "display frequency")
	flag.BoolVar(&fw, "fw", false, "display frequency of words")
	flag.BoolVar(&st, "st", true, "display statistiques")
	flag.BoolVar(&re, "re", false, "reproduct text")
	flag.StringVar(&what, "what", "words", "words || chars")
	flag.Int64Var(&combien, "c", 1, "how many words || chars")
	flag.Int64Var(&lenght, "l", 0, "lenght of text")
	flag.Parse()
	if file == "" || file == "-" {
		reader := bufio.NewReader(os.Stdin)
		for {
			b, err := reader.ReadByte()
			if err == io.EOF {
				break
			}
			raw = append(raw, b)
		}
	} else {
		var err error
		raw, err = ioutil.ReadFile(file)
		if err != nil {
			log.Fatal("Nelze přečíst soubor")
		}
	}
	data := string(raw)
	trim(&data)
	if an {
		analyse(&data)
	}
	if re {
		var w bool
		if what == "words" {
			w = true
		} else if what == "chars" {
			w = false
		} else {
			log.Fatal("Pozor words || chars")
		}
		fmt.Printf("%s\n", strings.TrimSpace(analyza.Reproduct(&data, w, uint64(combien), uint64(lenght))))
	}
}
