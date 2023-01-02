package main

import (
	"fmt"
	"time"
)

func main() {
	current_day := time.Now().Weekday()
	current_time := time.Now()
	fmt.Println("Current time is: ", current_time)
	fmt.Println("Current day is: ", current_day)
}
