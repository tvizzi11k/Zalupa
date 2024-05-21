package bot

import (
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func Run() {
	pref := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	menu := b.NewMarkup()
	btn := menu.WebApp("Site", &telebot.WebApp{URL: "https://176-99-11-185.cloudvps.regruhosting.ru/"})

	menu.Inline(menu.Row(btn))

	b.Handle("/start", func(c telebot.Context) error {
		return c.Send("Click to button below:", menu)
	})

	b.Start()
}
