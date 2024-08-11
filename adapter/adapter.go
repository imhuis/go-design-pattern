package _adapter

// AdapterInterface is an interface for adapter
type AdapterInterface interface {
	Implement() error
}

type ClientA struct {
}

func (c *ClientA) Request() error {
	return nil
}

type Adapter struct {
	ClientA
}

func (a *Adapter) Implement() error {
	return a.Request() // Adapter implements the ClientA unique method
}

type ClientB struct {
}

func (c *ClientB) Create() error {
	return nil
}

type AdapterB struct {
	ClientB
}

func (b *AdapterB) Implement() error {
	return b.Create() // AdapterB implements the ClientB unique method
}
