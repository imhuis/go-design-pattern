package _decorator

type Inter interface {
	Operation() error
}

type Component struct {
	i Inter
}

func (d *Component) Operation() error {
	return d.i.Operation()
}
