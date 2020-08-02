package main

import "fmt"

func Pingpong(s []int){
	s = append(s,3)
}

func main() {
	s := make([]int, 0)
	fmt.Println(s)
	Pingpong(s)
	fmt.Println(s)
}