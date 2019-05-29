package main

func EliminateAdjasentDuplicates(strings []string) []string {
	//if the sice is empty or there is only one element
	//there can no be any adjacent strings
	if len(strings) < 2 {
		return strings
	}

	//on this position will end the result array with no duplicated adjacent strings
	var resultBound int

	for index := 1; index < len(strings); index++ {
		//if the next string deffers from the last in the result move the bound one forward
		if strings[index] != strings[resultBound] {
			resultBound++
			//if the index and bound are one the same place no assignmet is needed
			if index != resultBound {
				strings[resultBound] = strings[index]
			}
		}
	}
	//retun the result slice from 0 to bound included
	return strings[:resultBound+1]
}
