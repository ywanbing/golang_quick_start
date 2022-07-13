package arr

import (
	"fmt"
	"testing"
)

/* 数组的特性
   1. 数组：是同一种数据类型的固定长度的序列。
   2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
   3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
   4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
   for i := 0; i < len(a); i++ {
   }
   for index, v := range a {
   }
   5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
   6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
   7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
   8.指针数组 [n]*T，数组指针 *[n]T。
   9.[0]T也是一个数组，但是它占用的内存为0
*/


// 全局：
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func Test_A1(t *testing.T) {
	// 局部：
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}

	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
}



// 多维数组
var arr10 [5][3]int
var arr11 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func Test_A2(t *testing.T) {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。 把它看作 [...]T 这个T必须是确定的类型。所以 第2维 不能被推导

	fmt.Println(arr10, arr11)
	fmt.Println(a, b)
}


// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。
func test(x [2]int) {
	fmt.Printf("test x: %p\n", &x)
	x[1] = 1000
}

func test1(x *[2]int) {
	fmt.Printf("test1 x: %p\n", x)
	x[1] = 1000
}

func Test_A3(t *testing.T) {
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)

	test(a)
	fmt.Println(a)

	test1(&a)
	fmt.Println(a)
}


// 内置函数 len 和 cap 都返回数组长度 (元素数量)。
func Test_A4(t *testing.T) {
	a := [2]int{}
	println(len(a), cap(a))
}


// 多维数组遍历
func Test_A5(t *testing.T) {
	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

	for k1, v1 := range f { 	// v1 : [3]int{}
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}
}

// 数组拷贝和传参
func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func Test_A6(t *testing.T) {
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)

	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}


// 练习：
// 找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）

// 求元素和，是给定的值
func myTest(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历，只寻找后面的数
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func Test_A7(t *testing.T) {
	b := [5]int{1, 3, 5, 8, 7}
	myTest(b, 8)
}
