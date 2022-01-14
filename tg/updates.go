package tg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func MakeUpdates(tg Telega) Updates {
	return Updates{tg, GetUpdatesReq{}, nil}
}

type Updates struct {
	tg    Telega
	req   GetUpdatesReq
	batch []Update
}

type GetUpdatesReq struct {
	Offset         *int32    `json:"offset"`
	Limit          *int      `json:"limit"`
	Timeout        *int      `json:"timeout"`
	AllowedUpdates *[]string `json:"allowed_updates"`
}

func (u *Updates) NextUpdate() (Update, error) {
	if len(u.batch) == 0 {
		batch, err := u.FetchUpdates(u.req)
		msg, _ := json.MarshalIndent(batch, "", "\t")
		log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile).Println(
			err, "Fetched updates: ", msg,
		)
		if err != nil {
			return Update{}, nil
		}
		u.batch = batch
	}

	upd := u.batch[0]
	u.batch = u.batch[1:]
	return upd, nil
}

type ApiResponse struct {
	Ok     bool        `json:"ok"`
	Result interface{} `json:"result"`
}

func (u *Updates) FetchUpdates(req GetUpdatesReq) ([]Update, error) {
	var batch []Update
	payload, err := json.Marshal(req)
	if err != nil {
		return batch, err
	}

	status, body, err := u.tg.Call("getUpdates", payload)

	log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile).Println(
		err, "Fetched : ", string(status), string(body),
	)

	if err == nil {
		if status == 200 {
			var batch []Update
			resp := ApiResponse{Ok: false, Result: &batch}
			err = json.Unmarshal(body, &resp)
		} else {
			err = fmt.Errorf(
				"got status %d getting updates. offset: %d", status, *req.Offset,
			)
		}
	}

	return batch, err
}
