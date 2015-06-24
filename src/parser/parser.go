package parser

import (
	"dateutils"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL_ROOT = "https://999.md"
)

var (
	URL_MSGS = fmt.Sprintf("%s/list/real-estate/apartments-and-rooms?applied=1&o_30_241=894&r_31_2_unit=eur&view_type=1&o_33_1=912&r_31_2_to=250&r_31_2_from=&o_1074_253=916&o_1074_253=937&o_32_9_12900_13859=13859&o_32_9_12900_13859=15665", URL_ROOT)
)

type ApartmentAdvert struct {
	sync.Mutex
	msgIds map[string]string
	Urls   chan []string
	Adv    chan string
	AllAdv chan string
}

func New() *ApartmentAdvert {
	aa := &ApartmentAdvert{
		msgIds: make(map[string]string),
		Urls:   make(chan []string),
		Adv:    make(chan string),
		AllAdv: make(chan string),
	}
	return aa
}

func (aa *ApartmentAdvert) GetMsgIds() map[string]string {
	return aa.msgIds
}

func (aa *ApartmentAdvert) GetAdvertList() {
	doc, err := goquery.NewDocument(URL_MSGS)
	if err != nil {
		log.Println(err)
	}
	now := time.Now()
	day_after := time.Unix(now.Unix()-86400, 0)
	var msgs []string
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {

		d, err := dateutils.FmtDate(s.Find(".ads-list-table-date").Text())
		if err != nil {
			return
		}
		if !dateutils.InTimeSpan(day_after, now, d) {
			return
		}
		p := strings.TrimSpace(s.Find(".ads-list-table-price").Text())
		if p == "" {
			return
		}
		link, _ := s.Find("h3 a").Attr("href")
		if _, ok := aa.msgIds[link]; ok {
			return
		}
		aa.Lock()
		aa.msgIds[link] = ""
		aa.Unlock()
		msgs = append(msgs, fmt.Sprintf("%s%s", URL_ROOT, link))
	})
	aa.Urls <- msgs
}

func (aa *ApartmentAdvert) getAdvertData(u string) {
	doc, err := goquery.NewDocument(u)
	if err != nil {
		log.Println(err)
	}
	s := doc.Find(".adPage")

	title := strings.TrimSpace(s.Find("h1").Text())
	price := strings.TrimSpace(s.Find(".adPage__content__price").Find("dd").Text())
	location := strings.TrimSpace(s.Find(".adPage__content__region").Text())
	phone := strings.TrimSpace(s.Find(".adPage__content__phone").Text())
	date := strings.TrimSpace(s.Find(".adPage__header__stats__date").Next().Text())
	delim := "---------------------------------------------"
	result := strings.Join([]string{
		fmt.Sprintf("Title: %s", title), fmt.Sprintf("Price: %s", price),
		strings.Replace(location, "\n", "", -1), strings.Replace(phone, "\n", "", -1),
		fmt.Sprintf("Date: %s", date), fmt.Sprintf("URL: <%s>", u), delim, "\n"}, "\n")
	aa.Adv <- result
}

func (aa *ApartmentAdvert) ProcessUrls(urls []string) {
	for _, url := range urls {
		go aa.getAdvertData(url)
	}
	var result []string
	for i := 0; i < len(urls); i++ {
		result = append(result, <-aa.Adv)
	}
	aa.AllAdv <- strings.Join(result, "\n")
}
