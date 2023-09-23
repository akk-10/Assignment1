package Administrator

type Administrator struct {
	name     string
	position string
	salary   float64
	address  string
}

func NewAdministrator(name, position, address string, salary float64) *Administrator {
	return &Administrator{
		name:     name,
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (a *Administrator) IsName() string {
	return a.name
}

func (a *Administrator) IsPosition() string {
	return a.position
}

func (a *Administrator) GetSalary() float64 {
	return a.salary
}

func (a *Administrator) GetAddress() string {
	return a.address
}

func (a *Administrator) SetPosition(position string) {
	a.position = position
}

func (a *Administrator) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Administrator) SetAddress(address string) {
	a.address = address
}
