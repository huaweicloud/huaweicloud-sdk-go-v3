package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRedisBigKeysResponse Response Object
type ShowRedisBigKeysResponse struct {

	// 查询到的大Key列表。
	Keys *[]BigKeysInfoResponseBody `json:"keys,omitempty"`

	// 大Key的总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRedisBigKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisBigKeysResponse struct{}"
	}

	return strings.Join([]string{"ShowRedisBigKeysResponse", string(data)}, " ")
}
