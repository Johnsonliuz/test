//練習
//有三個班，每班有五個學生，求每個班的平均分數及總平均分數
//學生成績用手動輸入
//統計三個班及格人數

package main

import "fmt"

func main() {
	
	var classnum int = 3
	var stunum int = 5
	var score int
	var sum float64
	var allsum float64 //三個班平均後班總和
	var passcount int
	
	for j := 1; j <= classnum; j++{
	sum = 0
		for i:= 1; i <= stunum; i++{
		fmt.Printf("請輸入第%d班第%d個學生的成績：",j,i)
		fmt.Scanf("%d", &score)
		sum += float64(score)
		if score >= 60{
		passcount++
		}
	}
	
	fmt.Printf("第%v個班的平均分數=%v\n",j,sum/5)
	allsum += float64(sum/5)
	}
	
	fmt.Printf("總平均=%v\n",allsum/3)
	fmt.Printf("及格人數=%v\n",passcount)
}