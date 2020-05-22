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

	b := []int{4, 5, 6, 7}
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	b = b[:cap(b)]
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	indexBeforeDoubling := len(b)

	b = append(b, make([]int, len(b))...)
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	fmt.Printf("Clearing Trailing Values\n")
	//b = b[:indexBeforeDoubling]
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)
	b[indexBeforeDoubling] = 22

	b = append(b[:0], append(make([]int, len(b)), b[0:]...)...)
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	arr1 := [9]int{1, 2, 3}
	fmt.Printf("len=%d cap=%d %v\n", len(arr1), cap(arr1), arr1)
	arr1[len(arr1)-1] = 200
	fmt.Printf("len=%d cap=%d %v\n", len(arr1), cap(arr1), arr1)

}
