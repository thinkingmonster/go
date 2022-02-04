package main

import "fmt"

type Employee struct {
	Name       string
	EmployeeId int
	Location   string
	Salary     int
}

func main() {
	jim := Employee{
		Name:       "Jim",
		EmployeeId: 183,
		Location:   "India",
		Salary:     50000,
	}

	var jack Employee
	jack.Name = "jack"
	jack.Salary = 20000
	jack.EmployeeId = 1300
	jack.Location = "Mohali"

	var employees []Employee
	employees = append(employees, jim)
	employees = append(employees, jack)

	for _, ele := range employees {
		fmt.Println(ele.Name)

	}

}
