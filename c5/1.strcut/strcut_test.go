package strcut

import (
	"fmt"
	"testing"
)

/*
结构体的定义：
	type 类型名 struct {
        字段名 字段类型
        字段名 字段类型
        …
    }
    1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
    2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
    3.字段类型：表示结构体字段的具体类型。
*/

type person struct {
	name string
	city string
	age  int8
}

func TestS1(t *testing.T) {
	var p = person{
		name: "",
		city: "",
		age:  0,
	}

	// 结构体初始化 = 上面
	p = person{}
	_ = p.name

	// 创建指针并赋值,使用 & 对结构体进行取地址操作相当于对该结构体类型进行了一次 new 实例化操作。
	p1 := &person{}
	_ = p1.name

	// 通过 new 函数新建 得到的是结构体的地址。
	p1 = new(person)
	_ = p1.name

	// 当某些字段没有初始值的时候，该字段可以不写。此时，
	// 没有指定初始值的字段的值就是该字段类型的零值。
	p2 := &person{
		city: "北京",
	}
	fmt.Printf("p2=%#v\n", p2) // p2=&main.person{name:"", city:"北京", age:0}

	// 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值。
	p3 := &person{
		"xqw",
		"北京",
		18,
	}
	fmt.Printf("p3=%#v\n", p3) // p3=&main.person{name:"xqw", city:"北京", age:18}
	/*
			1.必须初始化结构体的所有字段。
		    2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
		    3.该方式不能和键值初始化方式混用。
	*/
}

// go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入结构体
// 嵌入式结构体
type Programmer struct {
	person
	company string
}

func TestS2(t *testing.T) {
	var p Programmer

	p.name = "abc"
	p.person.name = "def"
	fmt.Printf("%#v \n", p)

	// 初始化 的方式
	s1 := Programmer{
		person{"5lmh", "man", 20},
		"bj",
	}
	fmt.Println(s1)

	s2 := Programmer{person: person{"5lmh", "man", 20}}
	fmt.Println(s2)

	s3 := Programmer{person: person{name: "5lmh"}, company: "bj"}
	fmt.Println(s3)
}

// 同名字段的情况
type Student struct {
	//  person {
	//    name string
	//    city string
	//    age  int8
	// }
	person
	name string
}

func TestS3(t *testing.T) {
	var s Student
	// 给自己字段赋值了
	s.name = "5lmh"
	fmt.Println(s)

	// 若给父类同名字段赋值，如下
	s.person.name = "abc"
	fmt.Println(s)
}

// 所有的内置类型和自定义类型都是可以作为匿名字段去使用
type Any struct {
	int
	int32
	bool
	string
	Student
	*person
}

// 结构体比较：
// 如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，
// 那样的话两个结构体将可以使用==或!=运算符进行比较。相等比较运算符==将比较两个结构体的每个成员。
// 可比较的结构体类型和其他可比较的类型一样，可以用于map的key类型。

// 内存对齐 .....
/*
Go struct 内存对齐：https://geektutu.com/post/hpg-struct-alignment.html
详解内存对齐：https://mp.weixin.qq.com/s/ig8LDNdpflEBWlypU1NRhw

视频讲解：
为什么要内存对齐：https://www.bilibili.com/video/BV1Ja4y1i7AF
*/
