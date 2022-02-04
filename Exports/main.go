package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Dullagon", Salary: 30000, FullTime: true},
	{FirstName: "Garry", LastName: "Quint", Salary: 40000, FullTime: true},
	{FirstName: "Dave", LastName: "Smith", Salary: 50000, FullTime: true},
	{FirstName: "Cornor", LastName: "Quint", Salary: 60000, FullTime: true},
	{FirstName: "Williams", LastName: "Smith", Salary: 10000, FullTime: false},
	{FirstName: "Jose", LastName: "Zaramano", Salary: 15000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}
	log.Println(myStaff.All())
	log.Println("OverPaid staff members are: ", myStaff.OverPaid())
	log.Println("UnderPaid staff members are: ", myStaff.UnderPaid())
}
