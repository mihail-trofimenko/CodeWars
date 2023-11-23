package main

import "fmt"

func HumanReadableTime(seconds int) string {
	hours := int(seconds / 3600)
	minutes := int((seconds - hours*3600) / 60)
	seconds = seconds - (hours*3600 + minutes*60)
	var HH, MM, SS string

	if hours < 10 {
		HH = "0"
	}
	HH += fmt.Sprint(hours) + ":"
	if minutes < 10 {
		MM = "0"
	}
	MM += fmt.Sprint(minutes) + ":"
	if seconds < 10 {
		SS = "0"
	}
	SS += fmt.Sprint(seconds)

	return HH + MM + SS
}

func main() {
	fmt.Println(HumanReadableTime(60))
}
