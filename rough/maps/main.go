package main

import (
	"fmt"
	"maps"
)

func main(){
	
	// m := make(map[string]string)

	// setting elements 

	// m["name"] = "golang"
	// m["area"] = "backend"

	// get element
	// fmt.Println(m["name"])
	// fmt.Println(m["area"])
	// fmt.Println(m["phone"]) //if key does not exists in map it returns zeroed value4

	// fmt.Println(len(m))

	// delete(m, "area")
	// fmt.Println(m["area"])
	// clear(m)

	// fmt.Println(m)

	// v, ok := m["name"]

	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("All ok")
	// }else{
	// 	fmt.Println("Not ok")
	
	m1 := map[string]int{"age":20,"phones":3}
	m2 := map[string]int{"age":20,"phones":3}

	fmt.Println(maps.Equal(m1,m2))
}