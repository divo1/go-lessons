package lessons

import(
	"fmt"
)

type person struct {
	name string
	age int
}

func Structs() {
	// This syntax creates a new struct
	fmt.Println(person{"Bob", 20})
	
	// You can name the fields when initializin a struct
	fmt.Println(person{name: "Alice", age: 30})
	
	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})
	
	// An `&` prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})
	
	// Access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	
	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)
}