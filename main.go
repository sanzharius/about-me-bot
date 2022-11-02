package main

import (
	"fmt"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/about"),
		tgbotapi.NewKeyboardButton("/links"),
	),
)

func main() {
	cfg, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv(cfg.My_token))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "/start", "/help":
			msg.ReplyMarkup = numericKeyboard
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		log.Print("starting server...")

		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

		/*KeyboardButtons := tgbotapi.NewKeyboardButton("")

		switch KeyboardButtons {
		case tgbotapi.NewKeyboardButton("/about"):
			fmt.Println("Hi, my name is Sanzhar, I study Go")
		case tgbotapi.NewKeyboardButton("/links"):
			fmt.Println("https://github.com/sanzharius, https://www.linkedin.com/in/sanzhar-umarov-713818252/")
		}*/

	}
}
