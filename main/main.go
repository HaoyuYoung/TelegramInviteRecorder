package main

import (
	"TelegramInviteRecorder"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	//BotToken String
	BotToken = "YOUR-BOT-TOKEN"
	//ChatID uint
	ChatID = "YOUR-CHART-ID"
	//Your can get these two from BotFather and RawDataBot

	GoogleAPI = "YOUR-API-JSON-PATH"
	//How to create: https://developers.google.com/sheets/api/guides/concepts?hl=en

	Gmail = "YOUR-GMAIL-to-READ-DATA"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}

	id := TelegramInviteRecorder.NewSheet(GoogleAPI, "YOUR-SHEET-NAME")
	TelegramInviteRecorder.AddEditor(GoogleAPI, id, Gmail)
	ctx := context.Background()
	client, err := sheets.NewService(ctx, option.WithCredentialsFile(GoogleAPI))
	if err != nil {
		fmt.Println(err)
	}
	go TelegramInviteRecorder.InviteRecorder(bot, client, id)
}
