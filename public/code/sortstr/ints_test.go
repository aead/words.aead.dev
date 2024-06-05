package strsort_test

import (
	"math/rand"
	"slices"
	"sort"
	"testing"
)

const size = 10000

func BenchmarkSlicesSort_Int(b *testing.B) {
	rand := random(size)
	nums := make([]int, size)
	copy(nums, rand)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slices.Sort(nums)

		b.StopTimer()
		copy(nums, rand)
		b.StartTimer()
	}
}

func BenchmarkSort_Ints(b *testing.B) {
	rand := random(size)
	nums := make([]int, size)
	copy(nums, rand)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(nums)

		b.StopTimer()
		copy(nums, rand)
		b.StartTimer()
	}
}

// random returns a slice of len n containing random elements.
func random(n int) []int {
	v := make([]int, n)
	for i := range v {
		v[i] = rand.Int()
	}
	return v
}
