// map 练习

package main

import (
	"fmt"
)

func create_map() {
	var map1 = make(map[string]int)
	var map2 map[string]int
	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	//map2["foo"] = 1  // 不能分配 panic
	map1["foo"] = 1 // 不能分配 panic
	fmt.Println(map1 == nil)
}

func del_key() {
	map1 := map[string]int{"foo": 1, "bar": 2}
	delete(map1, "foo1")
	fmt.Printf("%v", map1)
}

func access_key() {
	map1 := map[string]int{"foo": 1, "bar": 2}
	foo := map1["foo"]
	bar, isBarExisted := map1["bar"]
	bar2, isBarExisted2 := map1["bar2"]
	fmt.Printf("%d, %d, %t", foo, bar, isBarExisted)
	fmt.Printf("%d, %t", bar2, isBarExisted2)
}

func slicemap() {
	// 值是map的slice
	sliceMap := make([]map[string]string, 5)
	for index, value := range sliceMap {
		isnil := value == nil
		fmt.Printf("index: %d, value: %v ,isnil:%v\n", index, value, isnil)
	}
	sliceMap[0] = make(map[string]string, 10)
	sliceMap[0]["foo"] = "v_foo"
	sliceMap[0]["bar"] = "v_bar"
	for index, value := range sliceMap {
		isnil := value == nil
		fmt.Printf("index: %d, value: %v ,isnil:%v\n", index, value, isnil)
	}
}

func mapslice() {
	// 值是slice的map,  先声明在逐个填充
	mapslice := make(map[string][]string, 3)
	isnil := mapslice == nil
	fmt.Printf("%v, isnil:%v\n", mapslice, isnil)

	// mapslice["name"] = []string{"1", "2"}
	// fmt.Printf("%v\n", mapslice)
	key1 := "foo"
	value1 := make([]string, 3)
	el1 := "e_foo1"
	el2 := "e_foo2"
	el3 := "e_foo3"
	value1 = append(value1, el1, el2, el3)
	mapslice[key1] = value1

	// 初始的时候做内容填充
	mapslice2 := map[string][]string{
		"name": []string{"1", "2"},
		"foo":  []string{"2", "2"}}
	fmt.Printf("mapslice2: %v\n", mapslice2)
}

func func_mapmap() {
	// {"foor": {"k1":1, "k2":2},
	// "bar": {"k1":1, "k2":2}
	// }
	mapmap := make(map[string]map[string]int, 2)
	isnil := mapmap == nil
	fmt.Printf("%v, %v", mapmap, isnil)

	fmt.Println("after init:")

	key1 := "foo"
	submap := make(map[string]int, 2)
	submap["k1"] = 1
	submap["k2"] = 2
	mapmap[key1] = submap

	key2 := "bar"
	submap2 := make(map[string]int, 2)
	submap2["k1"] = 1
	submap2["k2"] = 2
	mapmap[key2] = submap2

	fmt.Printf("%v", mapmap)

}

func main() {
	//create_map()
	//del_key()
	// access_key()
	// slicemap()
	func_mapmap()
}
