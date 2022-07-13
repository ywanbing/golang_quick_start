package __concurrencyControl

import (
	"fmt"
	"testing"
	"time"
)

/*
	select多路复用

select的使用类似于switch语句，它有一系列case分支和一个默认的分支。
每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，
直到某个case的通信操作完成时，就会执行case分支对应的语句。

select {
    case <-chan1:
       // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
       // 如果成功向chan2写入数据，则进行该case处理语句
    default:
       // 如果上面都没有成功，则进入default处理流程
}

执行步骤：
1. 所有channel表达式都会被求值、所有被发送的表达式都会被求值。求值顺序：自上而下、从左到右.
结果是选择一个发送或接收的channel，无论选择哪一个case进行操作，表达式都会被执行。
RecvStmt 左侧短变量声明或赋值未被评估。
2. 如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行，
否则的话，如果有default分支，则执行default分支语句，
如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行.
3. 除非所选择的情况是默认情况，否则执行相应的通信操作。
4. 如果所选case是具有短变量声明或赋值的RecvStmt，则评估左侧表达式并分配接收值（或多个值）。
5. 执行所选case中的语句
*/

// select可以同时监听一个或多个channel，直到其中一个channel ready
func TestC1(t *testing.T) {
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go test1(output1)
	go test2(output2)
	// 用select监控
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}

func test1(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

// 如果多个channel同时ready，则随机选择一个执行
func TestC2(t *testing.T) {
	output1 := make(chan string)
	output2 := make(chan string)
	go func() {
		output1 <- "hello1"
	}()
	go func() {
		output2 <- "hello2"
	}()

	select {
	case value := <-output1:
		fmt.Println("string1:", value)
	case value := <-output2:
		fmt.Println("string2:", value)
	}
	fmt.Println("main结束")
}


// 所有channel表达式都会被求值、所有被发送的表达式都会被求值。求值顺序：自上而下、从左到右.
var ch1 chan int
var ch2 chan int
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func TestC3(t *testing.T) {
	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default!.")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)

	return numbers[i]
}
func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)

	return chs[i]
}