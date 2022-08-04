package interfaces

import "fmt"

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func main() {
	rs := Person{"Maysa Pereira", "Brazil"}
	ab := Person{"Byun Baekhyun", "South Korea"}

	fmt.Printf("%s\n%s\n", rs, ab)
}
