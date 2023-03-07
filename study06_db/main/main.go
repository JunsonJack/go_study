package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局变量db
var db *sql.DB
// 初始化数据库连接
func initDB()(err error){
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err!=nil {
		return err
	}
	// 尝试与数据库建立连接，校验dsn是否正确
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
// 插入数据
func insertData (){
	sqlStr := "insert into user_tb(username,psw) values(?,?)"
	ret,err := db.Exec(sqlStr,"wangwu","123admin")
	if err != nil {
		fmt.Printf("插入失败", err)
		return 
	}
	theID,err := ret.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Printf("获取插入数据id失败", err)
		return
	}
	fmt.Printf("插入成功，插入的数据id为%v", theID)

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("初始化失败!,err:%v\n",err)
		return
	}else{
		fmt.Printf("连接成功")
	}

	insertData() 
	//连接成功插入成功，插入的数据id为3
}