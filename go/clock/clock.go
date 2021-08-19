package clock

import (
	"fmt"
)

type Clock struct {
	hours   int
	minutes int
}

func New(hours int, minutes int) Clock {
	checkInput(&hours, &minutes)
	return Clock{hours, minutes}
}

func (clk Clock) Subtract(min int) Clock {
	return New(clk.hours, clk.minutes-min)
}

func (clk Clock) String() string {
	return fmt.Sprintf("%02d", clk.hours) + ":" + fmt.Sprintf("%02d", clk.minutes)
}

func (clk Clock) Add(min int) Clock {
	return New(clk.hours, clk.minutes+min)
}

func checkInput(hours *int, minutes *int) {
	if *hours < 0 {
		for *hours < 0 {
			*hours += 24
		}
	} else {
		*hours = *hours % 24
	}
	*hours = *hours % 24
	if *minutes >= 60 {
		for *minutes >= 60 {
			*minutes -= 60
			*hours++
		}
		if *hours > 23 {
			*hours = *hours % 24
		}
	}
	if *minutes < 0 {
		for *minutes < 0 {
			*minutes += 60
			*hours--
		}
		if *hours < 0 {
			for *hours < 0 {
				*hours += 24
			}
		}
	}
}
