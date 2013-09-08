package poller

import (
	"testing"
)

func TestNewEpoller(t *testing.T) {
	p, err := newEpoller()
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Close(); err != nil {
		t.Fatal(err)
	}
}
