package main

import (
	"time"
)

// clock is the list of emoji's one will be returned on the basis
// of given time.
var clock = map[float64]string{
	0:    "🕛",
	1:    "🕐",
	2:    "🕑",
	3:    "🕒",
	4:    "🕓",
	5:    "🕔",
	6:    "🕕",
	7:    "🕖",
	8:    "🕗",
	9:    "🕘",
	10:   "🕙",
	11:   "🕚",
	0.5:  "🕧",
	1.5:  "🕜",
	2.5:  "🕝",
	3.5:  "🕞",
	4.5:  "🕟",
	5.5:  "🕠",
	6.5:  "🕡",
	7.5:  "🕢",
	8.5:  "🕣",
	9.5:  "🕤",
	10.5: "🕥",
	11.5: "🕦",
}

// hoursAndMinutesToEmoji will return the emoji on the basis of
// hours and minutes.
func hoursAndMinutesToEmoji(hours float64, minutes float64) string {
	if hours > 11 {
		hours = hours - 12
	}
	minutes = minutes / 60
	if minutes < 0.25 {
		minutes = 0
	} else if minutes >= 0.25 && minutes < 0.75 {
		minutes = 0.5
	} else {
		if hours == 11 {
			hours = 0
		} else {
			hours = hours + 1
		}
		minutes = 0
	}
	return clock[float64(hours)+minutes]
}

// TimeToEmoji pass the time string of which you want the clock emoji.
func TimeToEmoji(timeString string) (emoji string, err error) {
	timestamp, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return
	}

	hours := float64(timestamp.Hour())
	minutes := float64(timestamp.Minute())
	emoji = hoursAndMinutesToEmoji(hours, minutes)
	return
}
