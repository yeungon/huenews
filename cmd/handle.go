package cmd

import (
	"github.com/yeungon/huenews/internal/handle"
	tele "gopkg.in/telebot.v3"
)

func Handle(b *tele.Bot) {
	handle.Start(b)
	handle.News(b)
}
