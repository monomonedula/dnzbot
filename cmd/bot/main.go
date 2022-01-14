package main

import (
	"log"
	"os"
	"regexp"

	"github.com/monomonedula/dnzbot/tg"
)

func apiKey() string {
	token := os.Getenv("API_TOKEN")
	if token == "" {
		panic("API_TOKEN env variable must be set")
	}

	return token
}

func main() {
	telega := tg.NewTelega(apiKey())
	bot := tg.NewBotPolling(
		telega,
		tg.PsForkOf(
			tg.FkUpdType{
				Type: tg.UtMessage,
				Ps: tg.PsForkOf(
					tg.FkPattern{
						Pattern: *regexp.MustCompile(`^/start$`),
						Ps:      tg.PsFixedTxt{Tg: telega, Txt: "hello there 1234"},
					},
					tg.FkPattern{
						Pattern: *regexp.MustCompile(`^/yasla$`),
						Ps:      tg.PsFixedTxt{Tg: telega, Txt: "hello there 1234"},
					},

					tg.FkPattern{
						Pattern: *regexp.MustCompile(`^/molodsha$`),
						Ps:      tg.PsFixedTxt{Tg: telega, Txt: "hello there 1234"},
					},

					tg.FkPattern{
						Pattern: *regexp.MustCompile(`^/serednia$`),
						Ps:      tg.PsFixedTxt{Tg: telega, Txt: "hello there 1234"},
					},

					tg.FkPattern{
						Pattern: *regexp.MustCompile(`^/\d+$`),
						Ps:      tg.PsFixedTxt{Tg: telega, Txt: "hello there 1234"},
					},
				),
			},
		),
		*log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	)
	bot.Run()
}
