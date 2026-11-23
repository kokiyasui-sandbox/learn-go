package main

import "fmt"

func main() {
	employee_1 := struct {
		id        int
		firstName string
		lastName  string
	}{
		id:        1,
		firstName: "koki",
		lastName:  "yasui",
	}

	employee_2 := Employee{
		id:        2,
		firstName: "koki",
		lastName:  "yasui",
	}

	var employee_3 struct {
		id        int
		firstName string
		lastName  string
	}

	employee_3.id = 3
	employee_3.firstName = "koki"
	employee_3.lastName = "yasui"

	fmt.Println(employee_1)
	fmt.Println(employee_2)
	fmt.Println(employee_3)
}

type Employee struct {
	id        int
	firstName string
	lastName  string
}
