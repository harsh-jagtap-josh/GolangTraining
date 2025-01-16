package main

import "fmt"

func main() {

	var array = []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	var index1, index2 int

	fmt.Scanf("%v %v", &index1, &index2)

	// If any of the Indexes is out of bounds

	if index1 < 0 || index1 >= len(array) || index2 < 0 || index2 >= len(array) {

		fmt.Println("Index out of Bounds")

	} else if index1 > index2 {

		// If first index is greater than second index

		fmt.Println("Incorrect Indexes!")

	} else {

		slice1 := array[:index1+1]

		slice2 := array[index1 : index2+1]

		slice3 := array[index2:]

		fmt.Println("Slice 1 is: ", slice1)
		fmt.Println("Slice 2 is: ", slice2)
		fmt.Println("Slice 3 is: ", slice3)
	}
}
