package bitset

import (
    "fmt"
    "math/bits"
)

type Bitset uint64
type Iterator uint64

const EmptySet = Bitset(0)

// MakeSet makes a bitset from one or more integers.
func MakeSet(intValues ...int) Bitset {
    set := Bitset(0)
    for _, intValue := range intValues {
        if intValue > 63 || intValue < 0 {
            panic(fmt.Errorf("invalid value for bitset: %d (value must be >= 0 and <= 63)", intValue))
        }
        set |= Bitset(1) << intValue
    }
    return set
}

// AddInts returns a bitset containing the elements of bs and all
// added integer values
func AddInts(bs Bitset, intValues ...int) Bitset {
    for _, intValue := range intValues {
        bs |= 1 << intValue
    }
    return bs
}

// DeleteInts retunrs a bitset containing the elements of bs except
// for the given integers values
func DeleteInts(bs Bitset, intValues ...int) Bitset {
    for _, intValue := range intValues {
        bs &^= 1 << intValue
    }
    return bs
}

// Union returns the union o two bitsets bs1 ∪ bs2
func Union(bs1 Bitset, bs2 Bitset) Bitset {
    return bs1 | bs2
}

// SetMinus returns the set-difference bs1\bs2
func SetMinus(bs1 Bitset, bs2 Bitset) Bitset {
    return bs1 &^ bs2
}

// Intersection returns the intersection of two bitsets bs1 ∩ bs2
func Intersection(bs Bitset, elementOrSet Bitset) Bitset {
    return bs & elementOrSet
}

// IsSubsetOf returns true if b1 is subset of bs2 (or equal to b2)
func (bs1 Bitset) IsSubsetOf(bs2 Bitset) bool {
    return bs2&bs1 == bs1
}

// IsSupersetOf returns true if b1 is a superset of (or equal to) b2
func (bs1 Bitset) IsSupersetOf(bs2 Bitset) bool {
    return bs2&bs1 == bs2
}

// IsStrictSubsetOf returns true if b1 is subset of bs2 but not equal to b2
func (bs1 Bitset) IsStrictSubsetOf(bs2 Bitset) bool {
    return bs2&bs1 == bs1 && bs2 != bs1
}

// IsStrictSupersetOf returns true if b1 is a superset of b2 but not equal to b2
func (bs1 Bitset) IsStrictSupersetOf(bs2 Bitset) bool {
    return bs2&bs1 == bs2 && bs2 != bs1
}

// GetInts returns all integer values of a bitset
func (bs Bitset) GetInts() []int {
    intValues := make([]int, bits.OnesCount64(uint64(bs)))
    idx := 0
    for bs > 0 {
        intValue := bits.TrailingZeros64(uint64(bs))
        element := Bitset(1) << intValue
        intValues[idx] = intValue
        bs &^= element
        idx++
    }
    return intValues
}

// GetSmallestInt returns the smallest integer of a bitset.
func (bs Bitset) GetSmallestInt() int {
    if bs == 0 {
        panic("bitset is empty and has no smallest integer")
    }
    return bits.TrailingZeros64(uint64(bs))
}

// GetElements returns all elements of a bitsets. "Elements" are equivalent to
// to singleton sets containing one and only one integer.
func (bs Bitset) GetElements() []Bitset {
    elements := make([]Bitset, bits.OnesCount64(uint64(bs)))
    idx := 0
    for bs > 0 {
        intValue := bits.TrailingZeros64(uint64(bs))
        element := Bitset(1) << intValue
        elements[idx] = element
        idx++
        bs &^= element
    }
    return elements
}

// Len returns the number of elements in a set
func (bs Bitset) Len() int {
    return bits.OnesCount64(uint64(bs))
}

// GetIterator returns an iterator to loop across all elements of a set:
//
//  it := set.GetIterator()
//  for it.HasNext() {
//      fmt.Printf("Value: %d\n", it.NextInt())
//  }
//
// or after
//
//  resultSet := EmptySet()
//  it := set.GetIterator()
//  for it.HasNext() {
//      resultSet |= it.Next()
//  }
//
// resultSet and set coincide.
func (bs Bitset) GetIterator() *Iterator {
    it := Iterator(bs)
    return &it
}

// HasNext returns true if there is at least one element left to iterate
func (it *Iterator) HasNext() bool {
    return *it > 0
}

// Next returns the next element of the bitset for iteration
func (it *Iterator) Next() Bitset {
    intValue := bits.TrailingZeros64(uint64(*it))
    element := Bitset(1) << intValue
    *it &^= Iterator(element)
    return element
}

// NextInt returns the next integer contained in the bitset
func (it *Iterator) NextInt() int {
    intValue := bits.TrailingZeros64(uint64(*it))
    *it &^= Iterator(1) << intValue
    return intValue
}

// NumRemaining returns the remaining number of elements to iterate across
func (it *Iterator) NumRemaining() int {
    return bits.OnesCount64(uint64(*it))
}

// RemainingSet returns a bitset containing the remaining elements to loop across
func (it *Iterator) RemainingSet() Bitset {
    return Bitset(*it)
}
