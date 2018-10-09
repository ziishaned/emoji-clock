# emoji-clock [![Build Status](https://img.shields.io/travis/ziishaned/emoji-clock.svg?style=flat-square)](https://travis-ci.org/ziishaned/emoji-clock) [![Codecov](https://img.shields.io/codecov/c/github/ziishaned/emoji-clock.svg?style=flat-square)](https://codecov.io/gh/ziishaned/emoji-clock) [![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](https://github.com/ziishaned/emoji-clock)

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
  "time"
  . "github.com/ziishaned/emoji-clock"
)

func main() {
  timestamp := "2016-06-17T20:41:02.705Z"
  emoji, _ := TimeToEmoji(timestamp)
  fmt.Println(emoji) // 🕒
}
```

## Contributions

Feel free to submit pull requests, create issues or spread the word.

## License

MIT &copy; [Zeeshan Ahmad](https://twitter.com/ziishaned)
