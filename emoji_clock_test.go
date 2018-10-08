package main

import (
  "reflect"
  "testing"
)

func TestItShouldReturnString(t *testing.T) {
  currentTime := "2014-11-12T11:45:26.371Z"
  emoji := TimeToEmoji(currentTime, false)
  if reflect.TypeOf(emoji).String() != "string" {
    t.Errorf("Expected type was string but got %s", reflect.TypeOf(emoji).String())
  }
}

func TestItShouldReturnTheCorrectEmoji(t *testing.T) {
  var data = map[int]map[string]string{
    0: {
      "timestamp": "2014-11-12T11:45:26.371Z",
      "expected":  "🕛",
    },
    1: {
      "timestamp": "2018-06-05T01:15:26.371Z",
      "expected":  "🕜",
    },
    2: {
     "timestamp": "2018-06-05T20:45:26.371Z",
     "expected":  "🕘",
    },
    3: {
     "timestamp": "2018-06-05T21:45:26.371Z",
     "expected":  "🕙",
    },
    4: {
     "timestamp": "2018-06-05T22:45:26.371Z",
     "expected":  "🕚",
    },
    5: {
     "timestamp": "2018-06-05T23:45:26.371Z",
     "expected":  "🕛",
    },
  }

  for key, val := range data {
    emoji := TimeToEmoji(val["timestamp"], true)
    if emoji != val["expected"] {
      t.Errorf("%b: Expected emoji is %s but received %s", key, val["expected"], emoji)
    }
  }
}
