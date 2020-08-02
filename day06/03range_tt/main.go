package main

import (
	"fmt"
)

func chData(new int, sp *[3][]int) {
	sp[new] = []int{1, 2, new}
}

func main() {
	// os.FileMode
	var tv = [3][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	// var t2 = [3]int{1, 2, 3}
	for i, v := range &tv {
		fmt.Printf("before:%p\n", v)
		// v[2] = i
		v = []int{1, 2, i}
		fmt.Printf("after:%p\n", v)
		// fmt.Println(v)
	}
	// for i := 0; i < len(tv); i++ {
	// 	chData(i, &tv)
	// }
	fmt.Println("++++++++++++++++++++++++++++++++++++++")
	for _, v := range tv {
		// fmt.Printf("%p\n", &v)
		fmt.Printf("%v\n", v)
	}
}
