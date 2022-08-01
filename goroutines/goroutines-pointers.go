package main

import (
	"fmt"
	"time"
)

func funPointer(value *string) {
	for {
		fmt.Println(*value)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	var test string = "test"
	var pointTest *string = &test

	go funPointer(pointTest)

	time.Sleep(time.Second)

	*pointTest = "test 2"

	time.Sleep(3 * time.Second)
}
