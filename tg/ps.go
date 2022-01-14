package tg

import (
	"encoding/json"
	"regexp"
)

type Fork interface {
	route(Update) (bool, error)
}

type PsFork struct {
	forks []Fork
}

func (ps PsFork) serve(upd Update) error {
	for _, fk := range ps.forks {
		routed, err := fk.route(upd)
		if err != nil {
			return err
		}
		if routed {
			break
		}
	}
	return nil
}

type FkPattern struct {
	pattern regexp.Regexp
	ps      Peasant
}

func (fk FkPattern) route(upd Update) (bool, error) {
	if fk.pattern.MatchString(upd.Message.Text) {
		return true, fk.ps.serve(upd)
	}
	return false, nil
}

type PeasantFixedTxt struct {
	Tg  Telega
	Txt string
}

func (ps PeasantFixedTxt) serve(upd Update) error {
	msg, err := json.Marshal(
		SendMessage{
			upd.Message.From.Id,
			ps.Txt,
		},
	)
	if err == nil {
		_, err = ps.Tg.Call_(
			"sendMessage",
			msg,
		)
	}
	return err
}

type SendMessage struct {
	ChatId int32  `json:"chat_id"`
	Text   string `json:"text"`
}
