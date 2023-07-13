package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDnsServersResponse Response Object
type ListDnsServersResponse struct {

	// dns服务器列表
	Data *[]DnsServersResponseDto `json:"data,omitempty"`

	// dns服务器总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDnsServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDnsServersResponse struct{}"
	}

	return strings.Join([]string{"ListDnsServersResponse", string(data)}, " ")
}
