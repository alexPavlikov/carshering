package main

import (
	"fmt"
	"net/http"

	"github.com/alexPavlikov/scud.ru/algorithms"
	"github.com/alexPavlikov/scud.ru/config"
	"github.com/alexPavlikov/scud.ru/handler"
	"github.com/alexPavlikov/scud.ru/models"
)

func init() {
	config.Config()
}

func main() {
	fmt.Println("Hello world!")

	handler.Handler()

	err := http.ListenAndServe(models.Config.ServerPort, nil)
	algorithms.CheckErr(err)
}
