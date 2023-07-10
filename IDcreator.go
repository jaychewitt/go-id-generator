package main

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

// This program is used to created an ID number for new Employees

func GetNewEmployee() {

	type DOB struct {
		month string
		day   int
		year  int
	}

	type Dob map[string]DOB

	type User struct {
		FirstName string
		LastName  string
		Age       int
		Dob
	}

	newDob := DOB{}
	newUser := User{Dob: make(map[string]DOB)}

	color.Green("Employee First Name: \n")
	fmt.Scanln(&newUser.FirstName)
	emp := newUser.FirstName

	color.Green("Employee Last Name: \n")
	fmt.Scanln(&newUser.LastName)

	color.Green("Employee Age: \n")
	fmt.Scanln(&newUser.Age)

	// Collect Employees D.O.B and making the map

	color.Green("Employee D.O.B (mm/dd/year) \n")
	color.Green("mm: ")
	fmt.Scanln(&newDob.month)
	color.Green("dd: ")
	fmt.Scanln(&newDob.day)
	color.Green("yy: ")
	fmt.Scanln(&newDob.year)

	newUser.Dob[emp] = DOB{month: newDob.month, day: newDob.day, year: newDob.year}

	createEmpId := func(name string, age int, dobDD int, dobM string, dobY int) string {
		newName := newUser.FirstName + newUser.LastName
		newAge := strconv.Itoa(newUser.Age)
		newDobb := strconv.Itoa(newDob.day)
		newDobM := newDob.month[0:2]
		newDobY := strconv.Itoa(newDob.year)

		return newName[0:2] + newAge + newDobM + newDobb  + newDobY[0:2]
	}

	newUSerID := createEmpId(newUser.FirstName, newUser.Age, newDob.day, newDob.month, newDob.year)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	color.Cyan("%v ID Number is %v", newUser.FirstName, newUSerID)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")

}

func main() {
	GetNewEmployee()
}
