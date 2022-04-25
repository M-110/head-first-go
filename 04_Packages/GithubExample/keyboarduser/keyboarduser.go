package keyboarduser

import (
	"fmt"
	"github.com/M-110/head-first-go/04_Packages/Constants/dates"
	"github.com/M-110/head-first-go/04_Packages/keyboard"
)

func main() {
	fmt.Println(keyboard.GetFloat())
	fmt.Println(dates.DaysToWeeks(5))
}
