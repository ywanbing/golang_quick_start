package internal

/*
	append          -- 用来追加元素到数组、slice中,返回修改后的数组、slice
    delete          -- 从 map中删除 key对应的value
    make            -- 用来分配内存，返回 Type本身(只能应用于 slice, map, channel)
    new             -- 用来分配内存，主要用来分配值类型，比如 int、struct。返回指向 Type的指针
    cap             -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
    copy            -- 用于复制和连接 slice，返回复制的数目
    len             -- 来求长度，比如 string、array、slice、map、channel ，返回长度
    print、println  -- 底层打印函数，在部署环境中建议使用 fmt 包
*/

/*
%v	按值的本来值输出
%+v	在 %v 基础上，对结构体字段名和值进行展开
%#v	输出 Go 语言语法格式的值
%T	输出 Go 语言语法格式的类型和值
%%	输出 % 本体
%b	整型以二进制方式显示
%o	整型以八进制方式显示
%d	整型以十进制方式显示
%x	整型以十六进制方式显示
%X	整型以十六进制、字母大写方式显示
%U	Unicode 字符
%f	浮点数
%p	指针，十六进制方式显示
*/
