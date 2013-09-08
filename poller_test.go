package poller

import (
	"testing"
)

func TestNewPoller(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Fatal(err)
	}
}
