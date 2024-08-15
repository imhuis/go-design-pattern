package _chain

type Filter interface {
	Next() Filter
	Operation() error
}

type FilterA struct {
}

func (f FilterA) Operation() error {
	panic("implement me")
}

type FilterB struct {
}

func (f FilterB) Operation() error {
	panic("implement me")
}

type ChainFilter struct {
	filter []Filter
}

// AddFilter
func (c *ChainFilter) AddFilter(filter Filter) {
	c.filter = append(c.filter, filter)
}

// LinkFilter
func (c *ChainFilter) LinkFilter(filter Filter) {
	c.AddFilter(filter)
}

func (c *ChainFilter) DoFilter() error {
	for _, f := range c.filter {
		if err := f.Operation(); err != nil {
			return err
		}
	}
	return nil
}
