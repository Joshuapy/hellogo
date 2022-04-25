// 结构体 练习

package main

import "fmt"

// 基本结构体类型
type person struct {
	name string
	city string
	age  int8
}

// 方法
func (p *person) dream() {
	fmt.Printf("%s is dreamign.\n", p.name)
}

// 方法改值
func (p *person) setAge(age int8) {
	p.age = age
}

func custom_type() {
	type newint int  // 什么场景下使用呢？
	type myint = int // 感觉没啥用
	var a newint     // newint
	var b myint      // int 太鸡肋
	fmt.Printf("type of a is: %T\n", a)
	fmt.Printf("type of a is: %T\n", b)

}

func define_struct() {
	var p1 person
	fmt.Printf("type of p1: %T\n", p1)
	fmt.Printf("p1.name: %#v\n", p1.name)
	p1.name = "joker"
	fmt.Printf("%#v\n", p1)

	var p2 = new(person)
	fmt.Printf("type of p2: %T\n", p2)
	p2.name = "joker2"
	fmt.Printf("%#v\n", p2)

	var p3 = &person{} // 取地址操作  p3是个指针
	fmt.Printf("type of p3: %T\n", p3)

	p33 := person{} // 和p1 有什么区别吗？
	fmt.Printf("p33: %#v\n", p33)

	p4 := person{
		name: "simis",
		age:  18}
	fmt.Printf("p4: %#v\n", p4)

	p5 := person{
		"foo",
		"newyors",
		18}
	fmt.Printf("p5: %#v\n", p5)

}

func mianshiti() {
	// 影子变量 bug
	// 参考：https://medium.com/swlh/use-pointer-of-for-range-loop-variable-in-go-3d3481f7ffc9
	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	// stus := []*student{  // 修复 why? 指针是不变的。  m[key] = stu 是把地址的值赋值给映射所以可以准确读取原始值。
	stus := []student{
		{name: "xx", age: 18}, // 初始化没问题？
		{name: "zz", age: 18},
		{name: "yy", age: 18},
	}
	for _, stu := range stus {
		key := stu.name
		value := stu // 新声明一个变量(新内存空间的地址)
		m[key] = &value
		// m[key] = &stu // TODO 这里有问题。。。??? 影子变量的地址，循环完后是最后一个值的地址
		fmt.Printf("m: %#v\n", m)
		fmt.Printf("%v\n", m[key])
	}

	// fmt.Printf("m: %#v\n", m)

	// fmt.Printf("%#v", m)
	for k, v := range m {
		fmt.Printf("%v: %#v\n", k, *v)
		// fmt.Println(k, "=>", v.name)
	}
}

func struct_method() {
	p := person{name: "joker"}
	p.dream()
	fmt.Printf("age: %d\n", p.age)
	p.setAge(12)
	fmt.Printf("after age: %d\n", p.age)
}

// 嵌套匿名结构体定义和使用
func substruct() {
	type Address struct {
		proince     string
		city        string
		create_time string
	}
	type Message struct {
		title       string
		content     string
		create_time string
	}
	type Mail struct {
		create_time string
	}
	type person struct {
		name string
		age  int8
		addr Address  // 具名嵌套
		Message
		Mail
	}

	p := person{
		name: "joker",
		age:  18,
		addr: Address{proince: "shanghai", city: "SH", create_time: "addr_day"},
		Mail: Mail{"init_msg_time"},
	}
	p.title = "titl" // 通过匿名结构体访问
	p.content = "cnt"
	p.Message.create_time = "msg_time" // 多个匿名结构体同名变量需要指定结构体类型访问
	p.Mail.create_time = "mail_time"
	fmt.Printf("%#v\n", p)
}


// 实验append 的bug
type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) []student {
	ce[1].age = 99  // 因为是引用类型，所以可以修改值
	ce = append(ce, student{3, "baz", 33}) // 小心 append的返回值,
	return ce
}
func slice_struce_append() {

	var ce []student
	ce = []student{
		{1, "foo", 11},
		{2, "bar", 22},
	}
	fmt.Printf("%v\n", ce)
	ce = demo(ce) // append 有新变量生成
	fmt.Printf("%v\n", ce)
}

func main() {
	// custom_type()
	// define_struct()
	// mianshiti()
	// struct_method()
	// substruct()
	slice_struce_append()
}
