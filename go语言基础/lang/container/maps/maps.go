package main

//map[k]v
//map[k1]map[k2]v  复合map
/*
定义map的三种方式

1. var m1 map[string]int // map[]
2. m2 := map[string]int{}
3. m3 := make(map[string]int)   //map[]

创建: make(map[string]int)
获取元素: m[key]
key不存在时，获得value类型的初始值
用value, ok:= m[key] 来判断key是否存在
用delete删除key

不保证遍历顺序，如需顺序，需手动对key排序; 将key放到slice中排序后，在根据slice的排序去遍历map
使用len获取元素的个数

除了slice, map, function的内建类型都可以做为key，比如int string bool
struct类型不包含上述字段，也可作为key

*/
import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil    只声明类型
	fmt.Println("m, m2, m3:")
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map m")
	for k, v := range m { //遍历map， 遍历出来的k v是无序的，每次返回的顺序可能都不一样
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"] //获取key为cource的值; 如果获取到一个不存在的key，系统不会抛错,根据类型返回zero value。如果是字符串就是返回空, 数值类型就返回0
	fmt.Println(`m["course"] =`, courseName)
	if causeName, ok := m["cause"]; ok { // ok变量返回的是这个m["casuse"]是否存在
		fmt.Println(causeName)
	} else {
		fmt.Println("key 'cause' does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)

	delete(m, "name") //删除key为name的值
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
}
