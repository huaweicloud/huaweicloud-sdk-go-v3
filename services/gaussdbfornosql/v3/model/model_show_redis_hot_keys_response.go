package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRedisHotKeysResponse Response Object
type ShowRedisHotKeysResponse struct {

	// 查询到的热Key列表。
	Keys *[]HotKeysInfoResponseBody `json:"keys,omitempty"`

	// 总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRedisHotKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisHotKeysResponse struct{}"
	}

	return strings.Join([]string{"ShowRedisHotKeysResponse", string(data)}, " ")
}
