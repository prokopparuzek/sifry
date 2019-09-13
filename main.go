package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"github.com/prokopparuzek/sifry_lib/analyza"
	"github.com/prokopparuzek/sifry_lib/change"
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
var ciphre string
var height uint64
var weight uint64
var decrypt bool

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
	flag.Int64Var(&combien, "cmb", 1, "how many words || chars")
	flag.Int64Var(&lenght, "l", 0, "lenght of text")
	flag.StringVar(&ciphre, "c", "", "what ciphre")
	flag.Uint64Var(&weight, "w", 0, "weight")
	flag.Uint64Var(&height, "h", 0, "height")
	flag.BoolVar(&decrypt, "d", false, "crypt || decrypt")
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
	switch ciphre {
	case "RectangleL":
		if !decrypt {
			r := crypt.Rectangle{weight, height}
			fmt.Printf("%s", r.CryptL(&data))
		} else {
			r := crypt.Rectangle{weight, height}
			fmt.Printf("%s\n", r.DecryptL(&data))
		}
	case "RectangleR":
		if !decrypt {
			r := crypt.Rectangle{weight, height}
			fmt.Printf("%s", r.CryptR(&data))
		} else {
			r := crypt.Rectangle{weight, height}
			fmt.Printf("%s\n", r.DecryptR(&data))
		}
	case "Reverse":
		fmt.Printf("%s", crypt.Reverse(&data))
	case "Teeth":
		if !decrypt {
			r := crypt.Teeth(height)
			fmt.Printf("%s", r.Crypt(&data))
		} else {
			r := crypt.Teeth(height)
			fmt.Printf("%s\n", r.Decrypt(&data))
		}
	case "Stairs":
		if !decrypt {
			r := crypt.Stairs(weight)
			fmt.Printf("%s\n", r.Crypt(&data))
		} else {
			r := crypt.Stairs(weight)
			fmt.Printf("%s\n", r.Decrypt(&data))
		}
	case "Snake":
		if !decrypt {
			r := crypt.Snake(height)
			fmt.Printf("%s", r.Crypt(&data))
		} else {
			r := crypt.Snake(height)
			fmt.Printf("%s\n", r.Decrypt(&data))
		}
	case "JumpNS":
		if !decrypt {
			r := crypt.Jump(height)
			fmt.Printf("%s\n", r.CryptNS(&data))
		} else {
			r := crypt.Jump(height)
			fmt.Printf("%s\n", r.DecryptNS(&data))
		}
	default:
		log.Fatal("Nenalezená šifra!")
	}
}
