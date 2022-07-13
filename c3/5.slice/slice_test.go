package slice

import (
	"fmt"
	"testing"
)

/*
	1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
	2. 切片的长度可以改变，因此，可以认为切片是一个可变的数组。
	3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
	4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
	5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
	6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/

func Test_S1(t *testing.T) {
	// 1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}

	// 2.:=
	s2 := []int{}

	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int

	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)
}

// 通过数组来初始化切片

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]        // 可以简写为 var slice []int = arr[:end]
var slice2 []int = arr[5:10]       // 可以简写为 var slice[]int = arr[start:]
var slice3 []int = arr[0:len(arr)] // var slice []int = arr[:]
var slice4 = arr[:len(arr)-1]      // 去掉切片的最后一个元素
func Test_S2(t *testing.T) {
	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr2[2:8]
	slice6 := arr2[0:6]          // 可以简写为 slice := arr2[:end]
	slice7 := arr2[5:10]         // 可以简写为 slice := arr2[start:]
	slice8 := arr2[0:len(arr2)]  // slice := arr2[:]
	slice9 := arr2[:len(arr2)-1] // 去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)
}

// 通过make来创建切片
func Test_S3(t *testing.T) {
	slice3 := make([]int, 10)
	slice4 := make([]int, 10)
	slice5 := make([]int, 0, 20)
	fmt.Printf("make局部slice3 ：%v\n", slice3)
	fmt.Printf("make局部slice4 ：%v\n", slice4)
	fmt.Printf("make局部slice5 ：%v\n", slice5)
}

// 读写操作实际目标是底层数组，只需注意索引号的差别。
func Test_S4(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5}

	s := data[2:4]
	s[0] += 100
	s[1] += 200

	fmt.Println(s)
	fmt.Println(data)
}

// 可直接创建 slice 对象，自动分配底层数组。
func Test_S5(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))
}

// 切片引用的是数组指针
func Test_S6(t *testing.T) {
	// 数组 b
	b := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 切片 v
	v := b[:]
	fmt.Printf("%p \n", &b)
	fmt.Printf("%p \n", &v)
	fmt.Printf("%p \n", &v[0])

	// 数组的复制，对原来的数组没有影响
	c := b
	c[0] = 12

	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("--------------------------------")

	v[0] = 10

	// 切片的改变会影响数组的值
	fmt.Println(b)
	fmt.Println(v)
}

// 用 append 内置函数操作切片（切片追加）
func Test_S7(t *testing.T) {
	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)

	var b = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)

	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)

	d := append(c, 7)
	fmt.Printf("slice d : %v\n", d)

	e := append(d, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)
}

// 向 slice 尾部添加数据，返回新的 slice 对象。
func Test_S8(t *testing.T) {
	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1)

	// 在不超过 s1 cap 的情况下，底层数组其实为同一个
	s2 := append(s1, 1)
	fmt.Printf("%p\n", &s2)

	fmt.Println(s1, s2)
}

// 超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
func Test_S9(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[2:4:11] // len = 2(2-0), cap = 3(3-0)  4-2 = 2

	fmt.Println("-------",cap(s))


	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。

	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。

	data1 := [...]int{0, 1, 2, 3, 4, 10: 0}
	s1 := data1[:2:3]    // len = 2(2-0), cap = 3(3-0)
	s1 = append(s1, 100) // 不超过底层数组的容量

	fmt.Println(s1, data1)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s1[0], &data1[0]) // 比对底层数组起始指针。
}

// 切片拷贝
func Test_S10(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)

	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)

	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)

	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = s3[0:0]
	fmt.Printf("slice s3 : %v ,len = %v\n", s3,cap(s3))

	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)

	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)
}

// 字符串和切片（string and slice）
func Test_S11(t *testing.T) {
	str := "Hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)

	// string本身是不可变的，因此要改变string中字符。需要如下操作：
	s := []byte(str) // 中文字符需要用[]rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)

	// 含有中文字符串
	str2 := "你好，世界！hello world！"
	s3 := []rune(str2)
	s3[3] = '够'
	s3[4] = '浪'
	s3[12] = 'g'
	s3 = s3[:4]
	str2 = string(s3)
	fmt.Println(str2)
}

// 写一个函数在原地完成消除[]string中相邻重复(只存在一次重复)的字符串的操作
func Test_S12(t *testing.T) {
	strs := []string{"abc","abd","abe","abe","abf","abg","abg"}
	l := clear(strs)
	fmt.Println(strs[:l])
}

// 返回字符串数组的数量
func clear(strs []string) int {
	l := len(strs)
	for i := 0; i < len(strs); i++ {
		if i + 1 == len(strs) {
			break
		}
		// 减去下一个字符串
		if strs[i] == strs[i+1] {
			copy(strs[i+1:],strs[i+2:])
			l--
		}
	}
	return l + 1
}