/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package remote_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hyperledger/aries-framework-go/pkg/doc/jsonld/context/remote"
	"github.com/hyperledger/aries-framework-go/pkg/internal/jsonldtest"
)

func TestNewProvider(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		p, err := remote.NewProvider("endpoint")

		require.NoError(t, err)
		require.NotNil(t, p)
	})
}

func TestProvider_Contexts(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		resp := remote.Response{
			Documents: jsonldtest.Contexts(),
		}

		respBytes, err := json.Marshal(resp)
		require.NoError(t, err)

		p, err := remote.NewProvider("endpoint", remote.WithHTTPClient(&mockHTTPClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader(respBytes)),
				}, nil
			},
		}))
		require.NoError(t, err)

		contexts, err := p.Contexts()
		require.NoError(t, err)
		require.NotNil(t, contexts)
		require.Equal(t, len(jsonldtest.Contexts()), len(contexts))
	})

	t.Run("Response error during HTTP do", func(t *testing.T) {
		p, err := remote.NewProvider("endpoint", remote.WithHTTPClient(&mockHTTPClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("response error")
			},
		}))
		require.NoError(t, err)

		contexts, err := p.Contexts()
		require.Empty(t, contexts)
		require.Error(t, err)
		require.Contains(t, err.Error(), "http do: response error")
	})

	t.Run("Response status code not 200 OK", func(t *testing.T) {
		p, err := remote.NewProvider("endpoint", remote.WithHTTPClient(&mockHTTPClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				}, nil
			},
		}))
		require.NoError(t, err)

		contexts, err := p.Contexts()
		require.Empty(t, contexts)
		require.Error(t, err)
		require.Contains(t, err.Error(), "response status code: 500")
	})

	t.Run("Fail to decode response", func(t *testing.T) {
		p, err := remote.NewProvider("endpoint", remote.WithHTTPClient(&mockHTTPClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("invalid response"))),
				}, nil
			},
		}))
		require.NoError(t, err)

		contexts, err := p.Contexts()
		require.Empty(t, contexts)
		require.Error(t, err)
		require.Contains(t, err.Error(), "decode response:")
	})
}

type mockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
