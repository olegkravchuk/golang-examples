package clock

import (
	"fmt"
	"time"
)

const testVersion = 4

type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), h, m, now.Second(), 0, now.Location())
	return Clock{t.Hour(), t.Minute()}
}

func (t Clock) String() string {
	return fmt.Sprintf("%02d:%02d", t.h, t.m)
}

func (t Clock) Add(m int) Clock {
	return New(t.h, t.m+m)
}
