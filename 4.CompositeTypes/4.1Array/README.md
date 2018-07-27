An array is a ﬁxed-length sequence of zero or more elements of aparticular type.

`var a [3]int` // array of 3 integers 

`fmt.Println(a[0])` // print the first element 

`fmt.Println(a[len(a)-1])` // print the last element, a[2]

// Print the indices and elements. 
```
for i, v := range a { 
    fmt.Printf("%d %d\n", i, v) 
}
```

// Print the elements only. 
```
for _, v := range a {
    fmt.Printf("%d\n", v) 
}
```

New array variable are initiallyset to the zero value 
```
var q [3]int = [3]int{1, 2, 3} 
var r [3]int = [3]int{1, 2} 
fmt.Println(r[2]) // "0"
```

If ane llipsis ‘‘...’’ appears in place of the length, 
the array length is determined by the number of initializers
```
q := [...]int{1, 2, 3} 
fmt.Printf("%T\n", q) // "[3]int"
```

The size of an array(which is constant) is part of its type, 
so [3]int and [4]int are different types.
```
q := [3]int{1, 2, 3} 
q=[4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
```

Specify alist of index and value pairs.
```
type Currency int

const ( 
    USD Currency = iota 
    EUR 
    GBP 
    RMB 
)

symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: """}

fmt.Println(RMB, symbol[RMB]) // "3 ""
```
In this form, indices can appear in any order and some may be omitted
`r := [...]int{99: -1}` 
This deﬁnes an arrayr with 100 elements, all zero except for the last, which has value −1.

Array typ e is comparable.
```
a := [2]int{1, 2} 
b := [...]int{1, 2} 
c := [2]int{1, 3} 
fmt.Println(a == b, a == c, b == c) // "true false false" 
d := [3]int{1, 2} 
fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
 ``` 

When a function is called, a copy of each argument value is assigned to the corresponding 
parameter variable, so the function receives a copy, not the original. 
Passing large arrays in this way can be inefﬁcient, and any changes that the function 
makes to array elements affect only the copy, not the original.

Of course, we can explicitly pass a pointer to an array so that any modiﬁcations the 
function makes to array elements will be visible to the caller
```
func zero(ptr *[32]byte) { 
    for i := range ptr { 
        ptr[i] = 0 
    } 
}
```

The array literal `[32]byte{}` yields an array of 32 bytes. 
Each element of the array has the zero value for byte, which is zero.
```
func zero(ptr *[32]byte) { 
    *ptr = [32]byte{} //the outer array will take the new value
}
```
The zero function will not accept a pointer to a[16]byte variable,
for example, nor is there any way to add or remove array elements.
