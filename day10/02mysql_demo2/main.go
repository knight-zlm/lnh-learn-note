package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "crawlab:Crawlab@123@tcp(10.10.23.101:33306)/crawlab_articles"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn:%s invalid, error:%v\n", dsn, err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s fialed, error:%v\n", dsn, err.Error())
		return
	}
	return
}
func queryOne(id int) {
	sqlStr := "select id,name,age form user where id=?"
	var u user
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

func queryMore(n int) {
	sqlStr := "select id,name,age form user where id<?"
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("sql:%s faild,err:%s\n", sqlStr, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("row scan error,%s", err.Error())
		}
		fmt.Println(u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("db init error:%v\n", err.Error())
	}
	fmt.Println("数据库链接成功！")
	queryOne(2)
}
