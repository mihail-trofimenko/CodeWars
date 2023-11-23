package main

import "fmt"

func trim(n int) int {
	if n < 0 {
		n = 0
	}
	if n > 255 {
		n = 255
	}
	return n
}

func RGB(r, g, b int) string {
	r = trim(r)
	g = trim(g)
	b = trim(b)
	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}
