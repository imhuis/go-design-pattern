package _iterator

type Iterator interface {
	HasNext() bool
	Next()
	GetItem() interface{}
}

type ArrayInt []int

func (a ArrayInt) Iterator() Iterator {
	return &ArrayIterator{
		array: a,
		index: 0,
	}
}

type ArrayIterator struct {
	array ArrayInt
	index int
}

func (a *ArrayIterator) HasNext() bool {
	return a.index < len(a.array)-1
}

func (a *ArrayIterator) Next() {
	a.index++
}

func (a *ArrayIterator) GetItem() interface{} {
	return a.array[a.index]
}
