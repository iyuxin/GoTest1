package main

import "fmt"

func main() {
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("map is nil")
	}

	var map1 = make(map[string]string, 10)

	map1["one"] = "java"
	map1["two"] = "python"
	map1["three"] = "golang"
	fmt.Println(map1)
	fmt.Println("=========================")

	map2 := make(map[int]string)

	map2[1] = "java"
	map2[2] = "python"
	map2[3] = "golang"
	fmt.Println(map2)
	fmt.Println("=========================")

	map3 := map[int]string{
		1: "java",
		2: "python",
		3: "golang",
	}
	fmt.Println(map3)
}
