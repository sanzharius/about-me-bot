package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/🇺🇸 US"),
		tgbotapi.NewKeyboardButton("/🇬🇧 GB"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/🇩🇪 DE"),
		tgbotapi.NewKeyboardButton("/🇯🇵 JP"),
	),
)

var countrycode = map[string]string{"US": "US", "GB": "GB", "DE": "DE", "JP": "JP"}

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

		switch update.Message.Text {
		case "/start":
			msg.ReplyMarkup = numericKeyboard
		case "/close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		txt := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("No holidays today in %s", update.Message.Text))
		txt2 := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Holidays celebrated in %s today are:", update.Message.Text))

		switch update.Message.Text {
		case numericKeyboard.Keyboard[0][0].Text:
			holidays, _ := MakeRequest(update.Message.Text)
			fmt.Printf("message: %s\n", holidays)
			if _, err := bot.Send(txt2); err != nil {
				log.Panic(err)
			} else if len(holidays) == 0 {
				fmt.Printf("message: %s\n", update.Message.Text)
				if _, err := bot.Send(txt); err != nil {
					log.Panic(err)
				}
			}
		case numericKeyboard.Keyboard[0][1].Text:
			holidays, _ := MakeRequest(update.Message.Text)
			fmt.Printf("message: %s\n", holidays)
			if _, err := bot.Send(txt2); err != nil {
				log.Panic(err)
			} else if len(holidays) == 0 {
				fmt.Printf("message: %s\n", update.Message.Text)
				if _, err := bot.Send(txt); err != nil {
					log.Panic(err)
				}
			}
		case numericKeyboard.Keyboard[1][0].Text:
			holidays, _ := MakeRequest(update.Message.Text)
			fmt.Printf("message: %s\n", holidays)
			if _, err := bot.Send(txt2); err != nil {
				log.Panic(err)
			} else if len(holidays) == 0 {
				fmt.Printf("message: %s\n", update.Message.Text)
				if _, err := bot.Send(txt); err != nil {
					log.Panic(err)
				}
			}
		case numericKeyboard.Keyboard[1][1].Text:
			holidays, _ := MakeRequest(update.Message.Text)
			fmt.Printf("message: %s\n", holidays)
			if _, err := bot.Send(txt2); err != nil {
				log.Panic(err)
			} else if len(holidays) == 0 {
				fmt.Printf("message: %s\n", update.Message.Text)
				if _, err := bot.Send(txt); err != nil {
					log.Panic(err)
				}
			}

		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

func MakeRequest(country string) ([]Holiday, error) {
	cfg, err := Init()
	if err != nil {
		log.Fatal(err)
	}

	year, month, day := time.Now().Date()
	r := url.Values{}
	r.Add("country", countrycode[country])
	r.Add("api_key", cfg.HolidayApiKey)
	r.Add("day", strconv.Itoa(day))
	r.Add("month", strconv.Itoa(int(month)))
	r.Add("year", strconv.Itoa(year))

	resp, err := http.Get(cfg.HolidayApiHost)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

	var HolUpdate []Holiday
	err = json.Unmarshal(body, &HolUpdate)
	if err != nil {
		return nil, fmt.Errorf("сouldn't unmarshall to struct: %w", err)
	}
	return HolUpdate, nil

}
