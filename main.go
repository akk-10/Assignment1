package main

import (
	"Assignment1/Administrator"
	"Assignment1/Designer"
	"Assignment1/Employee"
	"Assignment1/Engineer"
	"Assignment1/Manager"
	"Assignment1/Salesperson"
	"fmt"
)

func main() {
	manager := Manager.NewManager("Maral", "Manager", "123 Main St", 150000.0)
	administrator := Administrator.NewAdministrator("Nazira", "Administrator", "123 Olivia St", 600000.0)
	salesperson := Salesperson.NewSalesperson("Aru", "Salesperson", "123 Train St", 200000.0)
	engineer := Engineer.NewEngineer("Anel", "Engineer", "123 Abay St", 250000.0)
	designer := Designer.NewDesigner("Ayala", "Designer", "123 Soul St", 180000.0)

	printEmployeeInfo(engineer)
	engineer.SetSalary(850000)
	printEmployeeInfo(engineer)

	printEmployeeInfo(salesperson)
	salesperson.SetPosition("Senior Sales Manager")
	printEmployeeInfo(salesperson)

	printEmployeeInfo(manager)
	manager.SetPosition("Senior Engineering Manager")
	manager.SetSalary(950000)
	manager.SetAddress("New Manager Address")
	printEmployeeInfo(manager)

	printEmployeeInfo(designer)
	printEmployeeInfo(administrator)
}

func printEmployeeInfo(e Employee.Employee) {
	fmt.Printf("Employee Name: %s\n", e.IsName())
	fmt.Printf("Employee Position: %s\n", e.IsPosition())
	fmt.Printf("Employee Salary: $%.2f\n", e.GetSalary())
	fmt.Printf("Employee Address: %s\n", e.GetAddress())
}
