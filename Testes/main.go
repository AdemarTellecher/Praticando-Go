package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func FibonacciComFor(n int) string {
	prev_n1 := 1
	prev_n2 := 0
	total := (prev_n1 + prev_n2)

	switch n {
	case 0:
		return fmt.Sprintf("%d", n)
	case 1:
		return fmt.Sprintf("%d", n)
	default:
		for i := 2; i <= n; i++ {
			fibo_i := prev_n1 + prev_n2
			total += fibo_i
			prev_n2 = prev_n1
			prev_n1 = fibo_i
		}
	}
	return fmt.Sprintf("%d", total)
}

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

	//resultFibo := fibonacciComFor(5)
	fmt.Println(FibonacciComFor(3))
}
