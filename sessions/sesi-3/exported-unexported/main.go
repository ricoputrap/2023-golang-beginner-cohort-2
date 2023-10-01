package main

import helpers "exported-unexported/helpers"

func main() {
	helpers.Greet();

	person := helpers.Person{}
	person.InvokeGreet();
}