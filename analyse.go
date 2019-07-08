package main

import (
	"fmt"
	"prokop/sifry/analyza"
	"strings"
)

func analyse(raw []byte) {
	data := analyza.Text(raw)
	w := data.Words()
	fmt.Printf("Počet znaků:\t%d\n", data.Chars())
	fmt.Printf("Počet slabik:\t%d\n", data.Slabiky())
	fmt.Printf("Počet slov:\t%d\n", w)
	fmt.Printf("Počet vět:\t%d\n", data.Sentences())
	fmt.Printf("Počet řádek\t%d\n", data.Lines())

	if !ws {
		data = analyza.Text(strings.Replace(string(data), " ", "", -1))
		data = analyza.Text(strings.Replace(string(data), "\t", "", -1))
		data = analyza.Text(strings.Replace(string(data), "\n", "", -1))
	}
	fr := data.Frekvence()
	Map := sort(fr)
	fmt.Println("Frekvence znaků:")
	for _, v := range Map {
		pourcent := float64((100.0 * float64(v.val)) / float64(w))
		fmt.Printf("%c:%d\t:%6.2f %%\n", v.key, v.val, pourcent)
	}
}
