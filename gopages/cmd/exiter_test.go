// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExiter(t *testing.T) {
	SetupTestExiter(t)
	assert.PanicsWithError(t, "Attempted to exit with exit code 1", func() {
		Exit(1)
	})
}
