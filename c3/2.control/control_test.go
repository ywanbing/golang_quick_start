package control

import (
	"fmt"
	"testing"
)

/*
条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true
来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。
*/

// if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
// 	• 可省略条件表达式括号。
//    	• 支持初始化语句，可定义代码块局部变量。
//    	• 代码块 左括号 必须在条件表达式尾部。
//
// if 布尔表达式 {
//   /* 在布尔表达式为 true 时执行 */
// }
func TestC1(t *testing.T) {
	x := 0

	if n := "abc"; x > 0 {     // 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
		println(n[2])
	} else if x < 0 {    // 注意 else if 和 else 左大括号位置。
		println(n[1])
	} else {
		println(n[0])
	}
}

/*
switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，
从上直下逐一测试，直到匹配为止。 Golang switch 分支表达式可以是任意类型，
不限于常量。可省略 break，默认自动终止。

switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
switch x.(type){
	case type:
	   	statement(s)
	case type:
	   	statement(s)
	// 你可以定义任意个数的case
	default: // 可选
		statement(s)
}
*/
func TestC2(t *testing.T) {
	var x interface{}
	// 写法一：
	switch i := x.(type) { // 带初始化语句
	case nil:
		fmt.Printf(" x 的类型 :%T \r\n", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	// 写法二
	var j = 0
	switch j {
	case 0:
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}

	// 写法三
	var k = 0
	switch k {
	case 0:
		println("fallthrough")
		fallthrough
		/*
		   Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
		   而如果switch没有表达式，它会匹配true。
		   Go里面switch默认相当于每个case最后带有break，
		   匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		   但是可以使用fallthrough强制执行后面的case代码。
		*/
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}

	// 写法三
	var m = 0
	switch m {
	case 0, 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}

	// 写法四
	var n = 0
	switch { //省略条件表达式，可当 if...else if...else
	case n > 0 && n < 10:
		fmt.Println("i > 0 and i < 10")
	case n > 10 && n < 20:
		fmt.Println("i > 10 and i < 20")
	default:
		fmt.Println("def")
	}
}

/*
循环语句 for

Golang for支持三种循环方式，包括类似 while 的语法。

	for init; condition; post { }
    for condition { }
    for { }

    init： 一般为赋值表达式，给控制变量赋初值；
    condition： 关系表达式或逻辑表达式，循环控制条件；
    post： 一般为赋值表达式，给控制变量增量或减量。

    for语句执行过程如下：
    ① 先对表达式 init 赋初值；
    ② 判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，
		则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；
		否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。
*/
func TestC3(t *testing.T) {
	/* for 循环 */
	for i := 0; i < 10; i++ {
		fmt.Printf("a 的值为: %d\n", i)
	}

	var b int = 15
	var a int
	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	n := 5
	for  {
		if n < 1 {
			break
		}
		fmt.Println(n)
		n--
	}
}

/*
Golang range类似迭代器操作，返回 (索引, 值) 或 (键, 值)。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。

for key, value := range oldMap {
    newMap[key] = value
}


			1st value		2nd value
string		index			s[index]		unicode, rune
array/slice	index			s[index]
map			key				m[key]
channel		element

可忽略不想要的返回值，或 "_" 这个特殊变量。

*/
func TestC4(t *testing.T) {
	s := "abc"
	// 忽略 2nd value，支持 string/array/slice/map。
	for i := range s {
		println(s[i])
	}

	// 忽略 index。
	for _, c := range s {
		println(c)
	}

	// 忽略全部返回值，仅迭代。
	for range s {

	}

	m := map[string]int{"a": 1, "b": 2}
	// 返回 (key, value)。
	for k, v := range m {
		println(k, v)
	}

	for k := range m {
		println(k)
	}
}

/*
循环控制语句

循环控制语句可以控制循环体内语句的执行过程。

Goto、Break、Continue:
	1.三个语句都可以配合标签(label)使用
    2.标签名区分大小写，定义以后若不使用会造成编译错误
    3.continue、break配合标签(label)可用于多层循环跳出
    4.goto是调整执行位置，与continue、break配合标签(label)的结果并不相同
*/
func TestC5(t *testing.T) {
LABEL1:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 1 {
				continue LABEL1   // = break
				// break LABEL1   // 退出外层循环
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	i:=0
HERE:
	println(i)
	i++
	if i==5 {
		return
	}
	goto HERE
}

//