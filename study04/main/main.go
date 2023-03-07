package main

import "fmt"
func main() {
	// 定义一个数组
	var scores [5] int
	scores[0] = 95
	scores[1] = 77
	scores[2] = 56
	scores[3] = 67
	scores[4] = 88

	// 求和
	sum := 0
	for i :=0;i< len(scores);i++ {
		sum += scores[i]
	}
	// 平均数
	avg := sum / len(scores)
	// 输出
	fmt.Printf("成绩总和为：%v,成绩的平均数为：%v",sum,avg)
	// 成绩总和为：383,成绩的平均数为：76

}