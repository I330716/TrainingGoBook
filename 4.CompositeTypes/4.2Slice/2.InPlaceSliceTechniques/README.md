# In-Place Slice Techniques

When we have to deal with function where the underlying array of a slice is changed we have to know that passed as an argument slice will be modified and thus this will efect the slice after the function returns. So we have to make that for each input to produce at leas one ouput which will be asssigned to the slice which is passed as an argument. For example.
```
data := []string{"one", "", "three"}
fmt.Printf("%q\n", Nonempty1(data)) // `["one" "three"]`
//the original slice is modified
fmt.Printf("%q\n", data) // `["one" "three" "three"]`
```

As you can see after this example the data slice is changed and if some one is using it thinking that it has {"one", "", "three"} will have unlanted result. So we have to do the following:
```
data := []string{"one", "", "three"}
data = Nonempty1(data))
fmt.Printf("%q\n", data) // `["one" "three"]`
```

As you can see we have to know two things:
- inplace function should retund a output
- the output should be assigned to the input slice

>The rule above may be not true when comes to a function that makes inplace modifications which are wanted as a result.
For example sorting, reversing and e.tc.

See the examples for more details.