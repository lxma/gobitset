# gobitset

Simple uint64 bitset for Go.

Sets are uint64 values. Elements are integers from 0 to 63 (inclusive). Sets
can be used as keys in maps. Basic operations like union and intersection are executed by
bit functions.

Using this library will not make your code much shorter. It will only enhance readability.
E.g., `Union(set1, set2)` isn't shorter than `set1 | set2`, but it's easier to read. And

```go
it := set.GetIterator()
for it.HasNext() {
    fmt.Printf("Value: %d\n", it.NextInt())
}
```
is probably easier to grasp than
```go
it := set
for it > 0 {
    intValue := TrailingZeros64(it)
    it &^= uint64(1) << intValue
    fmt.Printf("Value: %d\n", intValue)
}
```


