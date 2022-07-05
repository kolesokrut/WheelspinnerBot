package internal

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

type app struct {
	cfg string
}

type App interface {
	Run()
}

func NewApp(cfg string) (App, error) {
	return &app{
		cfg: cfg,
	}, nil
}

func (a *app) Run() {
	a.startBot()
}

func (a *app) startBot() {
	pref := tele.Settings{
		Token:  a.cfg,
		Poller: &tele.LongPoller{Timeout: 60 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
