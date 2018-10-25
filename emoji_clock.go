package emojiclock

import (
  "time"
)

var clocks = []rune("🕛🕧🕐🕜🕑🕝🕒🕞🕓🕟🕔🕠🕕🕖🕡🕗🕣🕢🕘🕤🕙🕥🕚🕦")

// TimeToEmoji pass the time string of which you want the clock emoji.
func TimeToEmoji(timeString string) (string, error) {
  timestamp, err := time.Parse(time.RFC3339, timeString)
  if err != nil {
    return "", err
  }

  // New emoji every 30 minutes
  t := (60*timestamp.Hour() + timestamp.Minute() + 15) / 30
  emoji := clocks[t%24]

  return string(emoji), nil
}
