//練習
//編寫一個程序，可以自己輸入年齡
//如果大於18歲則輸出"你的年齡大於18 要對自己的行為負責"
//不到18歲則輸出"多練幾年再來"

package main

import "fmt"

func main() {
	var age int

	fmt.Println("請輸入年齡:")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你的年齡大於18 要對自己的行為負責")
	} else {
		fmt.Println("多練幾年再來")
	}
}
//事實上