package main

import "fmt"

func Arrange(s []int) []int {

	var t []int
	var a, b int

	if s == nil {
		return nil
	}

	if len(s) == 1 {
		return s
	}

	b = len(s) - 1
	for {
		t = append(t, s[a])
		t = append(t, s[b])
		a++
		if a >= b {
			break
		}
		b--
		if a >= b {
			break
		}
		t = append(t, s[b])
		t = append(t, s[a])
		b--
		if a >= b {
			break
		}
		a++
		if a >= b {
			break
		}
	}
	if len(s)%2 == 1 {
		t = append(t, s[a])
	}
	return t
}

func main() {
	fmt.Println(Arrange([]int{1, 2, 3, 0, 4, 5, 6}))
}
