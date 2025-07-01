package main

import "fmt"

func PrintMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key, " value = ", value)

	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["china"] = "shanghai"
}

func main() {
	cityMap := make(map[string]string)
	cityMap["china"] = "beijing"
	cityMap["usa"] = "newyork"
	cityMap["japan"] = "tokyo"
	
	PrintMap(cityMap)
	fmt.Println("=====================")
	ChangeValue(cityMap)
	PrintMap(cityMap)

}
