package gobitset

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestBitset_BasicFuncs(t *testing.T) {
	assert.Equal(t, []int{0, 1, 63}, MakeSet(0, 1, 63).GetInts())
	assert.Equal(t, []int{}, MakeSet().GetInts())
	assert.Equal(t, Bitset(0b110), MakeSet(1, 2))
	assert.Equal(t, []int{1, 2, 3}, Union(MakeSet(1, 2), MakeSet(2, 3)).GetInts())
	assert.Equal(t, []int{2}, Intersection(MakeSet(1, 2), MakeSet(2, 3)).GetInts())
	assert.Equal(t, []int{1}, SetMinus(MakeSet(1, 2), MakeSet(2, 3)).GetInts())
	assert.Equal(t, []int{}, Union(MakeSet(), MakeSet()).GetInts())
	assert.Equal(t, []int{}, Intersection(MakeSet(), MakeSet()).GetInts())
	assert.Equal(t, []int{}, SetMinus(MakeSet(), MakeSet()).GetInts())
	assert.Equal(t, []int{1, 2, 3}, AddInts(MakeSet(1, 2), 2, 3).GetInts())
	assert.Equal(t, []int{1}, DeleteInts(MakeSet(1, 2), 2, 3).GetInts())
}

func TestBitset_GetSmallestInt(t *testing.T) {
	assert.Equal(t, 4, MakeSet(5, 4, 6).GetSmallestInt())
	assert.Panics(t, func() { EmptySet.GetSmallestInt() })
}

func TestBitset_GetElements(t *testing.T) {
	assert.Equal(t, []Bitset{}, EmptySet.GetElements())
	elts := MakeSet(4, 2, 3).GetElements()
	slices.Sort(elts)
	assert.Equal(t, []Bitset{0b100, 0b1000, 0b10000}, elts)
}

func TestIterator(t *testing.T) {
	values := []int{}
	set := MakeSet(2, 3, 4)
	it := set.GetIterator()
	for it.HasNext() {
		values = append(values, it.NextInt())
	}
	slices.Sort(values)
	assert.Equal(t, []int{2, 3, 4}, values)

	elements := []Bitset{}
	it = set.GetIterator()
	for it.HasNext() {
		elements = append(elements, it.Next())
	}
	slices.Sort(elements)
	assert.Equal(t, []Bitset{0b100, 0b1000, 0b10000}, elements)
}

func TestIterator_Remaining(t *testing.T) {
	baseSet := MakeSet(2, 3, 4)
	it := baseSet.GetIterator()
	assert.Equal(t, baseSet, it.RemainingSet())
	assert.Equal(t, 3, it.NumRemaining())
	elt := it.Next()
	assert.NotEqual(t, baseSet, it.RemainingSet())
	assert.Equal(t, 2, it.NumRemaining())
	assert.Equal(t, baseSet, Union(it.RemainingSet(), elt))
}

func TestBitset_IsSubsetOf(t *testing.T) {
	bsLarge := MakeSet(1, 2, 3)
	bsSmall := MakeSet(2, 3)
	assert.Equal(t, true, bsSmall.IsSubsetOf(bsLarge))
	assert.Equal(t, true, bsLarge.IsSubsetOf(bsLarge))
	assert.Equal(t, false, bsLarge.IsSubsetOf(bsSmall))

	assert.Equal(t, false, bsSmall.IsSupersetOf(bsLarge))
	assert.Equal(t, true, bsLarge.IsSupersetOf(bsLarge))
	assert.Equal(t, true, bsLarge.IsSupersetOf(bsSmall))

	assert.Equal(t, true, bsSmall.IsStrictSubsetOf(bsLarge))
	assert.Equal(t, false, bsLarge.IsStrictSubsetOf(bsLarge))
	assert.Equal(t, false, bsLarge.IsStrictSubsetOf(bsSmall))

	assert.Equal(t, false, bsSmall.IsStrictSupersetOf(bsLarge))
	assert.Equal(t, false, bsLarge.IsStrictSupersetOf(bsLarge))
	assert.Equal(t, true, bsLarge.IsStrictSupersetOf(bsSmall))
}

func TestBitset_Len(t *testing.T) {
	assert.Equal(t, 0, EmptySet.Len())
	assert.Equal(t, 3, MakeSet(2, 3, 4).Len())
	assert.Equal(t, 2, MakeSet(2, 2, 4).Len())
}
