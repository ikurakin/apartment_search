package dateutils

import (
	"testing"
	"time"
)

func TestFmtDate(t *testing.T) {
	date := "21 апр. 2014, 00:59"
	loc, _ := time.LoadLocation("Europe/Chisinau")
	expected_date := time.Date(2014, time.April, 21, 0, 59, 0, 0, loc)
	result_date, err := FmtDate(date)
	if err != nil {
		t.Error(err)
	}
	if !expected_date.Equal(result_date) {
		t.Errorf("expected date %s not equal to result date %s", expected_date.String(), result_date.String())
	}
}

func TestInTimeSpan(t *testing.T) {
	now := time.Now()
	day_after := time.Unix(now.Unix()-86400, 0)
	check_date := time.Unix(now.Unix()-43200, 0)
	if !InTimeSpan(day_after, now, check_date) {
		t.Errorf("%s not between %s and %s", check_date.String(), day_after.String(), now.String())
	}
}
