package _map

import (
	"fmt"
	"sort"
	"testing"
)

/*

map是一种无序的基于 key-value 的数据结构，Go 语言中的 map 是引用类型，必须初始化才能使用。
Go语言中 map的定义语法如下
    map[KeyType]ValueType
其中，
    KeyType:表示键的类型。
    ValueType:表示键对应的值的类型。

map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
    make(map[KeyType]ValueType, [cap])

其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

*/

// TestM1
// @description:
// parameter:
//		@t:
// return:
func TestM1(t *testing.T) {
	// map 基本使用
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	// map 也支持在声明的时候填充元素
	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(userInfo)

	// 如果key存在 ,ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 只想遍历 key 的时候,遍历 map 时的元素顺序与添加键值对的顺序无关。
	for k := range scoreMap {
		delete(scoreMap,k)
	}

	fmt.Println(scoreMap)
	fmt.Println("---------------")
	// 使用delete()函数删除键值对  delete(map, key)
	delete(scoreMap, "小明") // 将小明:100 从 map 中删除
	for k,v := range scoreMap{
		fmt.Println(k, v)
	}


}


// 按照指定顺序遍历 map
func TestM2 (t *testing.T) {
	var scoreMap = make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("%02d", i) // 生成字符串
		scoreMap[key] = i
	}

	// 取出map中的所有key存入切片keys
	var keys = make([]string, 0, 10)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	fmt.Println(keys)

	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的 key遍历 map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// map 并不支持并发的读写
func TestM3(t *testing.T) {
	m := make(map[int]int)

	go func() {
		for {
			m[0] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
			// m[1] = 2
		}
	}()

	select {}
}


/*
map 原理部分
	1. 整理存储结构
	2. 初始化
	3. 写入数据
	4. 读取数据
	5. 扩容 和 迁移
 */