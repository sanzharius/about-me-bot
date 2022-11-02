package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	My_token string `env:"MY_TOKEN"`
	Port     string `env:"PORT"`
}

func Init() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	My_token := os.Getenv("MY_TOKEN")

	fmt.Println(My_token)

	Port := os.Getenv("PORT")

	fmt.Println(Port)

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
	return &cfg, nil

}
