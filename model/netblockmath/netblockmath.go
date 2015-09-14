package netblockmath

import (
	"fmt"
	"net"
)

func NetblocksOverlap(a_start, a_end, b_start, b_end *net.IP) bool {
	starta_int := Inet_aton(a_start)
	enda_int := Inet_aton(a_end)
	startb_int := Inet_aton(b_start)
	endb_int := Inet_aton(b_end)

	// if startb is between starta and enda
	if between(starta_int, enda_int, startb_int) ||
		between(starta_int, enda_int, endb_int) {
		return true
	}

	// or endb is between starta and enda
	if between(startb_int, endb_int, starta_int) ||
		between(startb_int, endb_int, enda_int) {
		return true
	}

	return false
}

func between(lower, upper, x uint32) bool {
	if x >= lower && x <= upper {
		return true
	}

	return false
}

func Inet_aton(ip *net.IP) uint32 {
	ip_byte := ip.To4()
	return uint32(ip_byte[0])<<24 | uint32(ip_byte[1])<<16 |
		uint32(ip_byte[2])<<8 | uint32(ip_byte[3])
}

func Inet_ntoa(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}
