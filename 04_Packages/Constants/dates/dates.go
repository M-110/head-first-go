package dates

const daysInWeek int = 7

func WeeksToDays(weeks int) int {
	return weeks * daysInWeek
}

func DaysToWeeks(days int) float64 {
	return float64(days) / float64(daysInWeek)
}
