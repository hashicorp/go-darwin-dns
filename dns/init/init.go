package init

import (
	"net"

	"github.com/hashicorp/go-darwin-dns/dns"
)

func init() {
	net.DefaultResolver = dns.New()
}
