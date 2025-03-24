package main

import "fmt"

/*
No Inheritance in Go --

Declaring a type based on another type looks a bit like inheritance, but it isn’t.
The two types have the same underlying type, but that’s all. There is no hierarchy between these types.

In languages with inheritance, a child instance can be used anywhere the parent instance is used. The
child instance also has all the methods and data structures of the parent instance. That’s not the case
in Go. You can’t assign an instance of type 'HighScore' to a variable of type 'Score' or vice versa without
a type conversion,

While Go doesn’t have inheritance, it encourages code reuse via built-in support for composition and promotion:
*/

type Employee struct {
	Name string
	ID   int
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

/*
Note that Manager contains a field of type Employee, but no name is assigned to that field. This
makes employees an embedded field. Any fields or methods declared on an embedded field are
promoted to the containing struct and can be invoked directly on it.

see 'Manager' struct below --
*/
type Manager struct {
	Employee // embedded field
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	newEmployees := []Employee{}
	for _, emp := range m.Reports {
		newEmployees = append(newEmployees, emp)
	}
	return newEmployees
}

func main() {

	emp := []Employee{
		Employee{
			Name: "Aman",
			ID:   123,
		},
		Employee{
			Name: "Bimal",
			ID:   456,
		},
	}

	m := Manager{
		Employee: Employee{Name: "Bob Bobson", ID: 12345},
		Reports:  emp,
	}
	fmt.Printf("%v", m.ID)       // prints 12345
	fmt.Println(m.Description()) // prints Bob Bobson (12345)
	fmt.Println(m.FindNewEmployees())
}
