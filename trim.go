package main

import (
	"prokop/sifry/analyza"
)

func trim(data *analyza.Text) {
	if std { // Standartizace textu, pouze malá bez interpunkce
		*data = data.Stdr()
	}
	if ad { // Pouze alfanumerické znaky
		data.AlphaD()
	}
}

func trimSpace(data *analyza.Text) {
	if !ws { // Odstranění bílých znaků
		data.RemoveWS()
	}
}
