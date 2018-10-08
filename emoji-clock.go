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

// TimeToEmoji convert a provided time to emoji
func TimeToEmoji(time time.Time, utc bool) {
	var hours int
	var minutes int
	if utc == true {

	}
}

func main() {
	currentTime := time.Now()
	TimeToEmoji(currentTime, true)
}
