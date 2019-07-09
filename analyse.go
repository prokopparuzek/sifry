package main

import (
	"fmt"
	"prokop/sifry/analyza"
)

func analyse(raw []byte, data analyza.Text) {
	var c uint64
	w := data.Words()
	if st { // Statistiky textu
		c = data.Chars()
		fmt.Printf("Počet znaků:\t%d\n", c)
		fmt.Printf("Počet slabik:\t%d\n", data.Slabiky())
		fmt.Printf("Počet slov:\t%d\n", data.Words())
		fmt.Printf("Počet vět:\t%d\n", data.Sentences())
		fmt.Printf("Počet řádek:\t%d\n", data.Lines())
		if fl { // Fleshův index
			fmt.Printf("Fleshův index:\t%f\n", data.Flesh())
		}
	}
	if fw { // Frekvence slov
		fre := data.FrekvenceSlov()
		Map := sort(fre)
		fmt.Println("frekvence slov:")
		for _, v := range Map {
			pourcent := float64((100.0 * float64(v.val)) / float64(w))
			fmt.Printf("%8s\t:%d\t:%6.2f %%\n", v.key, v.val, pourcent)
		}
	}
	trimSpace(&data)
	c = data.Chars()
	if fr { // frekvence znaků
		fre := data.Frekvence()
		Map := sort(fre)
		fmt.Println("frekvence znaků:")
		for _, v := range Map {
			pourcent := float64((100.0 * float64(v.val)) / float64(c))
			fmt.Printf("%s:%d\t:%6.2f %%\n", v.key, v.val, pourcent)
		}
	}
}
