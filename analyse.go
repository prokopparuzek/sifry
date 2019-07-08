package main

import (
	"fmt"
	"log"
	"prokop/sifry/analyza"
	"regexp"
	"strings"
)

func analyse(raw []byte) {
	data := analyza.Text(raw)
	c := data.Chars()
	fmt.Printf("Počet znaků:\t%d\n", c)
	fmt.Printf("Počet slabik:\t%d\n", data.Slabiky())
	fmt.Printf("Počet slov:\t%d\n", data.Words())
	fmt.Printf("Počet vět:\t%d\n", data.Sentences())
	fmt.Printf("Počet řádek\t%d\n", data.Lines())

	if !ws {
		data = analyza.Text(strings.Replace(string(data), " ", "", -1))
		data = analyza.Text(strings.Replace(string(data), "\t", "", -1))
		data = analyza.Text(strings.Replace(string(data), "\n", "", -1))
		c = data.Chars()
	}
	if std || an {
		data = data.Stdr()
	}
	if an {
		reg, err := regexp.Compile("[^A-Za-z0-9]+")
		if err != nil {
			log.Fatal("Nelze zkompilovat regex!")
		}
		data = analyza.Text(reg.ReplaceAllString(string(data), ""))
	}
	fr := data.Frekvence()
	Map := sort(fr)
	fmt.Println("Frekvence znaků:")
	for _, v := range Map {
		pourcent := float64((100.0 * float64(v.val)) / float64(c))
		fmt.Printf("%c:%d\t:%6.2f %%\n", v.key, v.val, pourcent)
	}
}
