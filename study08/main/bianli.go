package main

import "fmt"

func main() {
	s1 := []int{1,2,3,4,5,6}
	// for循环遍历
	for i := 0 ; i < len(s1) ; i++ {
		fmt.Printf("s1[%d]: %v\n", i,s1[i])
		/* 
			s1[0]: 1
			s1[1]: 2
			s1[2]: 3
			s1[3]: 4
			s1[4]: 5
			s1[5]: 6
		 */
	}
	fmt.Println("-----------------")
	// range 遍历
	for i, v := range s1 {
		fmt.Printf("s1[%d]:%v\n", i,v)
	}
	/*
		s1[0]:1
		s1[1]:2
		s1[2]:3
		s1[3]:4
		s1[4]:5
		s1[5]:6 
	 */
}