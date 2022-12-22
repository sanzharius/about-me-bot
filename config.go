package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	MyToken        string `env:"MYTOKEN"`
	Port           string `env:"PORT"`
	HolidayApiHost string `env:"HOLIDAYAPIHOST"`
	HolidayApiKey  string `env:"HOLIDAYAPIKEY"`
}

type Holiday struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func Init() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	MyToken := os.Getenv("MYTOKEN")
	fmt.Println(MyToken)

	Port := os.Getenv("PORT")
	fmt.Println(Port)

	HolidayApiHost := os.Getenv("HOLIDAYAPIHOST")
	fmt.Println(HolidayApiHost)

	HolidayApiKey := os.Getenv("HOLIDAYAPIKEY")
	fmt.Println(HolidayApiKey)

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
	return &cfg, nil

}
