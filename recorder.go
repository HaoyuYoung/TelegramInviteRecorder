package TelegramInviteRecorder

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/sheets/v4"
	"log"
	"strconv"
)

func InviteRecorder(bot *tgbotapi.BotAPI, googleClient *sheets.Service, sheetID string) {
	cell := make(map[string]string)
	num := make(map[string]int)
	count := 1
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)

	//Get Updates from your bot
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		//record the invitor
		if len(update.Message.NewChatMembers) != 0 {
			if update.Message.From == &update.Message.NewChatMembers[0] {
				fmt.Println("Joined")
			} else if update.Message.From.UserName != update.Message.NewChatMembers[0].UserName {
				a := cell[update.Message.From.UserName]
				if a == "" {
					fmt.Println(update.Message.From.UserName)
					cell[update.Message.From.UserName] = "B" + strconv.FormatInt(int64(count), 10)
					num[update.Message.From.UserName] = 1
					count = count + 1
					interfaceArray := [][]interface{}{{update.Message.From.UserName, len(update.Message.NewChatMembers)}}
					vr := sheets.ValueRange{
						MajorDimension:  "ROWS",
						Range:           "A1:A2",
						Values:          interfaceArray,
						ServerResponse:  googleapi.ServerResponse{},
						ForceSendFields: nil,
						NullFields:      nil,
					}
					_, err := googleClient.Spreadsheets.Values.Append(sheetID, "A1:A2", &vr).ValueInputOption("RAW").Do()
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					num[update.Message.From.UserName] = num[update.Message.From.UserName] + len(update.Message.NewChatMembers)
					interfaceArray := [][]interface{}{{num[update.Message.From.UserName]}}
					vr := sheets.ValueRange{
						MajorDimension:  "DIMENSION_UNSPECIFIED",
						Range:           cell[update.Message.From.UserName],
						Values:          interfaceArray,
						ServerResponse:  googleapi.ServerResponse{},
						ForceSendFields: nil,
						NullFields:      nil,
					}
					_, err := googleClient.Spreadsheets.Values.Update(sheetID, cell[update.Message.From.UserName], &vr).ValueInputOption("RAW").Do()
					if err != nil {
						fmt.Println(err)
						return
					}
				}

			}
		}

	}
}
