package main

import (
	"github.com/monomonedula/dnzbot/tg"
)

func main() {
	telega := tg.NewTelega("2042509427:AAEyXDl2GedH_1NbjhmW695qG2ed6QZIuMY")
	bot := tg.NewBotPolling(
		telega,
		tg.PeasantFixedTxt{Tg: telega, Txt: "Hello there"},
	)
	bot.Run()
}
