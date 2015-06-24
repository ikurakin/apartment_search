package parser

import (
	"reflect"
	"testing"
)

func TestGetAdvertList(t *testing.T) {
	p := New()

	ids1 := make(map[string]string)
	go p.GetAdvertList()
	var urls1 []string
	urls1 = <-p.Urls
	if len(urls1) == 0 {
		t.Errorf("Can't get adverts from %s", URL_ROOT)
		return
	}

	ids2 := p.GetMsgIds()
	if eq := reflect.DeepEqual(ids1, ids2); eq {
		t.Error("Method doesn't keep the message ids history")
	}

	go p.GetAdvertList()
	var urls2 []string
	urls2 = <-p.Urls
	if len(urls2) > 0 {
		t.Error("Method gets the same ids again")
	}
}

func TestProcessUrls(t *testing.T) {
	p := New()
	go p.GetAdvertList()
	go p.ProcessUrls(<-p.Urls)
	var r string
	r = <-p.AllAdv
	if r == "" {
		t.Error("Can't get adverts from processed urls")
	}
}
