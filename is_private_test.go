package ipip

import (
	"net"
	"testing"
)

func TestIsPrivate(t *testing.T) {
	testCases := []struct {
		in  net.IP
		out bool
	}{
		{nil, false},
		{net.IPv4(10, 0, 0, 0), true},
		{net.IPv4(11, 0, 0, 0), false},
		{net.IPv4(172, 16, 0, 0), true},
		{net.IPv4(172, 32, 0, 0), false},
		{net.IPv4(192, 168, 0, 0), true},
		{net.IPv4(192, 169, 0, 0), false},
		{net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, true},
		{net.IP{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
	}
	for _, tc := range testCases {
		if private := IsPrivate(tc.in); private != tc.out {
			t.Errorf("IsPrivate(%s): got %v, want %v", tc.in, private, tc.out)
		}
	}
}
