package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// to get the time in format we must need to use 01-02-2006 it is month-day-year and for time 15:04:05 it is hour-minute-second and for day Monday to be must and it must follow the case Sensetive.

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	//here the month should have time.---- while day time minute second are in int

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

// OUTPUT

// Welcome to time study of golang
// 2022-10-13 15:25:08.4757749 +0530 IST m=+0.003499401
// 10-13-2022 15:25:08 Thursday
// 2020-08-10 23:23:00 +0000 UTC
// 08-10-2020 Monday
