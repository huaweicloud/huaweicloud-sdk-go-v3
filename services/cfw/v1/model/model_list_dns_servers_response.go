package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDnsServersResponse Response Object
type ListDnsServersResponse struct {

	// **参数解释**： dns服务器列表 **取值范围**： 不涉及
	Data *[]DnsServersResponseDto `json:"data,omitempty"`

	// **参数解释**： dns服务器总数 **取值范围**： 不涉及
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
