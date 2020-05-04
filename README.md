# godl

Download file with Progress

## Installation

```bash
$ go get github.com/dqn/godl
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/dqn/godl"
)

func main() {
	err := godl.Download("https://api.github.com/users/dqn", "./user.json", true)
	if err != nil {
		// handle error
	}
	fmt.Println("completed!")
}
```

Output:

```bash
downloading... 1.23/1.23 MB
completed!
```

## API

### func Download(rawURL, dest string, showProgress bool) error

Download file.
