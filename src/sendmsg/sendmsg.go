package sendmsg

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type SlackMsg struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
}

func ToSlack(m string) int {
	s := &SlackMsg{
		Channel:   "#adv_msg",
		Username:  "aprtm_adv_bot",
		Text:      m,
		IconEmoji: ":ghost:",
	}

	jsonBytes, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
		return 500
	}
	b := bytes.NewBuffer(jsonBytes)
	r, err := http.Post("https://hooks.slack.com/services/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", "application/json", b)
	if err != nil {
		log.Println(err)
		return 500
	}
	defer r.Body.Close()
	// body, _ := ioutil.ReadAll(r.Body)
	log.Println(r.Status)
	return r.StatusCode
}
