package main

//RotateToTheLeft rotates the elements to the left given number of times
func RotateToTheLeft(slice []int, times int) {
	sliceSize := len(slice)
	times %= sliceSize
	if times <= 0 || sliceSize < 2 {
		return
	}

	for row := 0; row < gcd(sliceSize, times); row++ {
		itemToPutAtTheFinalStep := slice[row]
		previousMovedIndex := row
		for true {
			nextIndexToMoveNtimesLeft := previousMovedIndex + times
			if nextIndexToMoveNtimesLeft >= sliceSize {
				nextIndexToMoveNtimesLeft -= sliceSize
			}
			if nextIndexToMoveNtimesLeft == row {
				break
			}
			slice[previousMovedIndex] = slice[nextIndexToMoveNtimesLeft]
			previousMovedIndex = nextIndexToMoveNtimesLeft
		}
		slice[previousMovedIndex] = itemToPutAtTheFinalStep
	}
}

func gcd(devidend, devisor int) int {
	for devisor != 0 {
		devidend, devisor = devisor, devidend%devisor
	}
	return devidend
}
