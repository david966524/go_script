package main

import (
	"log"
	"mybot/aliyunapi"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  "",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle("/bill", func(ctx tele.Context) error {
		return ctx.Send(aliyunapi.DailyBill())
	})

	b.Handle("/banlance", func(ctx tele.Context) error {
		return ctx.Send(aliyunapi.AccoutBanlance())
	})

	b.Start()
}
