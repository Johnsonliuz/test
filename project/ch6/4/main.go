//已知條件：f(1)=3;f(n)=2*f(n-1)+1;

package main

import "fmt"

func f(n int)(a int){
	if n == 1{
		a = 3
	} else {
		a = 2 * f(n - 1 ) + 1
	}
	return a
}

func main() {
	var a int
	a = f(10)
	
	fmt.Println(a)
}