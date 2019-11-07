//練習
//根據用戶指定的月份，打印該月份所屬的季節
//3.4.5 春
//6.7.8 夏
//9.10.11 秋
//12.1.2 冬

package main

import "fmt"

func main() {
	var month int
	
	fmt.Println("請輸入月份：")
	fmt.Scanf("%d", &month)
	
	switch month {
		case 3,4,5 :
		fmt.Println("春")
		case 6,7,8 :
		fmt.Println("夏")
		case 9,10,11 :
		fmt.Println("秋")
		case 12,1,2 :
		fmt.Println("冬")
		default:
		fmt.Println("輸入錯誤")
	}
}