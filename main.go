package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var name string = "Akash"
	fmt.Printf("This is my name %s\n", name)
	age := 25
	fmt.Printf("My age is %d\n", age)
	var city, country string = "G. Noida", "India"
	var (
		isEmployed bool   = true
		sal        int    = 60000
		position   string = "Junior Developer"
	)
	fmt.Printf("My city is %s, and country is %s\n", city, country)
	fmt.Printf("I am %t. My sal is %d and my position is %s\n", isEmployed, sal, position)

	// default values
	var defaultInt int
	var defaultFloat float32
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultBool, defaultFloat, defaultString, defaultInt)
	// following line prints absolute gibberish
	// fmt.Printf("%d %f %f %d %t %f %d %s\n", defaultInt, defaultInt, defaultFloat, defaultFloat, defaultBool, defaultBool, defaultBool, defaultString)
	// following line prints good
	fmt.Printf("%d %f %t %s\n", defaultInt, defaultFloat, defaultBool, defaultString)
	// const type variable can be unused, editor gives warning but code compiles
	// whereas code doesn't compile when var is unused
	const pi = 3.14

	// when using var type variable, we can omit the type as well
	var x = 7
	fmt.Println(x)

	// following line  gives error coz d is undefined
	// d = 9

	// finally we can can omit writing var if using := operrator
	z := 99
	fmt.Println(z)

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	const typedAge int = 123
	const unTypedAge = 123

	fmt.Println(typedAge == unTypedAge)

	const (
		Jan = iota + 1 // 1
		Feb            // 2
		Mar            // 3
	)

	fmt.Println(Jan, Feb, Mar)

	fmt.Println(add(25, 78))

	sum, product := calculateSumAndProduct(6, 7)
	fmt.Println(sum, product)

	// this gives an error as you have to accepts two values as the function returns two values
	// result := calculateSumAndProduct(9, 10)
	// you can skip the second returned value like this
	// and this is not declaring a variable with name _, coz i did it later in line _
	// its sort of skipping as a feature provided by language
	newSum, _ := calculateSumAndProduct(9, 10)
	println(newSum)
	// println(newSum, _) => this line says _ cannot be used as

	var _ int = 29 // this line is fine, and gives no compile error ;-)
	// fmt.Println(_) => this line says _ cannot be used as

	if age >= 18 {
		fmt.Println("you are an adult")
	} else {
		fmt.Println("you are not an adult")
	}

	if 99 > 100 && 55 == 60 {
		fmt.Println("Thats not possible")
	} else if 99 > 98 && 11 >= 10 {
		fmt.Println("that is possible")
	}

	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Meh")
		break
	case "Tuesday", "Wednesday", "Saturday":
		fmt.Println("Less Meh")
		break
	default:
		fmt.Println("Muh")
	}

	// notice what is printed
	switch day {
	case "Tuesday":
		fmt.Println("Print no. 1")
	case "Wednesday", "Saturday":
		fmt.Println("Print no. 2")
	default:
		fmt.Println("Muh")
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()

	// can do the same thing using
	for i := range 5 {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()

	// no while loop, alternative is loop with condition
	counter := 0
	for counter <= 20 {
		fmt.Printf("%d ", counter)
		counter += 2
	}
	fmt.Println()

	// explicit infinite loops
	j := 0
	for {
		if j >= 10 {
			break
		}
		fmt.Printf("%d ", j+1)
		j += 1
	}
	fmt.Println()

	// Arrays and Slices
	numbers := [5]int{10, 20, 90, 80, 10}
	fmt.Printf("%v\n", numbers) // prints array
	fmt.Println(numbers)        // this too
	fmt.Println(len(numbers))

	// you can omit the number of elements
	numbersWithoutCount := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbersWithoutCount)

	// matrix
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)

	// slices
	allNumbers := numbers[:]
	firstThree := numbers[0:3]
	fmt.Println(allNumbers)
	fmt.Println(firstThree)

	// creating slice from start
	fruits := []string{}
	fruits = append(fruits, "kiwi", "apple")
	fmt.Println(fruits)
	fruits = append(fruits, "banana", "cherry")
	fmt.Println(fruits)
	var moreFruits []string = []string{"mango", "pineapple", "guava"}
	// add one slice to another
	fruits = append(fruits, moreFruits...)
	fmt.Println(fruits)

	// use range to iterate over arrays
	for index, val := range numbers {
		fmt.Printf("(%d-%d) ", index, val)
	}
	fmt.Println()

	// map
	squares := map[int]int{
		1: 1,
		2: 4,
		3: 9,
		4: 16,
	}
	fmt.Println(squares[3])
	// if key not present, zero value of value type is returned
	fmt.Println(squares[99])
	fmt.Println(squares)
	// a way to check if exists, coz zero value cannot be relied upon
	value, exists := squares[100]
	fmt.Println(value, exists)
	value, exists = squares[4]
	fmt.Println(value, exists)

	delete(squares, 4)
	// deleting keys
	fmt.Println(squares)

	// structures, gives power to create your own custom data types!
	type Person struct {
		name string
		age  int
	}

	p1 := Person{name: "Akash", age: 50}
	fmt.Printf("%v\n", p1)

	// anonumous structs
	employee := struct {
		position string
		salary   int
	}{
		position: "Senior Engineer",
		salary:   50000,
	}
	fmt.Println(employee)

	// structs cannot be redefined in same scope

	// function with names inide functions not allowed
	// func modifyPerson(p Person) {
	// }

	modifyPerson := func(p Person) {
		p.name = "Billu"
		p.age = 55
		fmt.Println("Here its modified", p)
	}

	modifyPersonUsingRef := func(p *Person) {
		p.name = "Sikka"
		p.age = 60
		fmt.Println("Here its modified", p)
	}

	modifyPerson(p1)

	// hence proved by default a copy of struct is passed
	fmt.Println("Here its not modified", p1)

	modifyPersonUsingRef(&p1)
	fmt.Println("Now its modified", p1)

	fmt.Printf("Value of p1 is %v and address of p1 is %p\n", p1, &p1)

	addressOfP1 := &p1

	// straight from c language
	*addressOfP1 = Person{
		name: "Howard",
		age:  70,
	}

	fmt.Println(p1)
	fmt.Println(*addressOfP1)

	var lang Language = Language{
		name:            "English",
		countryOfOrigin: "England",
	}
	lang.langPrinter("The country of origin of")
}

type Language struct {
	name            string
	countryOfOrigin string
}

// attaching function to struct, OOPs features
func (l *Language) langPrinter(msg string) {
	fmt.Printf("%s %s is %s.\n", msg, l.name, l.countryOfOrigin)
}

// two functions with same name i.e. langPrinter with diff arg list wont work, so no polymorphism

// this fn is not exported, you cannot omit arg types and return type
func add(a int, b int) int {
	return a + b
}

// you can specify the type of b and for the types before b int is inferred
// specify multiple return types in paranthesis
func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
