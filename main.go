package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/codenoid/minikv"
	tele "gopkg.in/tucnak/telebot.v3"
)

var (
	bot *tele.Bot

	db = minikv.New(15*time.Minute, 5*time.Second)
)

func main() {
	botToken := os.Getenv("TG_BOT_TOKEN")

	// listen for janitor expiration removal ( 5*time.Second )
	db.OnEvicted(onEvicted)

	b, err := tele.NewBot(tele.Settings{
		Token:  botToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot = b

	b.Handle("/testcaptcha", onJoin)
	b.Handle(tele.OnUserJoined, onJoin)
	b.Handle(tele.OnCallback, handleAnswer)
	b.Handle(tele.OnUserLeft, func(c tele.Context) error {
		kvID := fmt.Sprintf("%v-%v", c.Sender().ID, c.Chat().ID)
		if statusObj, found := db.Get(kvID); found {
			status := statusObj.(JoinStatus)
			bot.Delete(&status.CaptchaMessage)
			db.Delete(kvID)
		}
		c.Delete()
		return nil
	})

	b.Start()
}
