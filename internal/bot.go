package internal

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

func StartBot() {
	bot := initBot()
	CallbackHandler(bot)
}

func initBot() *linebot.Client{
	loadEnv()
	bot, err := linebot.New(os.Getenv("LINE_CHANNEL_SECRET"), os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal("Error starting Line Bot")
	}
	return bot
}


func CallbackHandler(client *linebot.Client) {
	http.HandleFunc("/callback", func(writer http.ResponseWriter, request *http.Request) {
		events, err := client.ParseRequest(request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				writer.WriteHeader(http.StatusBadRequest)
			} else {
				writer.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeBeacon:
				message := linebot.NewTextMessage("Test")
				_, err := client.ReplyMessage(event.ReplyToken, message).Do()
				if err != nil {
					log.Println(err)
					continue
				}
			}
		}
	})
}