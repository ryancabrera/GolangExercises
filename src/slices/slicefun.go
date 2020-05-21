//To Practice slices, based off of golang tour and the slice tricks https://ueokande.github.io/go-slice-tricks/
package slices

import "fmt"

func Main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
	printLen()
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printLen() {
	fmt.Printf("Expansion vs Extension exercises\n")

	b := []int{0, 1, 2, 3}
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	b = b[:cap(b)]
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	b = append(b, make([]int, len(b))...)
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			b[i] = i
		}
	}
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	b = append(b[:0], append(make([]int, len(b)), b[0:]...)...)
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)
}
