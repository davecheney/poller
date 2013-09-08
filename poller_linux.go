package poller

import (
	"os"
)

type Poller struct {
	r, w *os.File
}

// New returns a Poller which can be used register
// character devices.
func New() (*Poller, error) {
	r, w, err := os.Pipe()
	return &Poller{r: r, w: w}, err
}
