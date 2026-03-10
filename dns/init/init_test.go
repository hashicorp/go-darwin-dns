// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: Apache-2.0

package init

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.NotNil(t, net.DefaultResolver)
}
