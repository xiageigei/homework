package main

import "fmt"

func main() {
	a := 10
	point1(&a)
	fmt.Println(a)
	array := &[]int{1, 2, 3, 4}
	point2(array)
	fmt.Println(array)
}

func point1(param *int) {
	*param += 10
}
func point2(array *[]int) {
	for i, v := range *array {
		(*array)[i] += v * 2
	}
}
