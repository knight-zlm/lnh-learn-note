package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "crawlab:Crawlab@123@tcp(10.10.23.101:33306)/crawlab_articles"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn:%s invalid, error:%v\n", dsn, err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s fialed, error:%v\n", dsn, err.Error())
	}
	fmt.Println("数据库链接成功！")
}
