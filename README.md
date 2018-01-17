# go-dequeue
Simple efficient double ended queue

```package main

import (
  "fmt"
  "github.com/lleo/go-dequeue"
)

func main() {
	var d = dequeue.New().
		Push("a").
		Push("b").
		Push("c")

	var i int
	d.Range(func(data interface{}) bool {
		fmt.Printf("%d: %q\n", i, data)
		i++
	}
}
```
This Dequeue data structure supports Push(), Pop(), Shift(), and Unshift().
The key feature is that all operations, except Range(), are O(1) performance.
