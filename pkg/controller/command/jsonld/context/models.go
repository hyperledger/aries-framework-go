/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package context

import "github.com/hyperledger/aries-framework-go/pkg/doc/jsonld/context"

// AddRequest is a request model for adding JSON-LD contexts.
type AddRequest struct {
	Documents []context.Document `json:"documents"`
}
