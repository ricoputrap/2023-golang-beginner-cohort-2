package main

import "fmt"

func main() {
	// array -> fixed length
	var primeNumbers = [5]int{2, 3, 5, 7, 11};

	for i, num := range primeNumbers {
		fmt.Printf("Number at index %d, number is %d \n", i, num);
	}

	fmt.Println("========== MULTI DIMENSIONAL ARRAY ==========");
	// multidimentional array: row X column
	twoDimensions := [2][3]int{{1, 2, 3}, {4, 5, 6}} 
	for _, row := range twoDimensions {
		for _, col := range row {
			fmt.Printf("%d ", col);
		}
		fmt.Println();
	}

	fmt.Println("========== SLICE ==========");
	// slice -> mirip array tp non-fixed length
	// var techStacks = []string{"TypeScript", "ReactJS"};

	// define with type & size
	var techStacks = make([]string, 2);

	// modify at specific index
	techStacks[0] = "TypeScript";
	// techStacks[1] = "ReactJS"; // -> this index will be reserved (empty)

	// append
	techStacks = append(techStacks, "NodeJS", "Go");

	for i, stack := range techStacks {
		fmt.Printf("Stack ke-%d: %s \n", (i + 1), stack);
	}
	fmt.Println();

	// append with ellipsis
	var newStacks = []string{"Django", "Flask"};
	techStacks = append(techStacks, newStacks...);
	for i, stack := range techStacks {
		fmt.Printf("Stack ke-%d: %s \n", (i + 1), stack);
	}

	// copy functions: copy(arr1, arr2) -> replace contents of arr1 with arr2

	fmt.Println("========== X ==========");
	arr := []any{1, 2, "aku"};
	for i, item := range arr {
		fmt.Printf("Item ke-%d: %s \n", (i + 1), item);
	}

	/*
	Backing array: menyimpan element pada slice
	*/
}