package tg

import (
	"encoding/json"
	"fmt"
)

func MakeUpdates(tg Telega) Updates {
	timeout := 10
	return Updates{tg,
		GetUpdatesReq{Offset: nil, Limit: nil, Timeout: &timeout, AllowedUpdates: nil}, nil}
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
	for {
		if len(u.batch) > 0 {
			break
		} else {
			batch, err := u.FetchUpdates(u.req)
			if err != nil {
				return Update{}, nil
			}
			if len(batch) > 0 {
				newOffset := batch[len(batch)-1].UpdateId + 1
				u.req.Offset = &newOffset
			}
			u.batch = batch
		}
	}
	upd := u.batch[0]
	u.batch = u.batch[1:]
	return upd, nil
}

type ApiResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

func (u *Updates) FetchUpdates(req GetUpdatesReq) ([]Update, error) {
	var batch []Update
	payload, err := json.Marshal(req)
	if err != nil {
		return batch, err
	}

	status, body, err := u.tg.Call("getUpdates", payload)

	if err == nil {
		if status == 200 {
			resp := ApiResponse{}
			err = json.Unmarshal(body, &resp)
			batch = resp.Result
		} else {
			err = fmt.Errorf(
				"got status %d getting updates. offset: %d", status, *req.Offset,
			)
		}
	}

	return batch, err
}
