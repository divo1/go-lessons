// _Interfaces_ are named collections of method
// signatures.
package lessons

import(
	"fmt"
	"math"
)

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type irect struct {
	width, height float64
}

type icircle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (r irect) area() float64 {
	return r.width * r.height
}
func (r irect) perim() float64 {
	return 2 * r.width + 2 * r.height
}

// The implementation for `circle`s
func (c icircle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c icircle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func Interfaces() {
	r := irect{width: 3, height: 4}
	c := icircle{radius: 5}
	
	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of
	// these structs as arguments to `measure`.
	measure(r)
	measure(c)
}