package Designer

type Designer struct {
	name     string
	position string
	salary   float64
	address  string
}

func NewDesigner(name, position, address string, salary float64) *Designer {
	return &Designer{
		name:     name,
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (d *Designer) IsName() string {
	return d.name
}

func (d *Designer) IsPosition() string {
	return d.position
}

func (d *Designer) GetSalary() float64 {
	return d.salary
}

func (d *Designer) GetAddress() string {
	return d.address
}

func (d *Designer) SetPosition(position string) {
	d.position = position
}

func (d *Designer) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Designer) SetAddress(address string) {
	d.address = address
}
