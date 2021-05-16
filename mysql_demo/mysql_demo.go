package mysql_demo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func MysqlDemo() {
	dataSourceName := "root:mysafziy@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Open mysql connection error {%v}", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Close mysql connection error {%v}", err)
		}
	}(db)
	err = db.Ping()
	if err != nil {
		log.Fatalf("Open mysql connection error {%v}", err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10) // 设置与数据库建立连接的最大数目
	db.SetMaxIdleConns(10) // 设置连接池中的最大闲置连接数

	var id int
	var name string
	var age int
	querySql := "SELECT * FROM `user` WHERE id = 1"
	err = db.QueryRow(querySql).Scan(&id, &name, &age)
	if err != nil {
		log.Fatalf("Mysql query row error {%v}", err)
	}
	fmt.Println(id, name, age)

}
