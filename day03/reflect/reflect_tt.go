package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name string
	Id   int
	Age  int
}

type Mananger struct {
	User
	Title string
}

func (u User) Hello(names []string) {
	var totileName string
	totileName = strings.Join(names, ",")
	fmt.Printf("Hello %s, my name is %s", totileName, u.Name)
}

func main() {
	u := User{"ok", 1, 13}
	m := Mananger{User: User{"ok", 1, 13}, Title: "cto"}
	t := reflect.TypeOf(m)
	fmt.Printf("%v", t.Field(0))
	// usr := User{"ok", 1, 13}
	Set(&m)
	fmt.Printf("mananger %v\n", m)
	// 反射回调函数
	elem := reflect.ValueOf(&u).Elem()
	reFunc := elem.MethodByName("Hello")
	// reflect.
	// inputValue := []reflect.Value{reflect.ValueOf("world"), reflect.ValueOf("world2")}
	reFunc.Call([]reflect.Value{reflect.ValueOf([]string{"world", "world2"})})
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	// f := v.FieldByName("Title")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("zlm")
	}
}
