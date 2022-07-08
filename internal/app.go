package internal

import (
	"github.com/kolesokrut/WheelspinnerBot/internal/client/banking"
	"github.com/kolesokrut/WheelspinnerBot/internal/client/soundcloud"
	"github.com/kolesokrut/WheelspinnerBot/internal/client/tiktok"
	"github.com/kolesokrut/WheelspinnerBot/internal/client/youtube"
	t "gopkg.in/telebot.v3"
	"log"
	"net/url"
	"time"
)

type app struct {
	cfg string
	msg string
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

var (
	// Universal markup builders.
	menu = &t.ReplyMarkup{ResizeKeyboard: true}

	btnWeather  = menu.Text("Погода")
	btnCurrency = menu.Text("Валюта")
)

func (a *app) startBot() {
	pref := t.Settings{
		Token:   a.cfg,
		Poller:  &t.LongPoller{Timeout: 60 * time.Second},
		Verbose: false,
		OnError: a.OnBotError,
	}

	b, err := t.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	menu.Reply(
		menu.Row(btnWeather),
		menu.Row(btnCurrency),
	)

	b.Handle(&btnCurrency, func(c t.Context) error {
		return c.Send(banking.Banks(), &t.SendOptions{ParseMode: t.ModeMarkdown})
	})

	b.Handle(&btnWeather, func(c t.Context) error {
		return c.Send(banking.Weather(""))
	})

	b.Handle(t.OnText, func(c t.Context) error {
		var text = c.Text()

		u, err := url.Parse(text)
		if err != nil {
			return err
		}

		if u.Host == "youtube.com" || u.Host == "youtu.be" {
			a.msg, _ = youtube.DownloadMP3(text)
			c.Send(&t.Audio{File: t.FromURL(a.msg)})
		}

		if u.Host == "tiktok.com" || u.Host == "vm.tiktok.com" {
			a.msg = tiktok.DownloadVideo(text)
			c.Send(&t.Audio{File: t.FromURL(a.msg)})
		}

		if u.Host == "soundcloud.com" || u.Host == "soundcloud.app.goo.gl" {
			a.msg = soundcloud.DownloadMusic(text)
			c.Send(&t.Audio{File: t.FromURL(a.msg)})
		}

		return nil
	})

	b.Handle("/start", func(c t.Context) error {
		return c.Send("Hello!", menu)
	})

	b.Handle("/youtube", func(c t.Context) error {
		return c.Send("give youtube link")
	})

	b.Handle("/tiktok", func(c t.Context) error {
		return c.Send("give tiktok link")
	})

	b.Handle("/soundcloud", func(c t.Context) error {
		return c.Send("give soundcloud link")
	})

	b.Start()
}

func (a *app) OnBotError(err error, ctx t.Context) {
	log.Println(err)
}
