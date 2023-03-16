package xsort

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Asc[T constraints.Ordered](slices []T) {
	sort.Slice(slices, func(i, j int) bool {
		return slices[i] < slices[j]
	})
}

func Desc[T constraints.Ordered](slices []T) {
	sort.Slice(slices, func(i, j int) bool {
		return slices[i] > slices[j]
	})
}

// Sortable[T any] is an interface that defines a type T and a function Less that takes another type T as an argument and returns a boolean.
// The Less function is used to compare two values of type T and determine which one is "less" than the other.
type Sortable[T any] interface {
	Less(other T) bool
}

// SliceAsc is a function that takes in a slice of type T and sorts it in ascending order. The type T must implement the Sortable interface, which requires the Less() method to be defined.
// The function uses the sort.Slice() method from the sort package to sort the slice, using a comparison function that compares two elements of the slice using their Less() methods.
func SliceAsc[T Sortable[T]](slices []T) {
	sort.Slice(
		slices,
		func(i, j int) bool {
			return slices[i].Less(slices[j])
		},
	)
}

// SliceDesc is a function that takes in a slice of type T, where T is a type that implements the Sortable interface.
// The function sorts the slice in descending order using the Less method from the Sortable interface.
func SliceDesc[T Sortable[T]](slices []T) {
	sort.Slice(
		slices,
		func(i, j int) bool {
			return !slices[i].Less(slices[j])
		},
	)
}
