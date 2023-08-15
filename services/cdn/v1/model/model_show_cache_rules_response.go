package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCacheRulesResponse Response Object
type ShowCacheRulesResponse struct {
	CacheConfig    *CacheConfig `json:"cache_config,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowCacheRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCacheRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowCacheRulesResponse", string(data)}, " ")
}
