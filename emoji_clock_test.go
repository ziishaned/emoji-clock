package main

import (
  "log"
  "reflect"
  "testing"
  "time"
)

func TestItShouldReturnString(t *testing.T) {
  currentTime := time.Now()
  emoji := TimeToEmoji(currentTime, false)
  if reflect.TypeOf(emoji).String() != "string" {
    t.Errorf("Expected type was string but got %s", reflect.TypeOf(emoji).String())
  }
}

func TestItShouldReturnTheCorrectEmoji(t *testing.T) {
  timeStr := "2014-11-12T11:45:26.371Z"
  time2, err := time.Parse("2006-01-02T15:04:05.000Z", timeStr)
  if err != nil {
    log.Fatalln(err)
  }
  emoji := TimeToEmoji(time2, true)
  if emoji != "🕛" {
    t.Errorf("Expected emoji is 🕛 but received %s", emoji)
  }
}
