package __concurrencyControl

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
sync.WaitGroup
sync.WaitGroup 内部维护着一个计数器，计数器的值可以增加和减少。

(wg *WaitGroup) Add(delta int)	计数器+delta
(wg *WaitGroup) Done()	计数器-1
(wg *WaitGroup) Wait()	阻塞直到计数器变为 0

*/
func TestS1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("hello")
		time.Sleep(time.Second)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main goroutine done!")
}

/*
 sync.Once
sync.Once 其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，
而布尔值用来记录初始化是否完成。
这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。

*/

var icons map[string]string
var loadIconsOnce sync.Once

func TestS2(t *testing.T) {
	for i := 0; i < 20; i++ {
		Icon("left")
	}
}

// Icon 是并发安全的
func Icon(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func loadIcons() {
	fmt.Println("load icons ...")
	icons = map[string]string{
		"left":  "left.png",
		"up":    "up.png",
		"right": "right.png",
		"down":  "down.png",
	}
}

// 扩展内容：
// 互斥锁，读写锁。
// sync.Cond 用法
// sync.Map 用法
// sync.Pool 用法和作用
// 原子操作: atomic.Add、Swap、CompareAndSwap、Load、Store。
// 定时器:  Timer,Ticker
// context.Context 用法以及扩展用法。
// context.cancelCtx、timerCtx、valueCtx
// runtime 的协程调度：runtime.Gosched()，runtime.Goexit()，runtime.GOMAXPROCS
