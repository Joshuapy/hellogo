// 接口类型训练
// 2022/4/28


package main


import "fmt"


// 接口集
type Sayer interface{
	Say()
}

type Mover interface{
	Move(steps int) (out string)
}

// 猫类 - 方法接受者是值类型
type Cat struct {
	name string
}

func (c *Cat) SetName(name string) (err error) {
	if len(name) > 0 {
		c.name = name
		return
	}
	return fmt.Errorf("empty name!")
}

// 值类型接受者, 可以接受Cat类型和 *Cat指针类型
func (c Cat) Say() {
	fmt.Printf("I am a cat name is %v, and i say 喵喵\n", c.name)
}

func (c Cat) Move(n int) string {
	msg := fmt.Sprintf("I am a cat name is %v, i move %v steps!\n", c.name, n)
	return msg
}


// 狗类 - 方法接受者是值指针类型
type Dog struct {
	name string
}

func (d *Dog) SetName(name string) (err error){
	if len(name) == 0 {
		err = fmt.Errorf("name is empty.")
		return
	}
	d.name = name
	return
}

func (d *Dog) Say() {
	fmt.Printf("I am a dog name is %s, i say 汪汪\n", d.name)
}

func (d *Dog) Move(step int) (string){
	msg := fmt.Sprintf("I am a dog name is %s, i move %d steps!\n", d.name, step)
	return msg
}

// 接口调用
func whoyouare(s Sayer) {
	s.Say()
}
func gothere(m Mover, n int) {
	out := m.Move(n)
	fmt.Print(out)
}


func main() {
	var x Sayer
	var y Mover
	// 值类型实现接口
	x = Cat{name: "小黑"}  // 直接将自定义类型的实例 赋值 给 接口类型的变量
	y = Cat{name: "小灰"}
	whoyouare(x)  // 使用接口类型的值进行传参
	// gothere(x) // 参数类型错误
	gothere(y, 3)

	c := Cat{}
	y = c
	err := c.SetName("")
	if err != nil {
		fmt.Printf("set name error: %s\n", err.Error())
	}
	c.SetName("小白")
	whoyouare(c)  // 不需要使用中间接口类型的变量，可直接将自定义类型实例传参给函数(签名需要接口类型的函数)
	gothere(c, 2)

	// 指针类型实现接口
	wangcai := &Dog{name: "wangcai"}
	whoyouare(wangcai)
	gothere(wangcai, 23)

	whoyouare(Dog{name: "adog"}) // 传参失败 不能接收值类型
}