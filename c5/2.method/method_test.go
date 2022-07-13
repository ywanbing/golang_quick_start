package method

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
	Go语言中的方法（Method）是一种作用于特定类型变量的函数。
	这种特定类型变量叫做接收者（Receiver）。
	接收者的概念就类似于其他语言中的this或者 self。

	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
        函数体
    }

    1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。
		例如，Person 类型的接收者变量应该命名为 p，Cat 类型的接收者变量应该命名为 c 等。
    2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
    3.方法名、参数列表、返回参数：具体格式 与 函数定义相同。
*/

type Person struct {
	name string
	age  int
}

// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func TestM1(t *testing.T) {
	p1 := Person{"测试", 25}
	p1.Dream()
}

/*
指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，
在方法结束后，修改都是有效的。
这种方式就十分接近于其他语言中面向对象中的this或者self。

当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
*/

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

func TestM2(t *testing.T) {
	p1 := &Person{"测试", 25}
	fmt.Println(p1.age) // 25

	p1.SetAge(30)
	fmt.Println(p1.age) // 30
}

// 什么时候应该使用指针类型接收者
/*
 	1.需要修改接收者中的值
    2.接收者是拷贝代价比较大的大对象
    3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/

// 在 Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
// 注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

// MyInt 将int定义为自定义MyInt类型
type MyInt int

// SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

// 结构体的“继承” ,Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() { fmt.Printf("%s会动！\n", a.name) }

func (a *Animal) selfName() { fmt.Printf("%s \n", a.name) }

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构体实现继承
}

func (d *Dog) move() { fmt.Printf("%s会跑~\n", d.name) }

func (d *Dog) wang() { fmt.Printf("%s会汪汪汪~\n", d.name) }

func TestM3(t *testing.T) {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ // 注意嵌套的是结构体指针
			name: "小花",
		},
	}

	d1.selfName()
	d1.Animal.move() // 小花会动！

	d1.move() // 小花会跑！
	d1.wang() // 小花会汪汪汪~
}

// 结构体标签（Tag）

/*
Tag 是结构体的元信息，可以在运行的时候通过反射的机制读取出来。

Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
	`key1:"value1" key2:"value2"`
结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。
注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。
结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。
例如不要在key和value之间添加空格。
*/

type Teacher struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	CC   int    `json:"-"`
	abc  string `json:"abc"`
}

type teacher2 struct {
	Name string
	Age  int
	CC   int
	abc  string
}

func TestM4(t *testing.T) {
	var te Teacher
	m1, _ := json.Marshal(te)
	fmt.Println(string(m1))

	var te2 teacher2
	m2, _ := json.Marshal(te2)
	fmt.Println(string(m2))
}
