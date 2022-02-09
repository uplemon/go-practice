package quickstart

import "fmt"

/**
 * 结构体是一种聚合类型，里面可以包含任意类型的值，这些值就是我们定义的结构体的成员(也称为字段)
 * 在Go语言中，要自定义一个结构体，需要使用type+struct关键字组合
 * type structName struct {
 *     fieldName typeName
 *     ...
 *     ...
 * }
 * 总结：结构体是一种聚合类型，它比普通类型可以携带更多数据
 */

type Person struct {
    name string
    age uint
    addr address
}

type address struct {
    province string
    city string
}

/**
 * 工厂函数一般用于创建自定义的结构体，便于使用者调用
 * 通过工厂函数创建自定义结构体的方式，可以让调用者不用太关注结构体内部的字段，只需要给工厂函数传参就可以了
 */
func NewPerson(name string, age uint, addr address) *Person {
    return &Person{name: name, age: age, addr: addr}
}

func StructDemo() {
    p1 := Person{name: "Tom", age: 25, addr: address{province: "浙江", city: "杭州"}}
    p2 := Person{
        name: "Lina",
        age: 20,
        addr: address{
            province: "北京",
            city: "北京",
        },
    }
    p3 := NewPerson("Pony", 30, address{province: "山东", city: "济南"})
    fmt.Println(p1.name, p1.age, p1.addr.province, p1.addr.city)
    fmt.Println(p2.name, p2.age, p2.addr.province, p2.addr.city)
    fmt.Println(p3.name, p3.age, p3.addr.province, p3.addr.city)
    // 组合代替继承
    StructExtendsDemo()
}

/**
 * 在Go中没有继承的概念，结构/接口之间没有父子继承关系，Go语言提倡的是组合，利用组合达到代码复用的目的
 * 接口可以组合，结构体也可以组合
 * 把接口名/结构体名直接放到目标接口/结构体中就是组合
 */
type person struct {
    name string
    age uint
    // 把address结构体组合到结构体person中(不是当成一个字段)
    // 组合后，被组合的address称为内部类型，person称为外部类型
    // 外部类型不仅可以使用内部类型的字段，也可以使用内部类型的方法，就跟使用自己的字段和方法一样
    // 因为person组合了address，所以address的字段就像person自己的一样，可以直接使用
    address
}

func StructExtendsDemo()  {
    p := person{name: "Mike", age: 26, address: address{province: "北京", city: "北京"}}
    fmt.Println(p)
    fmt.Println(p.name, p.age, p.province, p.city)
}
