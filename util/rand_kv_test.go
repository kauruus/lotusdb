package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTestKey(t *testing.T) {
	for i := 0; i < 10; i++ {
		assert.NotNil(t, string(GetTestKey(i)))
	}
}

func TestRandomValue(t *testing.T) {
	for i := 0; i < 10; i++ {
		assert.NotNil(t, string(RandomValue(10)))
	}
}

func BenchmarkRandomValue(b *testing.B) {
	size := []int{10, 100, 1000}

	for _, s := range size {
		b.Run(fmt.Sprintf("input size %d", s), func(b *testing.B) {
			b.RunParallel(func(b *testing.PB) {
				for b.Next() {
					RandomValue(s)
				}
			})
		})
	}
}
