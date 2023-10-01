package main

import "fmt"

func main() {

}

// params with different types
func greet(name string, age int) {
	fmt.Printf("Name: %s, Age: %d", name, age);
}

// params with same type
func greet2(name, address string) {
	fmt.Printf("Name: %s, Address: %s", name, address);
}

// return a single value
func greet3(message string, names string) string {
	var result string = fmt.Sprintf("%s %s", message, names);
	return result;
}

// return multiple values
func calculate(num float64) (float64, float64) {
	multiple := num * num
	divided := num / num

	return multiple, divided
}

// predefine return values
func newCalculate(num float64) (multiple float64, divided float64) {
	multiple = num * num;
	divided = divided * divided;

	return;
}

type Person struct {
	name  string
	age		int
}

func getPerson(name string, age int) (person Person) {
	person = Person{ name, age }

	return;
}