package rfc

import (
	"github.com/darkweak/souin/cache/types"
	"github.com/darkweak/souin/configurationtypes"
	"net/http"
	"net/http/httputil"
	"time"
)

// VaryTransport type
type VaryTransport types.Transport

// IsVaryCacheable determines if it's cacheable
func IsVaryCacheable(req *http.Request) bool {
	method := req.Method
	rangeHeader := req.Header.Get("range")
	return (method == http.MethodGet || method == http.MethodHead) && rangeHeader == ""
}

// NewTransport returns a new Transport with the
// provided Cache implementation and MarkCachedResponses set to true
func NewTransport(p types.AbstractProviderInterface) *VaryTransport {
	return &VaryTransport{
		Provider:               p,
		VaryLayerStorage:       types.InitializeVaryLayerStorage(),
		CoalescingLayerStorage: types.InitializeCoalescingLayerStorage(),
		MarkCachedResponses:    true,
	}
}

// GetProvider returns the associated provider
func (t *VaryTransport) GetProvider() types.AbstractProviderInterface {
	return t.Provider
}

// SetURL set the URL
func (t *VaryTransport) SetURL(url configurationtypes.URL) {
	t.ConfigurationURL = url
}

// GetVaryLayerStorage get the vary layer storagecache/coalescing/requestCoalescing_test.go
func (t *VaryTransport) GetVaryLayerStorage() *types.VaryLayerStorage {
	return t.VaryLayerStorage
}

// GetCoalescingLayerStorage get the coalescing layer storage
func (t *VaryTransport) GetCoalescingLayerStorage() *types.CoalescingLayerStorage {
	return t.CoalescingLayerStorage
}

// SetCache set the cache
func (t *VaryTransport) SetCache(key string, resp *http.Response) {
	if respBytes, err := httputil.DumpResponse(resp, true); err == nil {
		t.Provider.Set(key, respBytes, t.ConfigurationURL, time.Duration(0))
	}
}
