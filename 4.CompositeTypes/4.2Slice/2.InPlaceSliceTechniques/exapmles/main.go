//This packeg represent the teaching about in place slice technigues
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", Nonempty1(data)) // `["one" "three"]`
	//the original slice is modified
	fmt.Printf("%q\n", data) // `["one" "three" "three"]`
	data = []string{"one", "", "three"}
	//It is best to use the same slice as assignmet operand on the left side
	//because this way we are not going to use the old modified slice
	data = Nonempty2(data)
	fmt.Printf("%q\n", data) // `["one" "three"]`

	stack := IntStack{} // make stack object with nil underlying slice
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)            // add element in the stack
	fmt.Println(stack.Top()) // '6'
	fmt.Println(stack)
	stack.Pop() // remove the top element
	fmt.Println(stack)

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(Remove(s, 2)) // "[5 6 8 9]"
	s = []int{5, 6, 7, 8, 9}
	fmt.Println(Remove2(s, 2)) // "[5 6 9 8]
}
