package clock

import "fmt"

// Define the Clock type here.

type Clock struct {
	Hours   int
	Minutes int
}

const (
	minPerHour = 60
	hourPerDay = 24
)

func norm(h int, m int) (int, int) {
	for m >= minPerHour {
		m -= minPerHour
		h++
	}
	for h >= hourPerDay {
		h -= hourPerDay
	}
	for m < 0 {
		m += minPerHour
		h--
	}
	for h < 0 {
		h += hourPerDay
	}
	return h, m
}

func New(h, m int) Clock {
	h, m = norm(h, m)
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hours, c.Minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hours, c.Minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02v:%02v", c.Hours, c.Minutes)
}
