# program-rlimit-tuning
Under Linux system, the number of files that a process can open is 1024 by default.   
If the process needs high performance, we need to change the file open limit.

### Install

Run the following command under your project:

> go get -u github.com/vvfly/base/program-rlimit-tuning

### Usage

Just import it directly, such as:

```go
package main

import (
	_ "github.com/vvfly/base/program-rlimit-tuning"
	"log"
)

func main() {
	log.Println("server is running")
}
```