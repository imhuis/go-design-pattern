package _adapter

// 提供一个接口，用于适配器的实现
// 适配器类组合被适配的类，实现了接口的方法

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
