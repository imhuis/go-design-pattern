package example

type SensitiveWordFilter interface {
	Filter(text string) bool
}

type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

func (c *SensitiveWordFilterChain) AddFilter(txt string) bool {
	for _, f := range c.filters {
		if f.Filter(txt) {
			return true
		}
	}
	return false
}

type AdSensitiveWordFilter struct{}

func (a *AdSensitiveWordFilter) Filter(text string) bool {
	return false
}

type PoliticalWordFilter struct{}

func (p *PoliticalWordFilter) Filter(text string) bool {
	return false
}
