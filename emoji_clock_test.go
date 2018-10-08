package main

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeToEmojiShouldReturnString(t *testing.T) {
	var utc bool
	currentTime := time.Now()
	emoji := TimeToEmoji(currentTime, utc)
	if reflect.TypeOf(emoji).String() != "string" {
		t.Errorf("Expected type was string but got %s", reflect.TypeOf(emoji).String())
	}
}
