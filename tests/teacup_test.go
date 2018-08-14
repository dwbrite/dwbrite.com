package tests

import "testing"
import "../teacup"

func AssertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic (codepoint? ?? ?)")
		}
	}()
	f()
}

func TestTeacupInstantiation(t *testing.T) {
	teacup.New(443, "../certs/dwbrite.com.cert", "../certs/dwbrite.com.key")
	privileged := false
	if r := recover(); r == nil {
		if r == "" {
			privileged = true
		}
	}
	println(privileged)
}