/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package embed_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hyperledger/aries-framework-go/pkg/doc/jsonld/context/embed"
)

func TestProvider_Contexts(t *testing.T) {
	p := embed.Provider{}
	require.NotNil(t, p)

	contexts, err := p.Contexts()

	require.NoError(t, err)
	require.Equal(t, 16, len(contexts))
}
