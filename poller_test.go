package poller

import (
	"os"
	"testing"
)

func TestNewPoller(t *testing.T) {
	p, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestPollerRegister(t *testing.T) {
	p, err := New()
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	fd, err := p.Register(r)
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()
}

func TestFDWaitWrite(t *testing.T) {
	p, err := New()
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	fd, err := p.Register(w)
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()
	if err := fd.WaitWrite(); err != nil {
		t.Fatal(err)
	}
}
