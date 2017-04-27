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
	//"net/url"
	//"io/ioutil"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
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
				if message.Text == "勤耕延吉"{
					msg = "勤耕延吉 http://www.hiyes.tw/allcase/yanji/index.html"				
				}
				if message.Text == "幸福莊園"{
					msg = "幸福莊園 http://www.hiyes.tw/allcase/happymanor/index.html"				
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
