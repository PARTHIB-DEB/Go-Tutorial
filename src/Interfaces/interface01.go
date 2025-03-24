package main

import "fmt"

/*

Interfaces are the real star while designing any complex architechture in Go.
While Go is not object-oriented but you can use the combination of Interfaces and Struct to make a similar design

An interface consists of several method signatures which will be applicable to those structs which will adhere that interface
Like Java , here `Structs Implements Interfaces` , but how ?
The struct's pointer-object or value-object is placed in the RECEIVER position of each of those methods.

But Is that All ? No
Till now we would just implement those methods by giving them different descriptions based on different structs

But how we will access those methods , because they are stacked in a interface , so somehow we have to access the interface ,
each time by an object of a struct.

Now either we can call each of those methods from the objects of each struct Or we can do this

We have to make a seperate function which will take an object of the interface (which will be actually an object of struct) and use its functions
It will be a stack of methods call using that function

*/

type details interface {
	Showspeed() int
	Showcolor() string
	Showwheels() int
}

type car struct {
	name     string
	showroom string
	speed    int
	color    string
	wheels   int
}

type bike struct {
	name     string
	showroom string
	speed    int
	color    string
	wheels   int
}

func (carobj *car) Showspeed() int {
	return carobj.speed
}
func (carobj *car) Showcolor() string {
	return carobj.color
}
func (carobj *car) Showwheels() int {
	return carobj.wheels
}

func (bikeobj *bike) Showspeed() int {
	return bikeobj.speed
}
func (bikeobj *bike) Showcolor() string {
	return bikeobj.color
}
func (bikeobj *bike) Showwheels() int {
	return bikeobj.wheels
}

// By default its taking the pointer instance of interface (now if a struct will implement it , its pointer instance will go as parameter)
func Showdetails(obj details) {
	fmt.Println(obj.Showcolor())
	fmt.Println(obj.Showspeed())
	fmt.Println(obj.Showwheels())
}

func main() {

	Bobj := bike{
		name:     "honda",
		showroom: "dunlop",
		speed:    12,
		color:    "blue",
		wheels:   2,
	}

	Cobj := car{
		name:     "hyundai",
		showroom: "dunlop",
		speed:    40,
		color:    "red",
		wheels:   4,
	}

	Cobj.Showcolor()  // Individual Call
	Bobj.Showwheels() // Individual Call

	Showdetails(&Bobj)
	Showdetails(&Cobj)

}
