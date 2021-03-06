package main

import (
	"github.com/prokopparuzek/sifry_lib/change"
)

func trim(data *string) {
	if std { // Standartizace textu, pouze malá bez interpunkce
		change.Stdr(data)
	}
	if ad { // Pouze alfanumerické znaky
		change.AlphaD(data)
	}
}

func trimSpace(data *string) {
	if !ws { // Odstranění bílých znaků
		change.RemoveWS(data)
	}
}
