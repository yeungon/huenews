package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/yeungon/huenews/config"
	tele "gopkg.in/telebot.v3"
)

func Init() {
	config.New()
	token := config.Get().Token
	Pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(Pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Hi, your bot is running as expected!")

	Handle(b)
	Schedule()
	b.Start()
}
