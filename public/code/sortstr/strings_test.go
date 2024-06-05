package strsort_test

import (
	"math/rand"
	"slices"
	"sort"
	"testing"
)

const (
	size   = 10000
	strLen = 120
)

func BenchmarkSlicesSort_String(b *testing.B) {
	rand := random(size, strLen)
	strs := make([]string, size)
	copy(strs, rand)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slices.Sort(strs)

		b.StopTimer()
		copy(strs, rand)
		b.StartTimer()
	}
}

func BenchmarkSort_Strings(b *testing.B) {
	rand := random(size, strLen)
	strs := make([]string, size)
	copy(strs, rand)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Strings(strs)

		b.StopTimer()
		copy(strs, rand)
		b.StartTimer()
	}
}

// random returns a []string of n strings.
//
// Each string is length characters long and
// contains length/2 leading 'a' characters
// followed by random characters between 'a' and 'z'.
func random(n, length int) []string {
	v := make([]string, n)
	for i := range v {
		b := make([]byte, length)
		for j := 0; j < length/2; j++ {
			b[j] = byte('a')
		}
		for j := length / 2; j < length; j++ {
			b[j] = byte('a' + rand.Intn('z'-'a'))
		}
		v[i] = string(b)
	}
	return v
}
