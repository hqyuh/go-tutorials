package main

import "fmt"

type Emp int

type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var emp Employee
	emp = Emp(1)
	emp.PrintName("Huy")
	fmt.Println("Employee Salary:", emp.PrintSalary(25000, 5))
}
