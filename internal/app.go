package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kolesokrut/WheelspinnerBot/internal/client/banking"
	"github.com/kolesokrut/WheelspinnerBot/internal/client/soundcloud"
	"github.com/kolesokrut/WheelspinnerBot/internal/client/tiktok"
	"github.com/kolesokrut/WheelspinnerBot/internal/client/youtube"
	"github.com/kolesokrut/WheelspinnerBot/internal/config"
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
	menu = &t.ReplyMarkup{ResizeKeyboard: true}

	btnWeather  = menu.Text("⛅️Погода")
	btnCurrency = menu.Text("🏛Валюта")

	db *sql.DB
)

func (a *app) startBot() {
	cfg, _ := config.LoadConfig("dev")

	db, err := sql.Open(cfg.DB.Driver, fmt.Sprintf("%s:%s@%s(%s:%s)/%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.Protocol, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

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
		row := db.QueryRow("select * from telegramdb.citynames where id = ?", c.Message().Sender.ID)

		cit := banking.City{}
		err = row.Scan(&cit.Id, &cit.Name)
		if err != nil {
			panic(err)
		}

		return c.Send(banking.Weather(cit.Name, cfg.Api.OpenWeather))
	})

	b.Handle(t.OnText, func(c t.Context) error {
		var text = c.Text()

		u, err := url.Parse(text)
		if err != nil {
			return err
		}

		if u.Host == "youtube.com" || u.Host == "youtu.be" {
			a.msg, _ = youtube.DownloadMP3(text, cfg.Api.Youtube)
			c.Send(&t.Audio{File: t.FromURL(a.msg)})
		}

		if u.Host == "tiktok.com" || u.Host == "vm.tiktok.com" {
			a.msg = tiktok.DownloadVideo(text, cfg.Api.Tiktok)
			c.Send(&t.Audio{File: t.FromURL(a.msg)})
		}

		if u.Host == "soundcloud.com" || u.Host == "soundcloud.app.goo.gl" {
			a.msg = soundcloud.DownloadMusic(text, cfg.Api.Soundcloud)
			c.Send(&t.Audio{File: t.FromURL(a.msg)})
		}

		c.Bot().Delete(c.Message())

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

	b.Handle("/setcity", func(c t.Context) error {
		if len(c.Message().Payload) == 0 {
			// переделать
			return c.Send("Введите город после команды")
		}

		cit := banking.City{}
		row := db.QueryRow("select * from telegramdb.citynames where id = ?", c.Message().Sender.ID)

		row.Scan(&cit.Id, &cit.Name)

		if cit.Id == 0 {
			_, err := db.Exec("insert into telegramdb.citynames (id, city) values (?, ?)",
				c.Message().Sender.ID, c.Message().Payload)
			if err != nil {
				panic(err)
			}

			return c.Send(fmt.Sprintf("Вы выбрали город: %s", c.Message().Payload))
		}

		_, err := db.Exec("update telegramdb.citynames set city = ? where id = ?", c.Message().Payload, c.Message().Sender.ID)
		if err != nil {
			panic(err)
		}
		println(c.Message().Payload)

		return c.Send(fmt.Sprintf("Вы выбрали город: %s", c.Message().Payload))
	})

	b.Start()
}

func (a *app) OnBotError(err error, ctx t.Context) {
	log.Println(err)
}
