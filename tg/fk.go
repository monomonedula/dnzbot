package tg

import "regexp"

type Fork interface {
	route(Update) (bool, error)
}

type FkMulti struct {
	Forks []Fork
}

func (fk FkMulti) route(upd Update) (bool, error) {
	for _, fk := range fk.Forks {
		routed, err := fk.route(upd)
		if err != nil {
			return true, err
		}
		if routed {
			break
		}
	}
	return true, nil
}

type FkPattern struct {
	Pattern regexp.Regexp
	Ps      Peasant
}

func (fk FkPattern) route(upd Update) (bool, error) {
	if fk.Pattern.MatchString(upd.Message.Text) {
		return true, fk.Ps.serve(upd)
	}
	return false, nil
}

type UpdType string

const (
	UtMessage            UpdType = "message"
	UtEditedMessage      UpdType = "edited_message"
	UtChannelPost        UpdType = "channel_post"
	UtEditedChannelPost  UpdType = "edited_channel_post"
	UtInluneQuery        UpdType = "inline_query"
	UtChosenInlineResult UpdType = "chosen_inline_result"
	UtCallbackQuery      UpdType = "callback_query"
	UtShippingQuery      UpdType = "shipping_query"
	UtPreCheckoutQuery   UpdType = "pre_checkout_query"
	UtPoll               UpdType = "poll"
	UtPollAnswer         UpdType = "poll_answer"
	UtMyChatMember       UpdType = "my_chat_member"
	UtChatMember         UpdType = "chat_member"
	UtChatJoinRequest    UpdType = "chat_join_request"
)

type FkUpdType struct {
	Type UpdType
	Ps   Peasant
}

func (fk FkUpdType) route(upd Update) (bool, error) {
	match := false
	switch fk.Type {
	case UtMessage:
		match = upd.Message != (Message{})
	}
	if match {
		return true, fk.Ps.serve(upd)
	}
	return false, nil
}
