package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [10]int{}
	// pArr := new([10]int)
	// fmt.Printf("pArr type:%T\n", pArr)
	// pArr[2] = 10
	// pArr[3] = 11
	// fmt.Printf("arr:%T len:%d arr:%v\n", arr, len(arr), pArr)
	arr2 := []int{1, 3}
	arr3 := append(arr2[:1], arr2[2:]...)
	fmt.Println(arr3)
	if reflect.TypeOf(arr) == reflect.TypeOf(arr2) {
		fmt.Println("类型一致")
	} else {
		fmt.Println("类型不一致")
	}

}
