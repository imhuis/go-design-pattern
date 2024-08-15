package example

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckupDone bool
	medicineDone      bool
	paymentDone       bool
}

type Department interface {
	execute(*Patient)
	setNext(department Department)
}

type Registration struct {
	next Department
}

func (r *Registration) execute(p *Patient) {
	if p.registrationDone {
		r.next.execute(p)
		return
	}
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Registration) setNext(department Department) {
	r.next = department
}

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckupDone {
		d.next.execute(p)
		return
	}
	p.doctorCheckupDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(department Department) {
	d.next = department
}

type Medicine struct {
	next Department
}

func (m *Medicine) execute(p *Patient) {
	if p.medicineDone {
		m.next.execute(p)
		return
	}
	p.medicineDone = true
	m.next.execute(p)
}

type Payment struct {
	next Department
}

func (c *Payment) execute(p *Patient) {
	if p.paymentDone {
		c.next.execute(p)
		return
	}
	p.paymentDone = true
	c.next.execute(p)
}
