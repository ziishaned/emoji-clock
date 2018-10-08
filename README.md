# emoji-clock [![Build Status](https://travis-ci.org/ziishaned/emoji-clock.svg?branch=master)](https://travis-ci.org/ziishaned/emoji-clock)

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
)

func main() {
  currentTime := time.Now()
  emoji := TimeToEmoji(currentTime, false)
  fmt.Println(emoji)
}
```

### Output

```txt
🕒
```

## Contributions

Feel free to submit pull requests, create issues or spread the word.

## License

MIT &copy; [Zeeshan Ahmad](https://twitter.com/ziishaned)
