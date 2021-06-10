package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowCacheRulesResponse struct {
	CacheConfig    *CacheConfig `json:"cache_config,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowCacheRulesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCacheRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowCacheRulesResponse", string(data)}, " ")
}
