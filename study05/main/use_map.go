package main

import "fmt"
func main (){
	// 声明变量  map[key的类型]值的类型
	var cityMap map[string]string
	// 创建集合
	cityMap = make(map[string]string)

	// map插入键值对
	cityMap["Shanghai"] = "Shanghai"
	cityMap["Zhejiang"] = "hangzhou"
	cityMap["jiangxi"] = "nanchang"

	// map 遍历
	for city := range cityMap {
		fmt.Println("city of ",city,"is",cityMap[city])
	}
	/*
		city of  Zhejiang is hangzhou
		city of  jiangxi is nanchang
		city of  Shanghai is Shanghai
	*/

	// 查看元素在集合中是否存在
	city, ok := cityMap["Zhejiang"]
	if ok {
		fmt.Println("浙江的省会城市是：",city)
	}else{
		fmt.Println("没有查询到浙江的省会城市")
	}

	/*
		浙江的省会城市是： hangzhou
	*/

	// 删除map数据
	delete(cityMap,"jiangxi")
	fmt.Println("删除后的map：")
	// map 遍历
	for city := range cityMap {
		fmt.Println("city of ",city,"is",cityMap[city])
	}

	/*
		删除后的map：
		city of  Shanghai is Shanghai
		city of  Zhejiang is hangzhou
*/

}