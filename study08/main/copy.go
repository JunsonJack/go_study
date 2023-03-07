package main

import "fmt"

func main() {
	s1 := []int{1,2,3}
	s2 := s1 //直接赋值
	s1[0] =100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Println("--------")

	s3 := make([]int,3)

	copy(s3,s1)

	s1[0] = 1

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

	/*
		s1: [100 2 3]
		s2: [100 2 3]
		--------
		s1: [1 2 3]
		s2: [1 2 3]  
	 */
}