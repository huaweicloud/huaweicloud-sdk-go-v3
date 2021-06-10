package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CacheConfigRequestBody struct {
	CacheConfig *CacheConfigRequest `json:"cache_config"`
}

func (o CacheConfigRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CacheConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CacheConfigRequestBody", string(data)}, " ")
}
