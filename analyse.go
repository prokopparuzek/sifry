package main

import (
	"fmt"
	"prokop/sifry/analyza"
)

func analyse(data *string) {
	var c uint64
	w := analyza.Words(data)
	if st { // Statistiky textu
		c = analyza.Chars(data)
		fmt.Printf("Počet znaků:\t%d\n", c)
		fmt.Printf("Počet slabik:\t%d\n", analyza.Slabiky(*data))
		fmt.Printf("Počet slov:\t%d\n", analyza.Words(data))
		fmt.Printf("Počet vět:\t%d\n", analyza.Sentences(data))
		fmt.Printf("Počet řádek:\t%d\n", analyza.Lines(data))
		if fl { // Fleshův index
			fmt.Printf("Fleshův index:\t%f\n", analyza.Flesh(data))
		}
	}
	if fw { // Frekvence slov
		fre := analyza.FrekvenceSlov(data)
		Map := sort(fre)
		fmt.Println("frekvence slov:")
		for _, v := range Map {
			pourcent := float64((100.0 * float64(v.val)) / float64(w))
			fmt.Printf("%8s\t:%d\t:%6.2f %%\n", v.key, v.val, pourcent)
		}
	}
	trimSpace(data)
	c = analyza.Chars(data)
	if fr { // frekvence znaků
		fre := analyza.Frekvence(data)
		Map := sort(fre)
		fmt.Println("frekvence znaků:")
		for _, v := range Map {
			pourcent := float64((100.0 * float64(v.val)) / float64(c))
			fmt.Printf("%s:%d\t:%6.2f %%\n", v.key, v.val, pourcent)
		}
	}
}
