package generics

import "fmt"

type User[T any, B any] struct {
	name T
	age  B
}

func main() {
	userTest := User[string, int64]{
		name: "Maysa",
		age:  22,
	}

	fmt.Println(userTest)
}
