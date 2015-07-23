package main

import (
	"log"
	"parser"
	"sendmsg"
	"time"
)

func main() {
	log.Println("Start apartment adverts search")
	p := parser.New()
	go p.GetAdvertList()
	for {
		select {
		case r := <-p.AllAdv:
			sendmsg.ToSlack(r)
		case u := <-p.Urls:
			go p.ProcessUrls(u)
		case <-time.After(120 * time.Minute):
			go p.GetAdvertList()
		}

	}
}
