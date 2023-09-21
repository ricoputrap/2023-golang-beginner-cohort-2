package main

import "fmt"

func variables() {
	fmt.Println("====== VARIABLES ======")

	// with data type
	var name string = "Rico";
	fmt.Println("Hello, my name is", name);

	// without data type
	age := 25;
	fmt.Printf("I am %d years old!\n", age);

	// multiple declaration
	var jobTitle, specialization, techStack string = "Software Engineer", "Front-end Web Development", "React JS";
	fmt.Printf("I work as a %s focusing on %s using %s.\n", jobTitle, specialization, techStack);
}

func main() {
	variables();
}