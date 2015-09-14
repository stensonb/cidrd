package netblockmath

import (
	"net"
	"testing"
)

func TestBitMath(t *testing.T) {
	ip := "10.20.30.40"
	ip_parsed := net.ParseIP(ip)
	if ip != Inet_ntoa(Inet_aton(&ip_parsed)) {
		t.Fail()
	}
}

func TestSeparate(t *testing.T) {
	oneStartingIP := net.ParseIP("10.20.30.40")
	oneEndingIP := net.ParseIP("10.20.30.50")

	twoStartingIP := net.ParseIP("10.20.30.60")
	twoEndingIP := net.ParseIP("10.20.30.70")

	ans := NetblocksOverlap(&oneStartingIP, &oneEndingIP, &twoStartingIP, &twoEndingIP)
	if ans {
		t.Fail()
	}
}

func TestFirstContainsSecond(t *testing.T) {
	oneStartingIP := net.ParseIP("10.20.30.40")
	oneEndingIP := net.ParseIP("10.20.30.50")

	twoStartingIP := net.ParseIP("10.20.30.45")
	twoEndingIP := net.ParseIP("10.20.30.47")

	ans := NetblocksOverlap(&oneStartingIP, &oneEndingIP, &twoStartingIP, &twoEndingIP)
	if !ans {
		t.Fail()
	}
}

func TestSecondContainsFirst(t *testing.T) {
	oneStartingIP := net.ParseIP("10.20.30.45")
	oneEndingIP := net.ParseIP("10.20.30.47")

	twoStartingIP := net.ParseIP("10.20.30.40")
	twoEndingIP := net.ParseIP("10.20.30.50")

	ans := NetblocksOverlap(&oneStartingIP, &oneEndingIP, &twoStartingIP, &twoEndingIP)
	if !ans {
		t.Fail()
	}
}

func TestFirstContainsStartOfSecond(t *testing.T) {
	oneStartingIP := net.ParseIP("10.20.30.40")
	oneEndingIP := net.ParseIP("10.20.30.50")

	twoStartingIP := net.ParseIP("10.20.30.45")
	twoEndingIP := net.ParseIP("10.20.30.55")

	ans := NetblocksOverlap(&oneStartingIP, &oneEndingIP, &twoStartingIP, &twoEndingIP)
	if !ans {
		t.Fail()
	}
}

func TestFirstContainsEndOfSecond(t *testing.T) {
	oneStartingIP := net.ParseIP("10.20.30.40")
	oneEndingIP := net.ParseIP("10.20.30.50")

	twoStartingIP := net.ParseIP("10.20.30.35")
	twoEndingIP := net.ParseIP("10.20.30.45")

	ans := NetblocksOverlap(&oneStartingIP, &oneEndingIP, &twoStartingIP, &twoEndingIP)
	if !ans {
		t.Fail()
	}
}

func TestSecondContainsStartOfFirst(t *testing.T) {
	oneStartingIP := net.ParseIP("10.20.30.50")
	oneEndingIP := net.ParseIP("10.20.30.60")

	twoStartingIP := net.ParseIP("10.20.30.45")
	twoEndingIP := net.ParseIP("10.20.30.55")

	ans := NetblocksOverlap(&oneStartingIP, &oneEndingIP, &twoStartingIP, &twoEndingIP)
	if !ans {
		t.Fail()
	}
}

func TestSecondContainsEndOfFirst(t *testing.T) {
	oneStartingIP := net.ParseIP("10.20.30.40")
	oneEndingIP := net.ParseIP("10.20.30.50")

	twoStartingIP := net.ParseIP("10.20.30.45")
	twoEndingIP := net.ParseIP("10.20.30.55")

	ans := NetblocksOverlap(&oneStartingIP, &oneEndingIP, &twoStartingIP, &twoEndingIP)
	if !ans {
		t.Fail()
	}
}
