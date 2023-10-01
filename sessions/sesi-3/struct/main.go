package main

import (
	"fmt"
	"strings"
)

// struct biasa
type Employee struct {
	name			string
	age				int
	position	string
}

// embedded struct
type Person struct {
	name  string
	age		int
}

type Member struct {
	position	string
	person		Person
}


func main() {

	// ===== init struct =====
	fmt.Println("===== 1.1. INIT STRUCT =====")

	var emp2 = Employee{
		name: "Rico",
		age: 25,
		position: "Frontend Web Dev",
	}

	fmt.Printf("Employee 1: %+v \n", emp2);

	fmt.Println(strings.Repeat("-", 20));

	// ===== pointer to struct =====
	fmt.Println("===== 1.2. POINTER TO STRUCT =====")

	var emp3 *Employee = &emp2;	// emp2 ngasih (shares) alamat memory dgn emp3

	fmt.Printf("Employee 2 name: %+v \n", emp2.name);
	fmt.Printf("Employee 3 name: %+v \n", emp3.name);

	fmt.Println(strings.Repeat("-", 20));

	// akan mengubah value dari emp2.name juga
	// karena emp2 dan emp3 saling berbagi alamat memori yg sama
	emp3.name = "Putra";

	fmt.Printf("Employee 2 name: %+v \n", emp2.name);
	fmt.Printf("Employee 3 name: %+v \n", emp3.name);

	fmt.Println(strings.Repeat("-", 20));

	// ===== embedded struct =====
	fmt.Println("===== 1.3. EMBEDDED STRUCT =====")

	var member1 = Member{}

	member1.person.name = "Rico"
	member1.person.age = 25
	member1.position = "Frontend Web Dev"

	fmt.Printf("%+v \n", member1)

	// anonymous struct (deklarasi tipe saat init variabelnya)
	fmt.Println("===== 1.3. ANONYMOUS STRUCT =====")

	// tanpa pengisian field
	var member2 = struct {
		person		Person
		position	string
	}{};
	member2.person = Person{ name: "Putra", age: 25 }
	member2.position = "Frontend Web Dev";

	// dgn pengisian field
	var member3 = struct {
		person		Person
		position	string
	} {
		person: Person{ name: "Pradana", age: 25 },
		position: "Frontend Web Dev",
	}

	fmt.Printf("Member 2: %+v \n", member2);
	fmt.Printf("Member 3: %+v \n", member3);

	// slice to struct
	fmt.Println("===== 1.4. SLICE TO STRUCT =====")

	var people = []Person{
		{ name: "Agus", age: 30 },
		{ name: "Bambang", age: 40 },
		{ name: "Caca", age: 20 },
	}

	for _, value := range people {
		fmt.Printf("%+v \n", value);
	}

	// slice of anonymous structs
	fmt.Println("===== 1.5. SLICE OF ANONYMOUS STRUCT =====")
	var members = []struct {
		person		Person
		position	string
	} {
		{ person: Person{ name: "Dodi", age: 20 }, position: "Engineer" },
		{ person: Person{ name: "Enggar", age: 20 }, position: "Engineer" },
	}

	for _, value := range members {
		fmt.Printf("%+v \n", value);
	}
}