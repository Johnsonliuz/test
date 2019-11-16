package main

import "fmt"

fnuc main() {
	var a int
	var b int
	
	fmt.Println("請輸入第一個整數:")
	fmt.Scanln(&a)
	fmt.Println("請輸入第二個整數:")
	fmt.Scanln(&b)
	
	res1,res2 := utils.Getsumandsub(1,2)
	fmt.Printf("res1=%v,res2=%v",res1,res2)
	
}