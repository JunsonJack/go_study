package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type Product struct {
	gorm.Model
	Code string
	Price uint
}
var db *gorm.DB
// init数据连接
func initDB (){
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil {
		panic("failed to connect database")
	}
	db = d
}

// 创建数据表
func createTable(){
	// 直接把上面的结构体创建表到数据库
	db.AutoMigrate(&Product{})

	/*
		CREATE TABLE `products` (`id` bigint unsigned AUTO_INCREMENT,`created_at` 
		datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`code` longtext,`price` bigint unsigned,PRIMARY KEY (`id`),INDEX `idx_products_deleted_at` (`deleted_at`))
	*/
}
// 插入数据
func insert (){
	// 插入数据
	p := Product{
		Code:"1001",
		Price: 100,
	}
	db.Create(&p)
}
// 查询数据
func find (){
	var product Product
	db.First(&product,1) //根据整形主键进行查询
	fmt.Printf("product.Code1: %v\n", product.Code)
	db.Take(&product)
	fmt.Printf("product.Code2: %v\n", product.Code)
	db.Last(&product)
	fmt.Printf("product.Code3: %v\n", product.Code)
}
// 更新数据
func updateData(){
	var product Product
	db.Model(&product).Update("Price",200) //更新单个字段
	//更新多个字段
	db.Model(&product).Updates(Product{Price: 200,Code: "F421"}) //仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200,"Code":"xxxyyy"})
	fmt.Printf("product.Code3: %v\n", product.Code)

}
// 删除数据
func delete_db (){
	var p Product
	db.Delete(&p)
}
func main() {
	initDB()
	find()
	/*
		product.Code1: 1001
		product.Code2: 1001
		product.Code3: 1001
	*/
	// updateData()
	delete_db()

}