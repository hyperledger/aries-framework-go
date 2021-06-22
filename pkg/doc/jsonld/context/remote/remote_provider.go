/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package remote

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hyperledger/aries-framework-go/pkg/common/log"
	jsonldcontext "github.com/hyperledger/aries-framework-go/pkg/doc/jsonld/context"
)

const defaultTimeout = time.Minute

var logger = log.New("aries-framework/context/remote")

type options struct {
	httpClient HTTPClient
}

// ProviderOpt configures remote context provider.
type ProviderOpt func(*options)

// HTTPClient represents an HTTP client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// WithHTTPClient configures an HTTP client.
func WithHTTPClient(client HTTPClient) ProviderOpt {
	return func(opt *options) {
		opt.httpClient = client
	}
}

// Provider is a remote JSON-LD context provider.
type Provider struct {
	endpoint string
	http     HTTPClient
}

// NewProvider returns a new instance of remote Provider.
func NewProvider(endpoint string, opts ...ProviderOpt) (*Provider, error) {
	providerOpts := &options{httpClient: &http.Client{
		Timeout: defaultTimeout,
	}}

	for _, opt := range opts {
		opt(providerOpts)
	}

	return &Provider{
		endpoint: endpoint,
		http:     providerOpts.httpClient,
	}, nil
}

// Response represents expected response with JSON-LD contexts from the remote source.
type Response struct {
	Documents []jsonldcontext.Document `json:"documents"`
}

// Contexts returns JSON-LD contexts from the remote source.
func (p *Provider) Contexts() ([]jsonldcontext.Document, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, p.endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := p.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http do: %w", err)
	}

	defer func() {
		e := resp.Body.Close()
		if e != nil {
			logger.Errorf("Failed to close response body: %s", e.Error())
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}

	var response Response

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return response.Documents, nil
}
