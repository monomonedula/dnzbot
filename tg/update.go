package tg

type Update struct {
	UpdateId      int32         `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

type CallbackQuery struct {
	Id              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageId string  `json:"inline_message_id"`
	ChatInstance    string  `json:"chat_instance"`
	Data            string  `json:"data"`
	GameShortName   string  `json:"game_short_name"`
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
