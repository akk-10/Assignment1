package Salesperson

type Salesperson struct {
	name     string
	position string
	salary   float64
	address  string
}

func NewSalesperson(name, position, address string, salary float64) *Salesperson {
	return &Salesperson{
		name:     name,
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (s *Salesperson) IsName() string {
	return s.name
}

func (s *Salesperson) IsPosition() string {
	return s.position
}

func (s *Salesperson) GetSalary() float64 {
	return s.salary
}

func (s *Salesperson) GetAddress() string {
	return s.address
}

func (s *Salesperson) SetPosition(position string) {
	s.position = position
}

func (s *Salesperson) SetSalary(salary float64) {
	s.salary = salary
}

func (s *Salesperson) SetAddress(address string) {
	s.address = address
}
