package main

import "fmt"

type People struct {
	Message string
}

func (d People) Hello() {
	fmt.Println(d.Message)
}

func main() {
	p1 := People{"안녕하세요"}
	p1.Hello()
}
