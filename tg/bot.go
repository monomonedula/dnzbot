package tg

type Bot struct {
	tg      Telega
	updates Updates
	ps      Peasant
}

type Peasant interface {
	serve(Update) error
}

func NewBotPolling(tg Telega, ps Peasant) Bot {
	return Bot{tg, MakeUpdates(tg), ps}
}

func (bot *Bot) Run() error {
	for {
		upd, err := bot.updates.NextUpdate()
		if err != nil {
			return err
		}

		go bot.ps.serve(upd)
	}
}
