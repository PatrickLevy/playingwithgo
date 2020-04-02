package utilities

import (
	"fmt"
	"math"
	"strconv"
)

// Simple passing param to function
func SayHello(message string) {
	fmt.Println(message)
}

// For Loop
func RunAForLoop() {
	fmt.Println("**** Going to run a for loop *****")
	for i := 0; i < 10; i++ {
		SayHello(strconv.Itoa(i))
	}
}

// Loop over elements in an array and a slice
func LoopThroughAnArrayAndSlice() {
	fmt.Println("**** Going to loop through elements in an array *****")

	// arrays are initilized by indicating how many elements (or using ... to let the compiler figure it out)
	myArray := [5]int{5, 10, 15, 20, 25}
	mySecondArray := [...]string{"first thing", "second thing"}
	fmt.Println("myArray:", myArray)
	fmt.Println("mySecondArray:", mySecondArray)

	// slice is initialized by NOT indicating how many elements
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println("mySlice", mySlice)

	// use range to loop
	for i, val := range mySlice {
		fmt.Println("looping:", i, val)
	}

	for char, c := range "This is a string" {
		fmt.Println("char", char, c)
	}
}

// Different variable types
func DeclareDifferentVarTypes() {
	fmt.Println("**** Going to declare a bunch of variable types *****")

	// boolean
	myBoolean := true
	const myBooleanConstant bool = true
	fmt.Println("myBoolean", myBoolean)
	fmt.Println("myBooleanConstant", myBooleanConstant)

	// int
	myInteger := 7
	var myInteger2 int = 12
	const myIntegerConstant int = 17
	fmt.Println("myInteger", myInteger)
	fmt.Println("myInteger2", myInteger2)
	fmt.Println("myIntegerConstant", myIntegerConstant)

	// float
	myFloat := 3.14
	fmt.Println("myFloat", myFloat)

	// string
	myString := "This is a string type"
	fmt.Println("myStrinbg", myString)

	// map
	// note, the element types must all be of a single type (all ints, strings, etc.)
	myMap := map[string]int{"cost1": 300}
	fmt.Println("myMap", myMap)
	fmt.Println("myMap['cost1']", myMap["cost1"])
	fmt.Println("does cost1 == 300?", myMap["cost1"] == 300)

	// struct
	type person struct {
		name string
		age  int
	}

	me := person{name: "Pat", age: 43}
	fmt.Println("me", me)
	fmt.Println("me.name", me.name)
	fmt.Println("me.age", me.age)

	// pointer
	var someInteger int = 3
	var myIntPointer *int
	myIntPointer = &someInteger
	fmt.Println("myIntPointer", *myIntPointer)

	var someString string = "yo this is the original value of someString"
	var myStringPointer *string = &someString // myStringPointer := &someString
	*myStringPointer = "hello there from a pointer"
	fmt.Println("*myStringPointer", *myStringPointer)
}

// assign anon function to a variable inside another function
func PlayWithAnonymousFunction() {
	fmt.Println("**** Going to play with anonymous functions *****")

	// Here is an anonymous func
	sayHi := func() bool {
		fmt.Println("Hello from the anonymous Function!")
		return true
	}

	sayHi()
}

// pass an instance of a struct to a func
func PlayingWithFunctionsThatTakeInStructs() {
	fmt.Println("**** Going to play with methods on types *****")

	// declare a dog struct
	type Dog struct {
		name string
		age  int
	}

	// make a generic bark method that takes in an instance of a dog
	bark := func(d Dog) bool {
		fmt.Println(d.name, "says:", "Woof!")
		return true
	}

	// now play with the dog and bark method
	espen := Dog{name: "Espen", age: 3}
	bark(espen)
}

// methods - note that these may not be defined within another function as far as I can tell
type Car struct {
	make  string
	model string
	year  int
}

func (c Car) honk() bool {
	fmt.Println(c.make, c.model, "says: HONK!!!!")
	return true
}

func PlayWithMethodsOnStructs() {
	fmt.Println("****** Playing with Methods ********")
	vibe := Car{make: "Pontiac", model: "Vibe", year: 2005}
	vibe.honk()
}

// composition of structs
type Vehicle struct {
	engineSize int
	brakes     bool
}

type Automobile struct {
	Vehicle
	numberOfSeats int
}

type Motorcycle struct {
	Vehicle
}

func PlayWithComposition() {
	// note: we can't seem to initialize with all the values, we need to do it after the myAuto variable is initialized
	myAuto := Automobile{}
	myAuto.engineSize = 5000
	myAuto.brakes = true
	myAuto.numberOfSeats = 5
	fmt.Println("myAuto", myAuto)
}

// types
func PlayWithTypes() {
	// we can make a new type from an existing type
	type aNewBooleanType bool
	var someNumber aNewBooleanType = true
	// someNumber = true
	fmt.Println("someNumber", someNumber)
}

/*********************
// Interfaces
*********************/

// first define a couple of new types
type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// methods for rectangles
func (r rectangle) getArea() float64 { // operates on rectangles, name is "getArea", takes in no parameters, returns a float64
	return r.width * r.height
}

func (r rectangle) getPerimeter() float64 {
	return 2*r.width + 2*r.height
}

// methods for circles
func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

// define an interface
type geometry interface {
	getArea() float64
	getPerimeter() float64
}

// define a function that takes advantage of the interface,
// it will automatically use the correct getArea() or getPerimeter() based on the
// type of shape passed in
func measure(g geometry) {
	fmt.Println("Area: ", g.getArea())
	fmt.Println("Perimeter", g.getPerimeter())
}

func PlayWithInterfaces() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	// Note that we can just use the methods defined for shapes directly
	rectArea := r.getArea()
	fmt.Println("rectArea", rectArea)

	// We can also write funtions that take advantage of the interface
	measure(r)
	measure(c)

}