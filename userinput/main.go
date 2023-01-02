package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("enter your name: ")
	kr()
}
func kr() {
	a := bufio.NewReader(os.Stdin)
	b, _ := a.ReadString('\n')
	fmt.Println("Your name is: ", b)
}
