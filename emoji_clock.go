package main

import (
	"time"
)

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

func TimeToEmoji(time time.Time, utc bool) string {
	var hours float64
	var minutes float64
	if utc == true {
		hours = float64(time.UTC().Hour())
		minutes = float64(time.UTC().Minute())
	} else {
		hours = float64(time.Hour())
		minutes = float64(time.Minute())
	}
	return hoursAndMinutesToEmoji(hours, minutes)
}
