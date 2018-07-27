#Slices
Slices represent variable-length sequences whose elements all have the same type. A slice type
is written []T, where the elements have type T; it looks like an array typ e without a size.
A slice is a lightweight data structure that gives access to a subsequence (or perhaps all) 
of the elements of an array, which is known as the slice’s `underlying array`.
A slice has three components: a `pointer`, a `length`, and a `capacity`.

The `pointer` points to the first element of the array that is reachable through the slice, which is not
necessarily the array’s first element.
The `length` is the number of slice elements; it can’t exceed
the `capacity`, which is usually the number of elements between the start of the slice and the end
of the underlying array.

Multiple slices can share the same `underlying array` and may refer to overlapping parts of that array.
The slice operator s[i:j], where 0 <= i <= j <= cap(s), creates a new slice that refers to elements
i through j-1 of the sequence s, which may be an array variable, a pointer to an array, or another slice.
- s[i:j] refers from i to j - 1
- s[:j] from 0 to j - 1
- s[i:] from i to len(s) - 1
- s[:] refers to the whole array

```
months := [...]string{1: "January", /* ... */, 12: "December"}
//here months[0] is ""
Q2 := months[4:7]
summer := months[6:9]
fmt.Println(Q2) // ["April" "May" "June"]
fmt.Println(summer) // ["June" "July" "August"]
```

June is included in each and is the sole output of this (inefficient) test for common elements:
```
for _, s := range summer {
    for _, q := range Q2 {
        if s == q {
        fmt.Printf("%s appears in both\n", s)
        }
    }
}
```
Slicing beyond cap(s) causes a panic, but slicing beyond len(s) extends the slice, so the
result may be longer than the original:
```
fmt.Println(summer[:20]) // panic: out of range
endlessSummer := summer[:5] // extend a slice (within capacity)
fmt.Println(endlessSummer) // "[June July August September October]"
```
####Passing slice to function 
Since a slice contains a pointer to an element of an array, passing a slice to a function permits
the function to modify the underlying array elements.
See revers function.
```
func reverse(s []int){
    ....
}

//this declare and init a slice. !!!The size in [] is missing!!!
s := []int{0, 1, 2, 3, 4, 5}
reverse(s) // "[5 4 3 2 1 0]"
```
####Slice literal
A slice literal looks like an array literal, a sequence of values separated by commas and 
surrounded by braces, but the size is not given

`s := []int{0, 1, 2, 3, 4, 5}` //slice literal

As with array literals, slice literals may specify the values in order, or give their indices 
explicitly, or use a mix of the two styles

####Comparing slices
Unlike arrays, `slices are not comparable`, so we cannot use == to test whether two slices contain 
the same elements. The standard library provides the highly optimized `bytes.Equal` function for 
comparing two slices of bytes ([]byte), but for other types of slice, we must do the 
comp arisonourselves:
```
//deep equality
func equal(x, y []string) bool { 
    if len(x) != len(y) { 
        return false 
    } 
    for i := range x { 
        if x[i] != y[i] { 
            return false 
        } 
    } 
    return true 
}
```
Problem with deep equality
- slice can contain itself.
- slice can be mutated during comparison

This makes slices   unsuitablefor use as map keys.
#####Operator ==
The `==` operator tests reference identity,that is, whether the two entities refer to the samething.
Theonlylegal slice comp arison is against nil, as in

`if summer == nil { /* ... */ }`

####Default value
The zero value of a slice type is `nil`. A nil slice has no underlying array. 
The nil slice has length and capacity zero,
```
//no ynderlying array
var s []int // len(s) == 0, cap(s) == 0, s == nil 
s=nil // len(s) == 0, s == nil 
s=[]int(nil) // len(s) == 0, s == nil 
//with underlaying array
s=[]int{} // len(s) == 0, s != nil
 make([]int, 3)[3:] // len(s) == 0, cap(s) == 0, s != nil
```

So, if you need to test whether a slice is empty, use len(s) == 0, not s == nil

>Other than comparing equal to nil, a nil slice be haves like any other zero-length slice; 
`reverse(nil)` is perfectly safe, for example. Unless clearly documented to the contrary, 
Go functions should treat all zero-length slices the same way, whether nil or non-nil.


####build in function `make`
```
make([]T, len) 
make([]T, len, cap) // same as make([]T, cap)[:len]
```
Under the hood, make creates an unnamed array variable and returns a slice of it; 
the array is accessible only through the returned slice. In the ﬁrst form, the slice is a 
view of the entire array. In the second, the slice is a view of only the array’s ﬁrst len elements, 
but its capacity includes the entire array. The addition alelements are set a side for future growth.
