package Engineer

type Engineer struct {
	name     string
	position string
	salary   float64
	address  string
}

func NewEngineer(name, position, address string, salary float64) *Engineer {
	return &Engineer{
		name:     name,
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (e *Engineer) IsName() string {
	return e.name
}

func (e *Engineer) IsPosition() string {
	return e.position
}

func (e *Engineer) GetSalary() float64 {
	return e.salary
}

func (e *Engineer) GetAddress() string {
	return e.address
}

func (e *Engineer) SetPosition(position string) {
	e.position = position
}

func (e *Engineer) SetSalary(salary float64) {
	e.salary = salary
}

func (e *Engineer) SetAddress(address string) {
	e.address = address
}
