package main

import (
	"fmt"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("新协程：i = ", i)

	}
}
func main() {
	go newTask()

	i := 0
	for {
		i++
		fmt.Println("主协程：i = ", i)

	}
}
