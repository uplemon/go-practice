package basic

import (
	"fmt"
	"unicode/utf8"
)

// array slice map
func CollectionDemo() {
	ArrayDemo()
	SliceDemo()
	MapDemo()
	StringDemo()
}

/**
 * 数组存放的是固定长度、相同类型的数据，而且这些存放的元素是连续的
 * 所存放的数据类型没有限制，可以是整型、字符串甚至自定义
 */
func ArrayDemo() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr)
	fmt.Println(arr[0], arr[1], arr[2], arr[3], arr[4])
	// 在定义数组的时候，数组的长度可以省略，Go语言会自动根据大括号{}中元素的个数推导出长度
	arr1 := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr1)
	// 数组可以部分初始化，未初始化的元素是数组类型的零值
	arr2 := [5]string{1: "b", 3: "d"}
	fmt.Println(arr2)
	// 遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Printf("数组索引:%d, 对应值:%s\n", i, arr[i])
	}
	// 通过for range简化for循环，range表达式返回两个结果(第一个是数组的索引，第二个是数组的值)
	for i, v := range arr {
		fmt.Printf("数组索引:%d, 对应值:%s\n", i, v)
	}
	/**
	 * 使用for range遍历时，如果返回的值用不到，可以使用_下划线丢弃
	 * 数组的索引通过_就被丢弃了，只使用数组的值v即可
	 */
	for _, v := range arr {
		fmt.Printf("对应值:%s\n", v)
	}
}

/**
 * 切片和数组类似，可以把它理解为动态数组，其实切片就是基于数组实现的
 * 切片是基于数组实现的，它的底层就是一个数组，对数组任意分隔就可以得到一个切片
 * 在Go中切片是使用最多的，尤其是作为函数的参数时，相比数组通常会选择切片，因为它高效、内存占用小
 */
func SliceDemo() {
	/**
	 * 基于数组生成切片，包含索引start，但不包含索引end
	 * slice := array[start:end]
	 * 虽然切片底层用的是arr数组，但对数组切片后，切片的索引会重置
	 * 切片是一个具备三个字段的数据结构，分别是指向数组的指针data、长度len和容量cap
	 */
	arr := [5]string{"a", "b", "c", "d", "e"}
	slice := arr[2:5]
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice[0], slice[1], slice[2])

	/**
	 * 切片修改
	 * 切片的值也可以被修改，这里也同时可以证明切片的底层是数组
	 */
	slice[2] = "f"
	// 可以看到arr也被修改了
	fmt.Println(slice, arr)

	/**
	 * 切片声明
	 * 除了可以从一个数组得到切片外，还可以通过make函数声明切片或者通过字面量方式声明和初始化切片
	 * 切片会有长度和容量，当切片的长度要超过容量的时候，会进行扩容
	 */
	slice1 := make([]string, 3)
	fmt.Println(len(slice1), cap(slice1))
	slice2 := make([]string, 3, 6)        // len:3 cap:3
	fmt.Println(len(slice2), cap(slice2)) // len:3 cap:6
	slice3 := []string{"a", "b", "c"}
	fmt.Printf("len=>%d, cap=>%d\n", len(slice3), cap(slice3))

	/**
	 * 因为共用底层数组，对一个切片的修改会影响到底层数组及基于该数组的其他切片的修改
	 * 当切片扩容时，会生成新的底层数组，从而和原有数组分离
	 * 在创建切片时最好让其长度和容量一样，这样在追加操作时会自动生成新数组，就不会因为共用底层数组而影响多个切片
	 */
	arr1 := [3]string{"a", "b", "c"}
	s1 := arr1[:2]
	fmt.Printf("s1_len=>%d, s1_cap=>%d\n", len(s1), cap(s1))
	s2 := append(s1, "d")
	s3 := append(s1, "e", "f")
	fmt.Println(arr1, s1, s2, s3)

	// 切片的循环和数组一模一样，常用的也是for range方式
	for i, v := range s3 {
		fmt.Printf("%d=>%s ", i, v)
	}
}

/**
 * 在Go中map是一个无序的K-V键值对集合，结构为map[K]V
 * map中所有的key必须具有相同的类型，value也同样，但key和value的类型可以不同
 * key的类型必须支持==比较运算符，这样才可以判断它是否存在，并保证key的唯一性
 */
func MapDemo() {
	m1 := make(map[string]int)
	m2 := map[string]int{}
	fmt.Println(m1, m2)
	m1["Tom"] = 20
	m1["Lina"] = 18
	m1["Mike"] = 25
	fmt.Println(m1)

	/**
	 * 获取map中元素时，如果key不存在，返回的value是该类型的零值，比如int的零值就是0
	 * 通常我们需要判断map中的key是否存在
	 * map的[]操作符可以返回两个值：1.对应的value 2.标记该值是否存在(存在为true，不存在为false)
	 */
	val, ok := m1["Mike"]
	if ok {
		fmt.Println(val)
	}

	// 删除键值对
	delete(m1, "Mike")
	fmt.Println(m1)

	// map的遍历使用for range循环，返回两个值，第一个是map的key，第二个是map的value
	for k, v := range m1 {
		fmt.Printf("key:%s value:%d\n", k, v)
	}

	/**
	 * map的大小：map没有容量，只有长度，也就是map的大小(键值对的个数)
	 * 要获取map的大小，使用内置的len函数即可
	 */
	fmt.Println(len(m1), len(m2))
}

/**
 * 字符串string也是一个不可变的字节序列，所以可以直接转为字节切片[]byte
 * 字符串是类型为byte的只读切片，一个字符串就是一堆字节，字符串存储的是字符的字节
 * string不仅可以直接转为[]byte，还可以使用[]操作符获取指定索引的字节值
 */
func StringDemo() {
	s := "Hello世界"
	bs := []byte(s)
	// 因字符串是字节序列，每一个索引对应的是一个字节，在UTF8编码下，一个汉字对应三个字节，因此字符串长度为11
	fmt.Println(s, len(s), s[0], s[1], s[10]) // Hello世界 11 72 101 140
	fmt.Println(bs, len(bs))                  // [72 101 108 108 111 228 184 150 231 149 140] 11
	// 如果想把一个汉字当成一个长度计算，可以使用utf8.RuneCountInString函数
	// 该字符串为7个unicode(utf8)字符，和我们看到的字符的个数一致
	fmt.Println(utf8.RuneCountInString(s), utf8.RuneCount(bs))
	/**
	 * 使用for range对字符串循环时，也是按照unicode字符进行循环的
	 * 在下面的示例中，i是索引，r是unicode字符对应的unicode码点
	 * 这也说明了for range循环在处理字符串的时候，会自动地隐式解码unicode字符串
	 */
	for i, r := range s {
		fmt.Println(i, r, string(r))
	}
}
