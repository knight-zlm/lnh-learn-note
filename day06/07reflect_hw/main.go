package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// MysqlConfig MysqlConfig
type MysqlConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

func loadInit(v interface{}) {
	fObj, err := os.OpenFile("./my.ini", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("open file error:%s", err.Error())
	}
	defer fObj.Close()
	san := bufio.NewScanner(fObj)
	initData := make(map[string]interface{})
	for san.Scan() {
		// 判断
		line := san.Text()
		if line[0] == '[' {
			continue
		}
		spLine := strings.Split(line, "=")
		key := spLine[0]
		value := spLine[1]
		initData[key] = value
	}
	reft := reflect.TypeOf(v)
	refv := reflect.ValueOf(v)
	for i := 0; i < reft.NumField(); i++ {
		field := reft.Field(i)
		value, ok := initData[field.Tag.Get("init")]
		if !ok {
			continue
		}
		fmt.Println(value)
		// 需要知道类型 反射设置值
		// field.Kind()
		stValue := refv.FieldByName(field.Name)
		fmt.Println(stValue)
	}
}

func main() {
	var init MysqlConfig
	init = MysqlConfig{}
	reft := reflect.TypeOf(init)
	fmt.Println(reft.NumField())
	// for i := 0; i < reft.NumField(); i++ {
	// 	field := reft.Field(i)
	// 	fmt.Println(field.Tag.Get("host"))
	// 	// 需要知道类型 反射设置值
	// 	// field.Kind()
	// 	// stValue := refv.FieldByName(field.Name)
	// 	// fmt.Println(stValue)
	// }
	// loadInit(&init)
	fmt.Printf("Host:%s,Prot%d,UserName:%s,Password:%s", init.Host, init.Port, init.UserName, init.Password)
}
