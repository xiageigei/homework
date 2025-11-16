package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("name: %s , Age %d , EmployeeID %s\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	e := Employee{
		Person:     Person{Name: "zx", Age: 5},
		EmployeeID: "1234",
	}
	e.PrintInfo()
}
