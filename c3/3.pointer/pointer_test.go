package pointer

import (
	"fmt"
	"testing"
)

/*

区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。

要搞明白Go语言中的指针需要先知道3个概念：指针地址、指针类型和指针取值。

1. Go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。
2. 传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。
3. Go语言中的指针操作非常简单，只需要记住两个符号：&（取地址）和*（根据地址取值）。

*/


/*
每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
Go语言中使用 &字符放在变量前面对变量进行“取地址”操作。
Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等。
*/

// 取变量指针 ptr := &v
//  v:代表被取地址的变量，类型为 T
//	ptr:用于接收地址的变量，ptr的类型就为 *T，称做 T的指针类型。*代表指针。
func TestP1(t *testing.T) {
	var a int = 10  // 定义一个 int 变量,并赋值
	b := &a
	fmt.Printf("a: %d ptr: %p \n", a, &a)
	fmt.Printf("b: %p type: %T \n", b, b)
	fmt.Println(&b)
}

// 在对普通变量使用 & 操作符取地址后会获得这个变量的指针，然后可以对指针使用 *操作，也就是指针取值
func TestP2(t *testing.T) {
	// 指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)

	c := *b // 指针取值（根据指针去内存取值）

	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

/*
 取地址操作符 & 和取值操作符 * 是一对互补操作符，& 取出地址，* 根据地址取出地址指向的值。
	1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
    2.指针变量的值是指针地址。
    3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
*/


// 指针在函数上的应用
func TestP3(t *testing.T) {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

// 空指针
func TestP4(t *testing.T) {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是 %v \n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

// 执行下面的代码会引发panic，为什么呢？
func TestP5(t *testing.T) {
	var a *int
	*a = 10
	fmt.Println(*a)
}

/*
在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
*/


/*
new是一个内置的函数，它的函数签名如下：
	func new(Type) *Type
		1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
    	2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
*/
func TestP6(t *testing.T) {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T \n", a) // *int
	fmt.Printf("%T \n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}

func TestP7(t *testing.T) {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)
}

// 指向指针的指针
func TestP8(t *testing.T) {
	var a int
	type T *int
	var ptr *int // T
	var pptr **int // *T

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	fmt.Printf("%T \n", ptr) //  *int
	fmt.Printf("%T \n", pptr) // **int

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

// 练习: 判断这个交换是否成功！
func TestP9(t *testing.T) {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(a, b *int) {
	b, a = a, b
}
