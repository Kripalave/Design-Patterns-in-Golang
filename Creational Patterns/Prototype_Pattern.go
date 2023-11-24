package main

import "fmt"

// Cloneable defines the prototype interface
type Cloneable interface {
	Clone() Cloneable
}

// Employee represents an employee
type Employee struct {
	Name   string
	Role   string
	Salary float64
}

// NewEmployee creates a new instance of Employee
func NewEmployee(name, role string, salary float64) *Employee {
	return &Employee{Name: name, Role: role, Salary: salary}
}

// Clone creates a shallow copy of the Employee
func (e *Employee) Clone() Cloneable {
	// Perform a shallow copy to create a new instance with the same values
	return &Employee{Name: e.Name, Role: e.Role, Salary: e.Salary}
}

func main() {
	// Creating an instance of Employee
	originalEmployee := NewEmployee("Sravan Vedantam", "Developer", 10000.0)

	// Cloning the original employee
	clonedEmployee := originalEmployee.Clone().(*Employee)

	// Modifying the cloned employee
	clonedEmployee.Name = "Srvan Vedantam"
	clonedEmployee.Salary = 118000.0

	// Displaying the results
	fmt.Println("Original Employee:", originalEmployee)
	fmt.Println("Cloned Employee:", clonedEmployee)
}
