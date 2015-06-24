package sendmsg

import "testing"

func TestToSlack(t *testing.T) {
	code := ToSlack("test_message")
	if code >= 400 {
		t.Error("Message not delivered")
	}
}
