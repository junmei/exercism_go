package clock

import "fmt"

const testVersion = 4

// Clock ....
type Clock struct {
	hour, minute int
}

func (c Clock) normalize() Clock {
	c.hour = (c.hour + c.minute/60) % 24
	c.minute = c.minute % 60
	if c.minute < 0 {
		c.minute += 60
		c.hour--
	}
	if c.hour < 0 {
		c.hour += 24
	}
	return c
}

// New : create a new Clock
func New(hour, minute int) Clock {
	var c = Clock{hour, minute}
	return c.normalize()
}

func (c Clock) String() string {
	var s = fmt.Sprintf("%02d:%02d", c.hour%24, c.minute)
	return s
}

// Add Add minutes to Clock c
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c.normalize()
}
