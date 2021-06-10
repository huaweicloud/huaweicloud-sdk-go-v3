package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateCacheRulesResponse struct {
	CacheConfig    *CacheConfig `json:"cache_config,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateCacheRulesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCacheRulesResponse struct{}"
	}

	return strings.Join([]string{"UpdateCacheRulesResponse", string(data)}, " ")
}
