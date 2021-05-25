// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chiahsin's Line Bot.\n")
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg := strings.ToLower(message.Text)
				if strings.Contains(msg, "成績") {
					if err := handleGradeInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "grade") {
					if err := handleGradeInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "學歷") {
					if err := handleSchoolInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "education") {
					if err := handleSchoolInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "履歷") {
					if err := handleResumeInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "resume") {
					if err := handleResumeInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "art") {
					if err := handleArtsCenter(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "藝術中心") {
					if err := handleArtsCenter(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "專案1") {
					if err := handleArtsCenter(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "project1") {
					if err := handleArtsCenter(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "學校") {
					if err := handleSchoolInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "大學") {
					if err := handleSchoolInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "研究所") {
					if err := handleSchoolInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "碩士") {
					if err := handleSchoolInformation(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "streaming analytics system") {
					if err := handleStreamingAnalyticsSystem(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "專案2") {
					if err := handleStreamingAnalyticsSystem(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "project2") {
					if err := handleStreamingAnalyticsSystem(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "專案3") {
					if err := handleCloudVRGaming(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "project3") {
					if err := handleCloudVRGaming(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if strings.Contains(msg, "cloud vr gaming platform") {
					if err := handleCloudVRGaming(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if msg == "1" {
					if err := handleArtsCenter(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if msg == "2" {
					if err := handleStreamingAnalyticsSystem(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else if msg == "3" {
					if err := handleCloudVRGaming(event.ReplyToken); err != nil {
						log.Print(err)
					}
				} else {
					if err := handleUnknownMessage(event.ReplyToken); err != nil {
						log.Print(err)
					}
				}
			default:
				if err := handleUnknownMessage(event.ReplyToken); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func handleUnknownMessage(replyToken string) error {
	replyMessage := "抱歉，目前沒有此關鍵字的設定！可以點選選單或輸入以下幾組關鍵字獲取Chia-Hsin的資料:\n一、輸入「學歷」、「成績」、「履歷」獲取相關資訊\n二、輸入「1」、「2」、「3」獲取Chia-Hsin過去的專案資訊"
	if _, err := bot.ReplyMessage(replyToken,
		linebot.NewStickerMessage("2", "38"),
		linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		return err
	}
	return nil
}

func handleSchoolInformation(replyToken string) error {
	replyMessage := "Chia-Hsin目前是國立清華大學資訊工程學系大四的學生，即將於2021年6月畢業，畢業後會前往國立臺灣大學資訊工程所攻讀碩士。"
	if _, err := bot.ReplyMessage(replyToken,
		linebot.NewTextMessage(replyMessage),
		linebot.NewStickerMessage("2", "30")).Do(); err != nil {
		return err
	}
	return nil
}

func handleStreamingAnalyticsSystem(replyToken string) error {
	replyMessage := "這是我在大二下加入多媒體與網路系統實驗室時，與光寶科技合作的專案，目的是架設一個software defined的串流分析系統，並將此系統deploy到他們的edge設備上(智慧路燈)。"
	replyMessage2 := "我們實作了4個應用並使用Kubernentes去進行管理:\n1.Object Detection\n2.Illegal Parking\n3.People Counting\n4.Traffic Flow Detection"
	replyMessage3 := "最後將這些應用的結果顯示在Web UI上供使用者觀看，我們也將數據寫入Grafana中來加以監控。"
	if _, err := bot.ReplyMessage(replyToken, 
		linebot.NewImageMessage("https://i.imgur.com/iYaPDEB.jpg", "https://i.imgur.com/2Rz6ipct.jpg"), 
		linebot.NewTextMessage(replyMessage),
		linebot.NewTextMessage(replyMessage2),
		linebot.NewImageMessage("https://i.imgur.com/kZ18ke4.jpg", "https://i.imgur.com/CLNaCcOt.jpg"),
		linebot.NewTextMessage(replyMessage3)).Do(); err != nil {
		return err
	}
	return nil
}

func handleGradeInformation(replyToken string) error {
	if _, err := bot.ReplyMessage(replyToken, 
		linebot.NewImageMessage(os.Getenv("grade1_original"), os.Getenv("grade1_previous")), 
		linebot.NewImageMessage(os.Getenv("grade2_original"), os.Getenv("grade2_previous"))).Do(); err != nil {
		return err
	}
	return nil
}

func handleResumeInformation(replyToken string) error {
	if _, err := bot.ReplyMessage(replyToken, 
		linebot.NewImageMessage(os.Getenv("resume_original"), os.Getenv("resume_previous"))).Do(); err != nil {
		return err
	}
	return nil
}

func handleCloudVRGaming(replyToken string) error {
	replyMessage := "這是我大三下做的專題，我們嘗試去搭建Cloud VR的遊戲平台，並對這個平台進行performance的測量，更深入地了解不同網路狀況時，server跟client的效能會有何影響，使用者的user experience又會有何變化。"
	if _, err := bot.ReplyMessage(replyToken, 
		linebot.NewImageMessage("https://i.imgur.com/4rRXxOx.jpg", "https://i.imgur.com/o8QaFHrt.jpg"), 
		linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		return err
	}
	return nil
}

func handleArtsCenter(replyToken string) error {
	replyMessage := fmt.Sprintf("在大一下時，我加入了國立清華大學藝術中心的網管組，負責幫忙藝術中心搭建新網站，server的部份我們使用了Node.js與Nginx，前端的部分則是使用Vue.js去實作RWD UI。")
	replyMessage2 := fmt.Sprintf("並將一些監控數據(如連線人數)打進grafana中做監控，若有發生危險的情況(如連線人數過多)時，則會發送alert給我們。")
	if _, err := bot.ReplyMessage(replyToken, 
		linebot.NewImageMessage("https://i.imgur.com/Ocelkd3.jpg", "https://i.imgur.com/CuM80Jgt.jpg"), 
		linebot.NewTextMessage(replyMessage),
		linebot.NewImageMessage("https://i.imgur.com/rKHWweU.png", "https://i.imgur.com/rKHWweUt.png"), 
		linebot.NewTextMessage(replyMessage2)).Do(); err != nil {
		return err
	}
	return nil
}
