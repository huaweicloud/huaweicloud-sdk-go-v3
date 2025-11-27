package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListByoipPoolsResponse Response Object
type ListByoipPoolsResponse struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	// 自带IP地址池列表。
	ByoipPools *[]ByoipPool `json:"byoip_pools,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListByoipPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListByoipPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListByoipPoolsResponse", string(data)}, " ")
}
