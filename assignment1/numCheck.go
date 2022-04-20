package main

import "fmt"

type nums []int

func generateSlices(num int) nums{
	nums := []int{}
	for i := 0 ; i <= num ; i++ {
		nums = append(nums, i)
	} 
	return nums
}

func (n nums) checkEvenOdd() {
	for i := range n {
		if n[i] % 2 == 0 {
			fmt.Println(n[i], "is even")
		}else{
			fmt.Println(n[i], "is odd")
		}
	}
}