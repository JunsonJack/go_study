package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局变量db
var db *sql.DB
// 初始化数据库连接
func InitDB()(err error){
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
// 定义一个结构体
type User struct {
	id int
	username string
	psw string
}
//查询单条数据
func queryRowData (){
	sqlStr := "select id ,username,psw from user_tb where id = ?"
	var u User
	err := db.QueryRow(sqlStr,2).Scan(&u.id, &u.username, &u.psw)
	if err != nil {
		fmt.Println("scan failed , err=" , err)
		return
	}
	fmt.Println("id:",u.id,"name:",u.username,"psw:",u.psw)
}

// 查询多条数据
func queryMultiRowData(){
	sql := "select id ,username,psw from user_tb where id > ?"
	rows,err := db.Query(sql,0)
	if err != nil {
		fmt.Println("查询失败",err)
		return
	}
	// 非常重要，关闭rows释放持有的数据库连接
	defer rows.Close()
	
	// 循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.username, &u.psw)
		if err != nil {
			fmt.Println("查询失败",err)
			return
		}
		fmt.Println("id:",u.id,"name:",u.username,"psw:",u.psw)
	}
}

// 删除数据
func deleteRowData (){
	sqlStr := "delete from user_tb where id = ?"
	ret,err := db.Exec(sqlStr,3)
	if err != nil {
		fmt.Println("删除失败",err)
		return 
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Println("获取影响行数失败",err)
		return
	}
	fmt.Println("删除成功，影响行数为：",n)
}

// 修改数据
func updateRowData (){
	sqlStr := "update user_tb set username=? where id =?"
	ret,err := db.Exec(sqlStr,39,2)
	if err!= nil {
		fmt.Println("更新失败",err)
		return
	}
	n,err := ret.RowsAffected()
	if err!=nil {
		fmt.Println("获取影响行数失败",err)
		return
	}
	fmt.Println("获取影响行数成功，行数为：",n)
}
func main() {
	err := InitDB()
	if err != nil {
		fmt.Printf("初始化失败!,err:%v\n",err)
		return
	}else{
		fmt.Printf("连接成功")
	}
	
	queryRowData()
	// 连接成功id: 2 name: wangwu psw: 123admin

	queryMultiRowData()
	/*
		id: 1 name: zhangsan psw: 122344
		id: 2 name: wangwu psw: 123admin
		id: 3 name: wangwu psw: 123admin
	*/
	deleteRowData()
	// 删除成功，影响行数为： 1
	updateRowData()
	//获取影响行数成功，行数为： 1
}