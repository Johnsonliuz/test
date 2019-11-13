//隨機產生兩個數 當和大於100則跳出

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	var a int
	var b int
	
	for{
	rand.Seed(time.Now().UnixNano())
	a = rand.Intn(100) + 1
	b = rand.Intn(100) + 1
	
	fmt.Printf("a=%v\n",a)
	fmt.Printf("b=%v\n",b)
	fmt.Printf("總和=%d\n",a+b)
	if a + b > 100{
		break
		}
	}
}