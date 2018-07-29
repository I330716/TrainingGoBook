# The append Function

The built-in append function appends items to slices:
```
var runes []rune 
for _, r := range "Hello, BF"{ 
    runes = append(runes, r) 
} 
fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l'  o' ',' ' ' 'B' 'F']"
```

If there is insufficient space for growth, append must allocate a new array big enough to hold the result. copy the valuens from the original array and then append the new element. For efficiency, the new array is usually somewhat larger than the minimum needed to hold the old values and the new ones together. This avoids an excessive number of allocations and ensures that appending a single element takes constant time on average.
```
func main() {
    var x, y []int
    for i := 0; i < 10; i++ {
        y = append(x, i)
        fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
        x = y
    }
}

Output:
0 cap=1  [0] //returns new undelaying array
1 cap=2  [0 1] //returns new undelaying array
2 cap=4  [0 1 2] //returns new undelaying array
3 cap=4  [0 1 2 3] //returns a slice that points to the old undelying array
4 cap=8  [0 1 2 3 4] //returns new undelaying array
5 cap=8  [0 1 2 3 4 5] //returns a slice that points to the old undelying array
6 cap=8  [0 1 2 3 4 5 6] //returns a slice that points to the old undelying array
7 cap=8  [0 1 2 3 4 5 6 7] //returns a slice that points to the old undelying array
8 cap=16 [0 1 2 3 4 5 6 7 8] //returns new undelaying array
9 cap=16 [0 1 2 3 4 5 6 7 8 9] //returns a slice that points to the old undelying array
```

Usually we don’t know whether a given call to append will cause a reallocation, so we can’t assume that the original slice refers to the same array as the resulting slice, nor that it refers to a different one. Similarly, we must not assume that operations on elements of the old slice will (or will not) be reflec ted in the new slice. As a result, it’s usual to assign the result of a call to append to the same slice variable whose value we passed to append: `runes = append(runes, r)`

Updating the slice variable(as above) is required for any function that may change the length or capacity of a slice or make it refer to a different underlying array. Although the elements of the underlying array are indirect, the slice’s pointer, length, and capacity are not. To update them requires an assignment like the one above. In this respect, slices are not ‘‘pure’’ reference types but resemble an aggregate type such as this struct:

```
type IntSlice struct {
    ptr *int
    len, cap int
}
```

> So when we whant to append a rune to slice 'runes' we must make assignment because 'runes' capacity, length and pointer are passed as value and will not ne changed after the function returns.

append lets us add more than one new element, or even a whole slice of them.
```
var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, 4, 5, 6)
x = append(x, x...) // append the slice x
fmt.Println(x) // "[1 2 3 4 5 6 1 2 3 4 5 6]"
```