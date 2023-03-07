package main

import "fmt"

func main() {
	var scores [5] int
	for i := 0;i<len(scores);i++ {
		fmt.Printf("请输入第%d个学生的成绩",i + 1)
		fmt.Scanln(&scores[i])
	}
	// 方式一：普通for循环
	for i :=0; i<len(scores);i++ {
		fmt.Printf("第%d个学生的成绩为: %d\n",i + 1, scores[i])
	}
	fmt.Println("------------------------------------")
	// 方式二：for-range循环
	for key,value := range scores {
		fmt.Printf("第%d个学生的成绩为: %d\n",key + 1, value)
	}
	// 如果不想输出key
	for _,value := range scores {
		fmt.Printf("学生的成绩为: %d\n", value)
	}
}