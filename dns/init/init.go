// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: Apache-2.0

package init

import (
	"net"

	"github.com/johnstarich/go/dns"
)

func init() {
	net.DefaultResolver = dns.New()
}
