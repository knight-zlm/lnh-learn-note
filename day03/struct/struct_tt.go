package main

import "fmt"

type human struct {
	Sex  int
	Name string
	Age  int
}

type student struct {
	human
	StudentID int
}

type teacher struct {
	human
	Course string
	Sex    int
}

type intm int

func main() {
	// st := student{StudentID: 1, human: human{Sex: 1, Name: "blb", Age: 12}}
	// fmt.Println("student ", st)
	// tc := teacher{Course: "语文", Sex: 1, human: human{Sex: 0, Name: "Mary", Age: 25}}
	// fmt.Println("teacher ", tc)
	// fmt.Printf("teacher Sex:%v", tc.human.Sex)
	i1 := intm(0)
	i1.Increace()
	fmt.Println("i1 ", i1)

}

func (i *intm) Increace() {
	for ; *i < intm(100); *i++ {
	}
}
