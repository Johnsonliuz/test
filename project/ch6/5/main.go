//有一堆桃子，猴子每天吃一半再多一顆，吃完九天後剩1顆
//請問原本有多少顆？

/*思路分析
第十天剩一顆，所以回傳一顆
第九天會是第十天+1在*2
可以列出(n+1)*2
再放入f()
*/

package main

import "fmt"

func f(n int)(a int){
	if n == 10{
		a = 1
	} else {
		a = (f(n+1)+1)*2
	}
	return a
}

func main() {
	var a int
	a = f(1)
	
	fmt.Println(a)
}