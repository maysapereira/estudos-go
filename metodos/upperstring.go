package main

import (
	"fmt"
	"strings"
)

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func main() {
	s := upperstring("Learning GO!")
	fmt.Println(s)
	fmt.Println(s.Upper())
}
