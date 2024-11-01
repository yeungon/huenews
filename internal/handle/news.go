package handle

import (
	tele "gopkg.in/telebot.v3"
)

func News(b *tele.Bot) {
	b.Handle("/news", func(c tele.Context) error {
		b.Send(tele.ChatID(-1002257060141), "Đây là Huế News")
		return nil
	})
}
