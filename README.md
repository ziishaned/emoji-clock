# emoji-clock [![Build Status](https://img.shields.io/travis/ziishaned/emoji-clock.svg?style=flat-square)](https://travis-ci.org/ziishaned/emoji-clock) [![Go Report](https://img.shields.io/badge/go%20report-A%2B-brightgreen.svg?style=flat-square)](https://goreportcard.com/report/github.com/ziishaned/emoji-clock) [![Codecov](https://img.shields.io/codecov/c/github/ziishaned/emoji-clock.svg?style=flat-square)](https://codecov.io/gh/ziishaned/emoji-clock) [![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](https://github.com/ziishaned/emoji-clock)

Get the emoji clock face for a given time.

## Install

To install this library run the following command inside your terminal:

```bash
go get github.com/ziishaned/emoji-clock
```

## Usage

```go
package main

import (
  "fmt"
  "log"
  "github.com/ziishaned/emoji-clock"
)

func main() {
  emoji, err := emojiclock.TimeToEmoji("2014-03-09T22:47:02.705Z")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(emoji) // 🕚
  
  emoji2, err := emojiclock.TimeToEmoji("2018-10-09T21:40:02.705Z")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(emoji2) // 🕤
}
```

## Contributions

Feel free to submit pull requests, create issues or spread the word.

## License

MIT &copy; [Zeeshan Ahmad](https://twitter.com/ziishaned)
