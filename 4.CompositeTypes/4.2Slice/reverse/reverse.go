//Reverse an slice inplace
package reverse

import "fmt"

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*A simple way to rotate a slice left by n elements is to apply the reverse function three times,
first to the leading n elements, then to the remaining elements, and finally to the whole slice*/
func rotateLeft(s []int, n int) {
	// Rotate s left by two positions.
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func main() {
	//this deaclere and init an array
	a := [...]int{0, 1, 2, 3, 4, 5}
	//this will affect the outer array
	reverse(a[:])  //passing array but receiveing slice in the function
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//this declare and init a slice. !!!The size in [] is missing!!!
	s := []int{0, 1, 2, 3, 4, 5}
	//this will affect the outer array
	rotateLeft(s, 2)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}
