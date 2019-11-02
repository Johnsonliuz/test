//練習
//有三個班，每班有五個學生，求每個班的平均分數及總平均分數
//學生成績用手動輸入

package main

import "fmt"

func main() {
	
	var class int
	var score int
	var sum int
	var allsum int //三個班平均後班總和
	
	for class = 1; class <= 3; class++{
	sum = 0
		for i := 1; i <= 5; i++{
		fmt.Printf("請輸入第%d班第%d個學生的成績：",class,i)
		fmt.Scanf("%d", &score)
		sum += score
	}
	
	fmt.Printf("第%v個班的平均分數=%v\n",class,sum/5)
	allsum += sum/5
	}
	
	fmt.Printf("總平均=%v\n",allsum/3)
}