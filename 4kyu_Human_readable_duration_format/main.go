package main

import "fmt"

func addS(n int64) string {
	if n > 1 {
		return "s"
	}
	return ""
}

func addSepar(s []string) string {
	// chech if value is last or pre-last in array
	cnt := 0
	if s[0] == "" {
		return ""
	}
	for i := 1; i < len(s); i++ {
		if s[i] != "" {
			cnt++
		}
	}
	if cnt == 1 {
		return " and "
	}
	if cnt > 1 {
		return ", "
	}
	return ""

}

func FormatDuration(seconds int64) string {

	var M int64 = 60       // 	seconds per minute
	var H int64 = 3600     // 	seconds per hour
	var D int64 = 84600    // 	seconds per day
	var Y int64 = 31536000 // 	seconds per year

	fmt.Println("Seconds:", seconds)
	years := int64(seconds / Y)
	seconds -= years * Y
	days := int64(seconds / D)
	seconds -= days * D
	hours := int64(seconds / H)
	seconds -= hours * H
	minutes := int64(seconds / M)
	seconds -= minutes * M
	var YDHMS = []string{"", "", "", "", ""}

	if years > 0 {
		YDHMS[0] += fmt.Sprint(years) + " year" + addS(years)
	}

	if days > 0 {
		YDHMS[1] += fmt.Sprint(days) + " day" + addS(days)
	}

	if hours > 0 {
		YDHMS[2] += fmt.Sprint(hours) + " hour" + addS(hours)
	}

	if minutes > 0 {
		YDHMS[3] += fmt.Sprint(minutes) + " minute" + addS(minutes)
	}

	if seconds > 0 {
		YDHMS[4] += fmt.Sprint(seconds) + " second" + addS(seconds)
	}

	for i := 0; i < len(YDHMS)-1; i++ {
		Separ := addSepar(YDHMS[i:])
		YDHMS[i] += Separ
	}

	return fmt.Sprintf("%s%s%s%s%s", YDHMS[0], YDHMS[1], YDHMS[2], YDHMS[3], YDHMS[4])
}

func main() {
	const M, H, D, Y int64 = 60, 3600, 84600, 31536000
	fmt.Println(FormatDuration(0*Y + 182*D + 1*H + 44*M + 40))
	fmt.Println(FormatDuration(15731080))
}
