package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	ci, err := strconv.Atoi("asc")
	if err != nil {
		errors.Wrap(err, "additional messageto give error")
		fmt.Println(err.Error())
	}
	fmt.Println(ci)
	if s == nil {
		fmt.Println("nil")
	}
}
