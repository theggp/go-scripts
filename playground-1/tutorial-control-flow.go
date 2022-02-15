package main

import (
	"fmt"
	"time"
)
func TutorialFor() {

	defer fmt.Println("This concludes the tutorial")

	sum_1 := 0
	for i := 0; i < 10; i++ {
		sum_1 += i
	}
	fmt.Println(sum_1)

	// For is Go's "while"
	sum_2 := 1
	for sum_2 < 1000 {
	sum_2 += sum_2
	}
	if sum_2 > 1000 {
		fmt.Println(sum_2)
	} else {
		fmt.Printf("Sum smaller than %v\n", 1000)
	}

	var string_1 string = "Peach"
	switch string_1{
		case "apple": fmt.Println("It was apple")
		case "pear": fmt.Println("It was pear")
		default: fmt.Println("It was not a known one")
	}



	// tutorials time
	today_1 := time.Now().Weekday()
	tomorrow_1 := today_1+1
	hour_now := time.Now().Hour()
	minute_now := time.Now().Minute()
	fmt.Println(tomorrow_1, hour_now, ":", minute_now)

}
