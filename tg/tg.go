package tg

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Telega struct {
	baseUrl string
	http    *http.Client
}

func NewTelega(token string) Telega {
	return Telega{
		fmt.Sprintf("https://api.telegram.org/bot%s/", token),
		&http.Client{},
	}
}

func (tg *Telega) Call(method string, payload []byte) (int, []byte, error) {
	var body []byte
	req, err := tg.NewReq(method, payload)
	if err != nil {
		return 0, body, err
	}
	resp, err := tg.http.Do(req)
	if err != nil {
		return 0, body, err
	}
	body, err = io.ReadAll(resp.Body)
	return resp.StatusCode, body, err
}

func (tg *Telega) Call_(method string, payload []byte) ([]byte, error) {
	status, body, err := tg.Call(method, payload)
	if status != 200 {
		err = fmt.Errorf("got status %d calling %s with payload %s", status, method, string(payload))
	}
	return body, err
}

func (tg *Telega) NewReq(method string, payload []byte) (*http.Request, error) {
	url := tg.baseUrl + method
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
