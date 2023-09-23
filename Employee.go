package Employee

type Employee interface {
	IsName() string
	IsPosition() string
	GetSalary() float64
	GetAddress() string
	SetPosition(position string)
	SetSalary(salary float64)
	SetAddress(address string)
}
