package _iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayInt_Iterator(t *testing.T) {
	array := ArrayInt{1, 2, 3, 4, 5}
	iterator := array.Iterator()

	i := 0
	for iterator.HasNext() {
		assert.Equal(t, array[i], iterator.GetItem())
		iterator.Next()
		i++
	}
}
