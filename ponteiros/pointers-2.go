package main

import "fmt"

func main() {
	var testValue string = "Maysa"
	copyStringVALUE(testValue)
	fmt.Println(testValue)

	originalStringValue(&testValue)
	fmt.Println(testValue)
}

func copyStringVALUE(stringValue string) {
	stringValue = "TEST 1"
	fmt.Println(stringValue)
}

func originalStringValue(stringValue *string) {
	*stringValue = "TEST 2"
	fmt.Println(*stringValue)
}
