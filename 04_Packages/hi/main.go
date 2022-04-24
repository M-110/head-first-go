package main

import (
	"fmt"
	"greeting/greeting"
	"mypck/birthdays"
	"mypck/holidays"
)

func main() {
	holiday := holidays.Holiday
	birthday := birthdays.Birthday
	fmt.Println(holiday)
	fmt.Println(birthday)
	greeting.Hello()
	greeting.Hi()
	greeting.AllGreetings()
}
