package main

import "fmt"

func change(a []int) {
	a[0] = 100
}

func changeByRef(a *[]int) {
	*a = append(*a, 100)
}

func main() {

	var list1 = []int{10, 20, 30, 40}
	var list2 = []int{10, 20, 30, 40}

	change(list1)
	changeByRef(&list2)

	fmt.Println(list1)
	fmt.Println(list2)

}
