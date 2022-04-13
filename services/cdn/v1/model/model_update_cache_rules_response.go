package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateCacheRulesResponse struct {
	CacheConfig    *CacheConfig `json:"cache_config,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateCacheRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCacheRulesResponse struct{}"
	}

	return strings.Join([]string{"UpdateCacheRulesResponse", string(data)}, " ")
}
