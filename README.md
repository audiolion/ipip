Package ipip (ip is private) checks if a `net.IP` is a private (local) address according to [RFC 1918](https://tools.ietf.org/html/rfc1918) for ip4 and [RFC 4193](https://tools.ietf.org/html/rfc4193) for ip6.

[![GoDoc](https://godoc.org/github.com/mikioh/ipaddr?status.png)](https://pkg.go.dev/github.com/audiolion/ipip)

**Install**

```shell
$ go get -u github.com/audiolion/ipip
```

**Usage**

```go

import (
  "github.com/audiolion/ipip"
)

func main() {

  var ip net.IP
  ip = net.IPv4(10, 0, 0, 0)
  
  ipip.IsPrivate(ip) // true
}
```

**Prior Art**

ipip is based on the work of [Max Semenik](https://github.com/MaxSem) with input from [Mikio Hara](https://github.com/mikioh) to [add this functionality](https://go-review.googlesource.com/c/go/+/162998/) directly in Go's `net` package. Track the status of this change request [here](https://github.com/golang/go/issues/29146).

**FAQ**

> What about multicast, limited-broadcast, unassigned, and reserved addresses?

The goal of ipip is to implement checking based on RFC 1918 and RFC 4193, and considers these edge cases to be out of scope.

> Why did you use IsPrivate instead of IsLocal?

RFC 1918 referred to these ip4 addresses as "private", meaning they are not routable on the public internet. When the same concept was applied to ip6 in RFC 4193 the term "local" address was used. While local is the more correct term, the layman probably is used to talking about "public" vs. "private" addresses and so the private terminology was used.

> There is no such thing as a local-untyped address, why not IsPrivateUnicast?

Yes, IsPrivateUnicast would be more technically correct, but the purpose of this package is to distinguish between "public" routable addresses and non-public (private) addresses as defined by RFC 1918 and RFC 4193. Adding "Unicast" makes the name more confusing to the layman, and we clarify unhandled edge cases in the comments and in this faq.

