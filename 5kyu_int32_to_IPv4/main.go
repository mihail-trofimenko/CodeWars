package main

import "fmt"

func Int32ToIp(n uint32) string {
	ip4 := n % 256
	ip3 := ((n - ip4) % (65536)) / 256
	ip2 := ((n - ip4 - ip3) % (16777216)) / (65536)
	ip1 := (n - (ip4 + (ip3 * 256) + (ip2 * 65536))) / (16777216)
	return fmt.Sprintf("%d.%d.%d.%d", ip1, ip2, ip3, ip4)
}

/* Best practices:

func Int32ToIp(n uint32) string {
  return fmt.Sprintf("%d.%d.%d.%d", (n >> 24) % 256, (n >> 16) % 256, (n >> 8) % 256, n % 256)
}

*/

func main() {
	fmt.Println(Int32ToIp(2154959208))
}
