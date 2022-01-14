package tg

import (
	"encoding/json"
	"fmt"
	"log"
)

func PsForkOf(forks ...Fork) PsFork {
	return PsFork{Forks: forks}
}

type PsFork struct {
	Forks []Fork
}

func (ps PsFork) serve(upd Update) error {
	for _, fk := range ps.Forks {
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

type PsFixedTxt struct {
	Tg  Telega
	Txt string
}

func (ps PsFixedTxt) serve(upd Update) error {
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

type PsErrorCatching struct {
	Log log.Logger
	Ps  Peasant
}

func (ps PsErrorCatching) serve(upd Update) error {
	err := ps.Ps.serve(upd)

	if err != nil {
		msg, _ := json.Marshal(upd)
		ps.Log.Println(
			fmt.Sprintf("Got error while handling %s :", msg), err,
		)
	}
	return nil
}
