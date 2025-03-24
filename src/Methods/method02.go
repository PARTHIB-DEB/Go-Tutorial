package main

import (
	"fmt"
	"time"
)

/*
Go uses parameters of POINTER TYPE to indicate that a parameter might be modified by the function.
The same rules apply for method receivers, too.

They can be POINTER RECEIVER (the type is a pointer) / VALUE RECEIVERS (the type is a value type).
The following rules help you determine when to use each kind of receiver:
1. If your method modifies the receiver, you must use a pointer receiver.
2. If your method needs to handle nil instances , then it must use a pointer receiver.
3. If your method doesn’t modify the receiver, you can use a value receiver.

Pointer receivers work just like pointer function parameters; it’s a copy of the pointer that’s
passed into the method. Just like nil parameters passed to functions, if you change the copy of
the pointer, you haven’t changed the original. This means you can’t write a pointer receiver method
that handles nil and makes the original pointer non-nil
*/

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
}
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {

	c := Counter{
		total:       12,
		lastUpdated: time.Now(),
	}

	c.Increment()

	output := c.String()
	fmt.Println(output)

}
