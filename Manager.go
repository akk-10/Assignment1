package Manager

type Manager struct {
	name     string
	position string
	salary   float64
	address  string
}

func NewManager(name, position, address string, salary float64) *Manager {
	return &Manager{
		name:     name,
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (m *Manager) IsName() string {
	return m.name
}

func (m *Manager) IsPosition() string {
	return m.position
}

func (m *Manager) GetSalary() float64 {
	return m.salary
}

func (m *Manager) GetAddress() string {
	return m.address
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}

func (m *Manager) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}
