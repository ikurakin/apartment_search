package sendmsg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SlackMsg struct {
	Channel   string `json:"channel"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
}

func SendToSlack(m string) {
	s := &SlackMsg{
		Channel:   "#adv_msg",
		Text:      m,
		IconEmoji: ":ghost:",
	}

	jsonBytes, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
		return
	}
	b := bytes.NewBuffer(jsonBytes)
	r, err := http.Post("https://hooks.slack.com/services/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", "application/json", b)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(r.Status)
	log.Println(string(body))
}
