package _interface

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

/*
	接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。

	interface是一组 method的集合，是 duck-type programming的一种体现。
	接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。
不关心属性（数据），只关心行为（方法）。
*/


/*
	接口是一个或多个方法签名的集合。
    任何类型的方法集中只要拥有该接口'对应的全部方法'签名。
    就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。
    这称为 Structural Typing。
    所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。
    当然，该类型还可以有其他方法。

    接口只有方法声明，没有实现，没有数据字段。
    接口可以匿名嵌入其他接口，或嵌入到结构中。
    对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
    只有当接口存储的类型和对象都为nil时，接口才等于nil。
    接口调用不会做receiver的自动转换。
    接口同样支持匿名字段方法。
    接口也可实现类似OOP中的多态。
    空接口可以作为任何类型数据的容器。
    一个类型可实现多个接口。
    接口命名习惯以 er 结尾。


    type 接口类型名 interface{
        方法名1( 参数列表1 ) 返回值列表1
        方法名2( 参数列表2 ) 返回值列表2
        …
    }

 	1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，
如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
    2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，
这个方法可以被接口所在的包（package）之外的代码访问。
    3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

	type writer interface{
    	Write([]byte) error
	}
*/

// 实现接口的条件
// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。

// Sayer 接口
type Sayer interface {
	say()
}

type dog struct {
	name string
}

type cat struct {}

// 因为 Sayer接口里只有一个 say方法，所以我们只需要给 dog和 cat 分别实现 say方法就可以实现 Sayer接口了。

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

func TestI1(t *testing.T) {
	var x Sayer // 声明一个 Sayer 接口类型的变量 x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog

	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵

	x = b       // 可以把dog实例直接赋值给 x
	x.say()     // 汪汪汪
}

// 值接收者和指针接收者实现接口的区别

type Mover interface {
	move()
}

func (d dog) move() {
	fmt.Println("狗会动")
}

func TestI2(t *testing.T) {
	var x Mover

	var d1 = dog{} // d1是dog类型
	x = d1         // x可以接收dog类型
	x.move()

	var d2 = &dog{}  // d2是*dog类型
	x = d2           // x可以接收*dog类型
	x.move()
}


// 接口嵌套 接口与接口间可以通过嵌套创造出新的接口。嵌套得到的接口的使用与普通接口一样！
type animal interface {
	Sayer
	Mover
}

func TestI3(t *testing.T) {
	var x animal
	x = dog{name: "花花"}
	x.move()
	x.say()
}

// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量。
func TestI4(t *testing.T) {
	// 定义一个空接口x
	var x interface{}

	s := "baidu.com"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)

	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)

	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)

}

// 空接口的应用
// 空接口作为函数的参数
// 使用空接口实现可以接收任意类型的函数参数。

// 空接口作为函数参数
func show(a interface{}) {
    fmt.Printf("type:%T value:%v\n", a, a)
}

// 空接口作为map的值类型
// 使用空接口实现可以保存任意值的字典。
func TestI5(t *testing.T) {
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "李白"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值
func TestI6(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = nil
	_ = w

// 	看图解
}

// 想要判断空接口中的值这个时候就可以使用类型断言
//   x.(T)
//  	x：表示类型为 interface{}的变量
//  	T：表示断言x可能是的类型。
// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
func TestI7(t *testing.T) {
	var x interface{}
	x = "string"
	v ,ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	// 如果要断言多次就需要写多个if判断，这个时候我们可以使用switch语句来实现
	switch v1 := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v1)
	case int:
		fmt.Printf("x is a int is %v\n", v1)
	case bool:
		fmt.Printf("x is a bool is %v\n", v1)
	default:
		fmt.Println("unsupport type！")
	}
}