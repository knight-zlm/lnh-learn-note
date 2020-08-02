package main

import (
	"fmt"
)

const (
	jk = iota
	kl
	hk
)

func main() {
	// src := "hello 你好，world 世界！"
	// var hzc int
	// for _, v := range src {
	// 	if unicode.Is(unicode.Han, v) {
	// 		hzc++
	// 	}
	// }
	// fmt.Println("汉字数量：", hzc)
	totalNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range totalNums {
		for i := 1; i <= v; i++ {
			fmt.Printf("%dx%d=%d,\t", i, v, i*v)
		}
		fmt.Printf("\n")
	}

}
