//Function which simulate the original append one
package main

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
