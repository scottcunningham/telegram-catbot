package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

const helpMessage = `🐱 Options: 🐈‍⬛

/cat -- 🤗 show me a 🐱
/catfact -- 🤔 tell me about 🐈

CatBot respects 🦮 dogs 🐩 too! Just a little bit less than cats.
Try these dog commands too:

/dog -- 🙄 show me a 🐕
/dotfact -- 🤭 tell me about 🐶
`

var telegramToken = os.Getenv("TELEGRAM_TOKEN")

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  telegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/help", func(m *tb.Message) {
		b.Send(m.Chat, helpMessage)
	})

	b.Handle("/cat", func(m *tb.Message) {
		catURL, err := getCatURL()
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		b.Send(m.Chat, catURL)
	})

	b.Handle("/dog", func(m *tb.Message) {
		dogURL, err := getDogURL()
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		b.Send(m.Chat, dogURL)
	})

	b.Handle("/catfact", func(m *tb.Message) {
		facts, err := getCatFacts()
		if err != nil {
			panic(err)
		}
		b.Send(m.Chat, facts.RandomFact().Text)
	})

	b.Handle("/dogfact", func(m *tb.Message) {
		fact, err := getDogFact()
		if err != nil {
			panic(err)
		}
		b.Send(m.Chat, fact)
	})

	b.Start()
}
