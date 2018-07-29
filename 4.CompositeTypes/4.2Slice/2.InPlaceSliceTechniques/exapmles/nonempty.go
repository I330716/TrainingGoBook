package main

// Nonempty1 returns a slice holding only the nonempty strings.
// The underlying array is modified during the call.
func Nonempty1(strings []string) []string {
	lastNotEmptyIndex := 0 //starts from the begging of the array
	for _, str := range strings {
		if str != "" {
			//move the non empty string in place somewhere back to the
			//underlying array
			strings[lastNotEmptyIndex] = str
			// move to the nex position of the underlying array where will will put the
			//next non empty string if any
			lastNotEmptyIndex++
		}
	}
	//lastNotEmptyIndex is one step after the end of the non emptier strings
	return strings[:lastNotEmptyIndex]
}

// Nonempty2 returns a slice holding only the nonempty strings.
// The underlying array is modified during the call. It uses build-in copy
func Nonempty2(strings []string) []string {
	//Make zero length slice which point to the same array
	nonEmptyStrings := strings[:0]
	for _, str := range strings {
		if str != "" {
			//this will move the str in the same underluing array
			nonEmptyStrings = append(nonEmptyStrings, str)
		}
	}
	return nonEmptyStrings
}
