package _composite

type Inter interface {
	Operation() error
}

type Leaf struct {
}

func (l *Leaf) Operation() error {
	//TODO implement me
	panic("implement me")
}

type Tree struct {
	Name  string
	Leafs []Inter
}

func (t *Tree) Operation() error {
	for _, leaf := range t.Leafs {
		leaf.Operation()
	}
	return nil
}

func (t *Tree) AddLeaf(leaf Inter) {
	t.Leafs = append(t.Leafs, leaf)
}

func NewTree() *Tree {
	return &Tree{Name: "root"}
}
