package basic

import "fmt"

/**
 * 类型断言
 * 有了接口和实现接口的类型，就会有类型断言。类型断言用来判断一个接口的值是否是实现该接口的某个具体类型
 */
type person1 struct {
	name string
	age  uint
	address1
}

type address1 struct {
	province string
	city     string
}

func NewPerson1(name string, age uint, addr address1) *person1 {
	return &person1{name: name, age: age, address1: addr}
}

func NewAddress1(province string, city string) address1 {
	return address1{province: province, city: city}
}

/**
 * 实现Stringer接口，Stringer是Go SDK的一个接口，属于fmt包(src/fmt/print.go)
 * type Stringer interface {
 *     String() string
 * }
 */
func (p *person1) String() string {
	return fmt.Sprintf("name:%s,age:%d,province:%s,city:%s", p.name, p.age, p.province, p.city)
}

func (addr address1) String() string {
	return fmt.Sprintf("province:%s,city:%s", addr.province, addr.city)
}

// person1和address1类型都实现了Stringer接口
func TypeAssertionDemo() {
	var s fmt.Stringer
	p1 := NewPerson1("Tom", 20, address1{province: "广东", city: "深圳"})
	s = p1
	// 判断接口s的值是否是*person1类型，这就是类型断言
	p2 := s.(*person1)
	fmt.Println(p2)
	/**
	 * 类型断言的多值返回
	 * address1也实现了Stringer接口，如果对s进行address1类型断言，就会抛出异常信息
	 * 程序异常这不符合设计初衷，本来想判断一个接口的值是否是某个具体类型，但不能因为判断失败就导致程序异常
	 * 基于这点，Go语言提供了类型断言的多值返回，类型断言返回的第二个值就是断言是否成功的标志(true/false)
	 */
	// a := s.(address1)
	// panic: interface conversion: fmt.Stringer is *basic.person1, not basic.address1
	// fmt.Println(a)
	a, ok := s.(address1)
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("s不是一个address1类型的值")
	}
}
