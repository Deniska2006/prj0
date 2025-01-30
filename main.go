package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Добрий день,як вас звати?")
	var name string
	fmt.Scan(&name)
	if name == "Denis" {
		fmt.Println("Ви крутий!")
	} else {
		fmt.Println("Ви lox")
	}
	time.Sleep(7 * time.Second)
}
