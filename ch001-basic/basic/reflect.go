package basic

import (
    "fmt"
    "reflect"
)

type User struct {
    Name   string `json:"name" xml:"name"`
    Age    int16  `json:"age" xml:"age"`
    Gender string `json:"gender" xml:"gender"`
}

func ReflectDemo() {
    reflectTypeDemo()
    reflectValueDemo()
}

// 通过反射获取类型信息
func reflectTypeDemo() {
    user := User{Name: "Tom", Age: 20, Gender: "male"}
    // 使用reflect.TypeOf()函数可以获得任意值的类型对象(reflect.Type)
    // 程序通过类型对象可以访问任意值的类型信息
    t := reflect.TypeOf(user)
    // reflect.Type.Name()
    // reflect.Type.Kind()
    // reflect.Type.NumField()
    fmt.Printf("[reflect.Type.Name]:%v\n[reflect.Type.Kind]:%v\n[reflect.Type.NumField]:%v\n", t.Name(), t.Kind(), t.NumField())
    // Output:
    // [reflect.Type.Name]:User
    // [reflect.Type.Kind]:struct
    // [reflect.Type.NumField]:2

    // 遍历结构体所有成员
    for i := 0; i < t.NumField(); i++ {
        // 获取每个成员的结构体字段类型
        fieldType := t.Field(i)
        // 输出成员名和tag
        fmt.Printf("name:%v, tag:%v\n", fieldType.Name, fieldType.Tag)
        // 解析Tag
        fmt.Printf("json:%v, xml:%v\n", fieldType.Tag.Get("json"), fieldType.Tag.Get("xml"))
    }

    // 通过字段名，找到字段类型信息
    if fieldType, ok := t.FieldByName("Age"); ok {
        // 解析Tag
        fmt.Printf("[Age] => json:%v, xml:%v\n", fieldType.Tag.Get("json"), fieldType.Tag.Get("xml"))
    }
}

// 通过反射获取值信息
func reflectValueDemo() {
    // 反射不仅可以获取值的类型信息，还可以通过reflect.Value动态地获取或者设置变量的值
    // 变量、interface{}和reflect.Value是可以相互转换的
    user := User{Name: "Tom", Age: 20, Gender: "male"}
    // 从接口变量到反射对象
    v1 := reflect.ValueOf(user)
    // 从反射对象到接口变量，并通过类型断言转换
    u1 := v1.Interface().(User)
    fmt.Println(v1, v1.FieldByName("Name").String())
    fmt.Println(u1)

    // 通过反射修改变量的值，需要传递变量的指针创建反射对象，保证其值是可写的
    v2 := reflect.ValueOf(&user)
    v3 := v2.Elem()
    v3.FieldByName("Name").SetString("Mike")
    fmt.Println(user)
}

/**
 * 反射的三大法则：
 * 1. 反射可以将"接口类型变量"转换为"反射类型对象"
 *    通过reflect.Type()和reflect.Value()将接口类型变量转换为反射类型对象
 * 2. 反射可以将"反射类型对象"转换为"接口类型变量"
 *    根据一个reflect.Value类型的变量，可以使用Interface()方法恢复其接口类型的值，
 *    该方法会把type和value信息打包并填充到一个接口变量中，然后返回
 * 3. 如果要修改"反射类型对象"，则其值必须是"可写的"
 *    创建反射对象时传入的变量是指针，使用Elem()方法返回指针指向的数据
 */
