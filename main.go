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
	//"net/url"
	//"io/ioutil"

	"github.com/line/line-bot-sdk-go/linebot"
	//"github.com/gorilla/sessions"
)

var bot *linebot.Client

func main() {
	var err error
	var channelSecret = os.Getenv("ChannelSecret")
    var channelAccessToken =  os.Getenv("ChannelAccessToken")
	bot, err = linebot.New(channelSecret, channelAccessToken)
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)


    /*resp, err := http.Get("https://hr.hiyes.tw:443/getMessage.php?mid=Kordan&message=Ou")
    if err != nil {
        fmt.Println(err)
    } else {
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println("GET OK: ", string(body), resp)
    }*/







	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
    var msg string
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(event.ReplyToken+":"+message.ID+"-"+message.Text+" OK!")).Do(); err != nil {
				if message.Text == "熱銷"{
					msg = "勤耕延吉 · 電話：(02)25705777 · 地址：台北市松山區光復南路58巷 http://www.hiyes.tw/allcase/yanji/index.html "					
					msg = msg + "幸福莊園 · 電話：02-2678-7222 · 地址：新北市鶯歌區鳳福路及鳳鳴路口 · 接待地址：新北市鶯歌區鶯歌路及鳳鳴路口 http://www.hiyes.tw/allcase/happymanor/index.html "
					//msg = "http://www.hiyes.tw/allcase/yanji/index.html "
					//msg = msg + "http://www.hiyes.tw/allcase/happymanor/index.html"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				}
				if message.Text == "預約"{
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您想預約的日期")).Do(); err != nil {
						log.Print(err)
					}					
				}
				if strings.Contains(message.Text, "月") || strings.Contains(message.Text, "日"){
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您想預約的日期是" + message.Text + ",您想預約的時間?")).Do(); err != nil {
						log.Print(err)
					}
				}
				if strings.Contains(message.Text, "點"){
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你想預約的時間是" + message.Text + " , 您的預約已建立,會有專人與您聯絡!")).Do(); err != nil {
						log.Print(err)
					}
				}
				if message.Text == "晚上"{
					msg = "現在是下班時間,造成你的不便敬請原諒!"				
				}

				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好,請輸入你想要詢問的文字,如 : 熱銷,預約")).Do(); err != nil {
						log.Print(err)
				}
			}
		}
	}
}
