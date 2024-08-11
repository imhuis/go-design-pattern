package _singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLazySingleton(t *testing.T) {
	assert.True(t, GetLazySingleton() == GetLazySingleton())
}

func BenchmarkGetLazySingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingleton() != GetLazySingleton() {
				b.Errorf("test failed")
			}
		}
	})
}
