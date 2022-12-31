package main

import (
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	secret := os.Getenv("LINE_CHANNEL_SECRET")

	token := os.Getenv("LINE_CHANNEL_TOKEN")

	bot, err := linebot.New(secret, token)
	if err != nil {
		panic(err)
	}

	if _, err := bot.BroadcastMessage(linebot.NewTextMessage("テストです。")).Do(); err != nil {
		panic(err)
	}
}
