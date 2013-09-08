package poller

import (
	"syscall"
)

type Poller struct {
	*epoller
}

// New returns a Poller which can be used register
// character devices.
func New() (*Poller, error) {
	e, err := newEpoller()
	if err != nil { return nil, err }
	return &Poller{
		epoller: e,
	}, nil
}

type Pollable interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
	Fd() uintptr
}

type FD struct {
	*epoller
}

func (P *Poller) Register(p Pollable) (*FD, error) {
	err := syscall.SetNonblock(int(p.Fd()), true)
	return &FD{
		epoller: P.epoller,
	}, err
}

func (f *FD) WaitWrite() error {
	return nil
}

type epoller struct {
	fd int
}

func newEpoller() (*epoller, error) {
	fd, err := syscall.EpollCreate1(syscall.EPOLL_CLOEXEC)
	if err != nil { return nil, err }
	e := &epoller{
		fd: fd,
	}
	return e, nil
}

func (e *epoller) Close() error {
	return syscall.Close(e.fd)
}
