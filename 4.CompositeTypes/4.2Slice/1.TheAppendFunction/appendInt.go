//Function which simulate the original append one
package main

import "fmt"

func appendInt(oldSlice []int, values ...int) []int {
	var newSlice []int // makes nill slice
	newLength := len(oldSlice) + len(values)
	if newLength <= cap(oldSlice) {
		//There is room to grow in the underlying array. Extend the slice
		newSlice = oldSlice[:newLength]
		/*the newSlice point to the underlying array of the oldSlice
		but has one length more that the old one*/
	} else {
		/*There is insufficient space in the underlying array.
		So we have to allocate new array. Grow by doubling, for
		amortized linear complexity*/
		newCapacity := newLength
		if newCapacity < 2*len(oldSlice) {
			newCapacity = 2 * len(oldSlice)
		}
		newSlice = make([]int, newLength, newCapacity)
		copy(newSlice, oldSlice) // a build-iin function
	}
	copy(newSlice[len(oldSlice):], values)
	return newSlice
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var z, c []int
	for i := 0; i < 10; i++ {
		c = append(z, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(c), c)
		z = c
	}
}
