package main

import "fmt"

func main() {
	
	var a int
	var positive int
	var negative int
	
	//打印1-100 用for+continue實現
	for i := 1; i <= 100; i++{
		if i % 2 == 0{
			fmt.Println(i)
		} else {
			continue
		}
	}
	
	
	//從鍵盤輸入幾個不確定的整數，判斷他是正數，負數，遇到O則跳出結束(運用for+continue+break)
	for {
	
	fmt.Println("請輸入數字：")
	fmt.Scanln(&a)
	
		if a > 0 {
			a = positive
			positive++
			continue
		} else if a < 0 {
			a = negative
			negative++
			continue
		} else if a == 0 {
			break
		}
	}
	fmt.Println("共有%v個正數",positive)
	fmt.Println("共有%v個負數",negative)
}