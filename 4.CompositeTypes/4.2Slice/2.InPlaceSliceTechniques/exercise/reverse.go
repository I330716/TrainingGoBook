package main

// Reverse takes as an argument pointer to an array and reverse this array in place
func Reverse(arrayPtr *[8]int) {
	for front, back := 0, 7; front < back; front, back = front+1, back-1 {
		(*arrayPtr)[front], (*arrayPtr)[back] = (*arrayPtr)[back], (*arrayPtr)[front]
	}
}
