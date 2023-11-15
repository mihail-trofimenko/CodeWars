package main

import "fmt"

func FirstNonRepeating(str string) string {

	for j, r0 := range str {
		repeating := false
		for i, r1 := range str {
			if ((r1 >= 'A') && (r1 <= 'z')) && ((r0 >= 'A') && (r0 <= 'z')) && ((r1 == r0+32) || (r0 == r1+32) || r0 == r1) && i != j {
				repeating = true
				break
			}
			if ((r0 < 'A') || (r0 > 'z')) && r0 == r1 && i != j {
				repeating = true
				break
			}
		}
		if !repeating {
			return string(r0)
		}
	}

	return ""
}

func main() {
	fmt.Println(FirstNonRepeating("ZBgAOv:Al:c;eu4YcM:guJ7"))
}
