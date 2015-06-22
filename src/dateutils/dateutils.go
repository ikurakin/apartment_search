package dateutils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	DATE_FMT = "2 Jan 2006, 15:04"
)

var (
	MONTHS = map[string]string{
		"апр.": "Apr", "мая": "May", "июня": "Jun", "июля": "Jul",
	}
)

func FmtDate(d string) (time.Time, error) {
	d_split := strings.Split(d, " ")
	m, ok := MONTHS[d_split[1]]
	if ok {
		new_d := strings.Replace(d, d_split[1], m, -1)
		loc, _ := time.LoadLocation("Europe/Chisinau")
		t, err := time.ParseInLocation(DATE_FMT, new_d, loc)
		if err != nil {
			log.Println(err)
		}
		return t, err
	}
	err := fmt.Errorf("%s not in MONTHS", d_split[1])
	log.Println(err)
	return time.Time{}, err
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
