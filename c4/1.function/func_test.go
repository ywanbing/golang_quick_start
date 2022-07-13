package function

import (
	"fmt"
	"os"
	"testing"
)

/*
GO 函数的特点：
   • 无需声明原型。
   • 支持不定 变参。
   • 支持多返回值。
   • 支持命名返回参数。
   • 支持匿名函数和闭包。
   • 函数也是一种类型，一个函数可以赋值给变量。

   • 不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
   • 不支持 重载 (overload)
   • 不支持 默认参数 (default parameter)。
*/

// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func F1(a, b, c int, d string, e *bool) {}

// 多返回值用括号, 可以命名返回参数
func F2(a int) (b bool, e error) {
	return true, nil
}

func F3(a int) (bool, error) {
	return true, nil
}

// 不定参数
func F4(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

// 命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
func F5(x, y int) (z int) {
	z = x + y
	return
}

// 匿名函数
func F6() {
	sum := func(nums ...int) (sum int) {
		for _, num := range nums {
			sum += num
		}
		return
	}
	sum(1, 2, 3, 4, 5)

	// 函数可以做为结构字段
	d := struct {
		fn func()
	}{
		fn: func() {
			fmt.Println("hello")
		},
	}
	d.fn()
}

// 闭包
func Test_F7(t *testing.T) {
	c := a()
	c()
	c()
	c()

	c = a()
	c()
	c()
}

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

// defer
/*
defer特性：
   1. 关键字 defer 用于注册延迟调用。
   2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
   3. 多个defer语句，按先进后出的方式执行。
   4. defer语句中的变量，在defer声明时就决定了。

defer用途：
    1. 关闭文件句柄
    2. 锁资源释放
    3. 数据库连接释放
*/
func TestF8(t *testing.T) {
	var whatever [5]struct{}

	// defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。
	for i := range whatever {
		defer fmt.Println(i)
	}

	for i := range whatever {
		// 类似闭包，最后执行在去取值，这个时候全部都是最后一个值。
		defer func() {
			fmt.Println(i)
		}()
	}
}

func TestF9(t *testing.T) {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y) // y 闭包引用
	}(x) // x 被复制

	x += 10
	y += 100
	println("x =", x, "y =", y)
}


// defer 陷阱练习
func TestF10(t *testing.T) {
	f, err := os.Open("1.txt")
	if err != nil {
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close e1.txt err %v\n", err)
		}
	}()

	// ..code...

	f, err = os.Open("2.txt")
	if err != nil {
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close 2.txt err %v\n", err)
		}
	}()
}