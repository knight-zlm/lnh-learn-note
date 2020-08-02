package main

import "fmt"

func main() {
	a := make([]int, 4, 6)
	a1 := a[2:4]
	a2 := a[3:6]
	fmt.Printf("a:%v len:%d, cap:%d, id:%p\n", a, len(a), cap(a), &a)
	fmt.Printf("a1:%v len:%d, cap:%d, id:%p\n", a1, len(a1), cap(a1), &a1)
	fmt.Printf("a2:%v len:%d, cap:%d, id:%p\n", a2, len(a2), cap(a2), &a2)
	a[3] = 4
	a = append(a, 1, 2, 3, 5)
	a[2] = 9
	fmt.Printf("a:%v len:%d, cap:%d, id:%p\n", a, len(a), cap(a), &a)
	fmt.Printf("a1:%v len:%d, cap:%d, id:%p\n", a1, len(a1), cap(a1), &a1)
	fmt.Printf("a2:%v len:%d, cap:%d, id:%p\n", a2, len(a2), cap(a2), &a2)
	a1 = a[2:4]
	a2 = a[3:6]
	fmt.Printf("new a1:%v len:%d, cap:%d, id:%p\n", a1, len(a1), cap(a1), &a1)
	fmt.Printf("new a2:%v len:%d, cap:%d, id:%p\n", a2, len(a2), cap(a2), &a2)
	// fmt.Println()
}
