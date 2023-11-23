package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countIP(ipstr string) uint32 {
	ipv4str := strings.Split(ipstr, ".")
	var ipv4 [4]int32
	for i, r := range ipv4str {
		x, _ := strconv.Atoi(r)
		ipv4[i] = int32(x)
	}
	return uint32(ipv4[3] + (ipv4[2] * 256) + (ipv4[1] * 65536) + (ipv4[0] * 16777216))
}

func IpsBetween(start, end string) int {
	return int(countIP(end) - countIP(start))
}

func main() {
	fmt.Println(countIP("128.114.17.104"))
	fmt.Println(IpsBetween("10.0.0.0", "10.0.0.50"))
}

/* Best practices:

func IpsBetween(start, end string) int {
  convertToInt := func(s string) int {
		strArr := strings.Split(s, ".")
		oct1, _ := strconv.Atoi(strArr[0])
		oct2, _ := strconv.Atoi(strArr[1])
		oct3, _ := strconv.Atoi(strArr[2])
		oct4, _ := strconv.Atoi(strArr[3])

		return oct1<<24 + oct2<<16 + oct3<<8 + oct4
	}

	intConvert1 := convertToInt(start)
	intConvert2 := convertToInt(end)

	return intConvert2 - intConvert1
}

*/
