package ipip

import "net"

// IsPrivate reports whether ip is a private (local) address.
//
// The identification of private, or local, unicast addresses uses address type
// indentification as defined in RFC 1918 for ip4 and RFC 4193 for ip6 with the
// exception of ip4 directed broadcast addresses.
//
// Unassigned, reserved, multicast and limited-broadcast addresses are not
// handled and will return false.
func IsPrivate(ip net.IP) bool {
	if ip4 := ip.To4(); ip4 != nil {
		return ip4[0] == 10 || (ip4[0] == 172 && ip4[1]&0xf0 == 16) || (ip4[0] == 192 && ip4[1] == 168)
	}
	return len(ip) == net.IPv6len && ip[0]&0xfe == 0xfc
}
