/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package jsonld

import (
	_ "embed" //nolint:gci // required for go:embed

	"github.com/hyperledger/aries-framework-go/pkg/doc/jsonld/context"
)

// nolint:gochecknoglobals // embedded test contexts
var (
	//go:embed contexts/credentials-examples_v1.jsonld
	credentialExamples []byte
	//go:embed contexts/odrl.jsonld
	odrl []byte
)

// Contexts returns JSON-LD contexts used in tests.
func Contexts() []context.Document {
	return []context.Document{
		{
			URL:     "https://www.w3.org/2018/credentials/examples/v1",
			Content: credentialExamples,
		},
		{
			URL:     "https://www.w3.org/ns/odrl.jsonld",
			Content: odrl,
		},
	}
}
