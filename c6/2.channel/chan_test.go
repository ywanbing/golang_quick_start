package __channel

import (
	"fmt"
	"testing"
	"time"
)

/*
Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，
总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

channel是一种类型，一种引用类型。声明通道类型的格式如下：
	var 变量 chan 元素类型
*/

var c1 chan int   // 声明一个传递整型的通道
var c2 chan bool  // 声明一个传递布尔型的通道
var c3 chan []int // 声明一个传递int切片的通道

/*
通道是引用类型，通道类型的空值是nil。
声明的通道后需要使用 make 函数初始化之后才能使用。

	make(chan 元素类型, [缓冲大小])
*/
func TestC1(t *testing.T) {
	ch4 := make(chan int)
	ch5 := make(chan bool, 10)
	ch6 := make(chan []int)

	_ = ch4
	_ = ch5
	_ = ch6
}

/*
通道有发送（send）、接收(receive）和关闭（close）三种操作。

发送和接收都使用 <- 符号。
*/
func TestC2(t *testing.T) {
	ch := make(chan int)
	go func() {
		fmt.Println("send ",1)
		ch <- 1
	}()

	a := <- ch
	fmt.Println("receive ",a)

	close(ch)
}

/*
关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。
通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

	1.对一个关闭的通道再发送值就会导致panic。
    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
    4.关闭一个已经关闭的通道会导致panic。

*/

// 无缓冲的通道又称为阻塞的通道。
func TestC3(t *testing.T) {
	ch := make(chan int)
	ch <- 10 	// 写入通道，一直等待接受者接收
	// <- ch	// 获取通道的值，一直等待发送者发送
	fmt.Println("发送成功")
}

// 只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
// 当通道的缓冲区被放满了，又会被阻塞，直到有接受者拿走其中的值。
func TestC4(t *testing.T) {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道

	ch <- 10
	fmt.Println("发送成功 10")

	go func() {
		time.Sleep(2 * time.Second)
		<- ch
	}()

	ch <- 11
	fmt.Println("发送成功 11")
}


// 判断通道是否已经关闭的操作
func TestC5(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}


/*
单向通道
有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，
比如限制通道在函数中只能发送或只能接收。
	1. chan<- int 是一个只能发送的通道，可以发送但是不能接收；
	2. <-chan int 是一个只能接收的通道，可以接收但是不能发送。
在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来是不可以的。
*/
func TestC6(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)
	go squarer(ch2, ch1)

	printer(ch2)
}

func counter(in chan<- int) {
	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(out <-chan int) {
	for i := range out {
		fmt.Println(i)
	}
}