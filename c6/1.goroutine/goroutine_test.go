package __goroutine

import (
	"fmt"
	"testing"
)

/*
Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。
*/

func hello() {
	fmt.Println("Hello Goroutine!")
}

func TestG1(t *testing.T) {
	hello()
	fmt.Println("main goroutine done!")
}


func TestG2(t *testing.T) {
	go hello()
	go hello()
	go hello()
	fmt.Println("main goroutine done!")
	// time.Sleep(time.Millisecond)
}