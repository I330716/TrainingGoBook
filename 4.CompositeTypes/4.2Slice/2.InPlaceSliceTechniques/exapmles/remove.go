package main

// Remove removes the element of a int slice denoted by index.
// The function preserves the order of the rest elements.
//The function work in place so the otiginal slice will be modified
func Remove(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// Remove2 removes the element of a int slice denoted by index.
// The function doesn not preserves the order of the rest elements.
//The function work in place so the otiginal slice will be modified
func Remove2(slice []int, index int) []int {
	//put the last index at the place we whant to remove
	slice[index] = slice[len(slice)-1]
	//now return the same slice without the last element
	return slice[:len(slice)-1]
}
