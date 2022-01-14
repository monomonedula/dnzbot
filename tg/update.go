package tg

type Update struct {
	UpdateId int32   `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessasgeId int32    `json:"message_id"`
	From       User     `json:"from"`
	Text       string   `json:"text"`
	Location   Location `json:"location"`
}

type Location struct {
	Longitude float32 `json:"longtitude"`
	Latitude  float32 `json:"latitude"`
}

type User struct {
	Id        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Lang      string `json:"language_code"`
}
