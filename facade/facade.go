package _facade

type Facade struct {
	DependencyA *DependencyA
	DependencyB *DependencyB
}

func NewFacade() *Facade {
	return &Facade{
		DependencyA: newDependencyA(),
		DependencyB: newDependencyB(),
	}
}

func (f *Facade) Operation() {
	// facade operation
	panic("implement me")
}

type DependencyA struct {
}

func newDependencyA() *DependencyA {
	return &DependencyA{}
}

type DependencyB struct {
}

func newDependencyB() *DependencyB {
	return &DependencyB{}
}
