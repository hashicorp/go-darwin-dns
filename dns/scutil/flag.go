// Copyright IBM Corp. 2019, 2020
// SPDX-License-Identifier: Apache-2.0

package scutil

type Flag string

const (
	// Scoped requires a resolver to only send queries on the specified interface
	Scoped Flag = "Scoped"

	RequestARecords Flag = "Request A records"
)
