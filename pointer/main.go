package main

import "fmt"

func main() {
	num := "123"
	point := &num
	point2 := &point
	**point2 = "Change"
	fmt.Println(num)
}
