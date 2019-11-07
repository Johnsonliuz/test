//進階練習
//學生成績大於等於60分輸出合格，低於60輸出不合格，成績不能高於100分

package main

import "fmt"

func main() {
	var score int
	
	fmt.Println("請輸入成績：")
	fmt.Scanf("%d", &score)
	
	switch int(score/60){
		case 1 :
			fmt.Println("合格")
		case 0 :
			fmt.Println("不合格")
	}
}