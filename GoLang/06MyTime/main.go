package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study ofngolang")

	presentTime := time.Now().AddDate(2006, 02, 01)
	fmt.Println("time is ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04 Monday"))

	createdDate := time.Date(2020, time.April, 5, 14, 12, 32, 21, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
