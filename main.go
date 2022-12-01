package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/about"),
		tgbotapi.NewKeyboardButton("/links"),
	),
)

func main() {

	log.WithFields(log.Fields{
		"out":  os.Stderr,
		"time": time.Now(),
	}).Info("A new message received")

	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	LogLevel, err := log.ParseLevel(os.Getenv("LOGLEVEL"))
	if err != nil {
		LogLevel = log.InfoLevel
	}

	log.SetLevel(LogLevel)

	cfg, err := Init()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.MyToken)
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

		switch update.Message.Command() {
		case "/start", "/help":
			msg.ReplyMarkup = numericKeyboard
		case "/close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		txt := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, my name is Sanzhar, I study Go")
		txt2 := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/sanzharius, https://www.linkedin.com/in/sanzhar-umarov-713818252/")

		switch update.Message.Text {
		case numericKeyboard.Keyboard[0][0].Text:
			fmt.Printf("message: %s\n", update.Message.Text)
			if _, err := bot.Send(txt); err != nil {
				log.Panic(err)
			}
		case numericKeyboard.Keyboard[0][1].Text:
			fmt.Printf("message: %s\n", update.Message.Text)
			if _, err := bot.Send(txt2); err != nil {
				log.Panic(err)
			}
		}

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
