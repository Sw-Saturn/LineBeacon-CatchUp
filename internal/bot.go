package internal

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

func StartBot() {
	bot := initBot()
	SetCallbackHandler(bot)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initBot() *linebot.Client {
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal("Error starting Line Bot")
	}
	return bot
}

func SetCallbackHandler(client *linebot.Client) {
	http.HandleFunc("/callback", func(writer http.ResponseWriter, request *http.Request) {
		events, err := client.ParseRequest(request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				writer.WriteHeader(http.StatusBadRequest)
				log.Println(err)
			} else {
				writer.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
			}
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeBeacon:
				replyMessage := linebot.NewTextMessage(
					"あああああああああああああああああああああああああああああああ!!!!!!!!!!!" +
						"(ﾌﾞﾘﾌﾞﾘﾌﾞﾘﾌﾞﾘｭﾘｭﾘｭﾘｭﾘｭﾘｭ!!!!!!ﾌﾞﾂﾁﾁﾌﾞﾌﾞﾌﾞﾁﾁﾁﾁﾌﾞﾘﾘｲﾘﾌﾞﾌﾞﾌﾞﾌﾞｩｩｩｩｯｯｯ!!!!!!!)",
				)
				_, err := client.ReplyMessage(event.ReplyToken, replyMessage).Do()
				if err != nil {
					log.Println(err)
					continue
				}
			case linebot.EventTypeMessage:
				switch incomingMessage := event.Message.(type) {
				case *linebot.TextMessage:
					replyMessage := linebot.NewTextMessage(incomingMessage.Text)
					_, err := client.ReplyMessage(event.ReplyToken, replyMessage).Do()
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}
		}
	})
}
