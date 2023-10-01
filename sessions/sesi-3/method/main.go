package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) introduce(msg string) string {
	return fmt.Sprintf("%s My name is %s and I am %d years old \n", msg, p.name, p.age);
}

// won't change the name
func (p Person) changeName1(name string) {
	p.name = name;
}

// pointer -> can change the name
func (p *Person) changeName2(name string) {
	p.name = name;
}

func main() {
	person := Person{ name: "Rico", age: 25 }
	fmt.Println(person.introduce("Helllooooo!"));

	person.changeName1("Putra");
	fmt.Println("Change name with first method - name:", person.name);

	person.changeName2("Pradana");
	fmt.Println("Change name with first method - name:", person.name);
}