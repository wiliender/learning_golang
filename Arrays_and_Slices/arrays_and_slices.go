package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	var arry1 [5]int
	fmt.Println(arry1)

	var array2 [5]string
	array2[0] = "Wiliender"
	fmt.Println(array2)

	array3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	array4 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array4)

	array5 := [...]int{1, 2, 3, 4, 5} //size is not specified
	fmt.Println(array5)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	slice = append(slice, 16) //adds 16 to the end of the slice
	fmt.Println(slice)

	slice2 := array4[1:3]
	fmt.Println(slice2)

	array4[1] = "Posição Alterada"
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	// ARRAYS INTERNOS
	slice3 := make([]float32, 14, 15) //size 10, capacity 15
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 7)

	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacity

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
