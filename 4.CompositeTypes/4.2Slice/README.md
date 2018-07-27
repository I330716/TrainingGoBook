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

Since a slice contains a pointer to an element of an array, passing a slice to a function permits
the function to modify the underlying array elements.
See revers function.

`s := []int{0, 1, 2, 3, 4, 5}` //slice literal

