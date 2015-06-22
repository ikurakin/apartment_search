package main

import (
	"parser"
	"sendmsg"
	"time"
)

func main() {
	p := parser.New()
	for {
		select {
		case r := <-p.AllAdv:
			sendmsg.SendToSlack(r)
		case u := <-p.Urls:
			go p.ProcessUrls(u)
		case <-time.After(180 * time.Minute):
			go p.GetAdvertList()
		}

	}
}
