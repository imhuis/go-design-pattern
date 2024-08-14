package example

import "fmt"

type Power interface {
	Output5V() int
}

type Voltage220V struct {
}

func (v *Voltage220V) Output220V() int {
	return 220
}

type Voltage5V struct {
}

func (v *Voltage5V) Output5V() int {
	return 5
}

type VoltageAdapter struct {
	Voltage220V
}

type Phone struct {
	VoltageAdapter
}

func (p *Phone) Charging() {
	voltage := p.Output220V() / 44
	fmt.Println(voltage)
}
