package handle

import (
	tele "gopkg.in/telebot.v3"
)

func Start(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		c.Send("Xin chào! Đây là bot của https://t.me/huenews")
		return nil
	})
}
