package main

import (
	"fmt"
	"prokop/sifry/analyza"
)

func analyse(raw []byte) {
	var c uint64
	data := analyza.Text(raw)
	if st {
		c = data.Chars()
		fmt.Printf("Počet znaků:\t%d\n", c)
		fmt.Printf("Počet slabik:\t%d\n", data.Slabiky())
		fmt.Printf("Počet slov:\t%d\n", data.Words())
		fmt.Printf("Počet vět:\t%d\n", data.Sentences())
		fmt.Printf("Počet řádek:\t%d\n", data.Lines())
	}
	if fl {
		fmt.Printf("Fleshův index:\t%f\n", data.Flesh())
	}

	if !ws {
		data.RemoveWS()
		c = data.Chars()
	}
	if std {
		data = data.Stdr()
	}
	if ad {
		data.AlphaD()
		c = data.Chars()
	}
	if fr {
		fre := data.Frekvence()
		Map := sort(fre)
		fmt.Println("frekvence znaků:")
		for _, v := range Map {
			pourcent := float64((100.0 * float64(v.val)) / float64(c))
			fmt.Printf("%c:%d\t:%6.2f %%\n", v.key, v.val, pourcent)
		}
	}
}
