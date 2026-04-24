package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDNatInfoResponse Response Object
type ListDNatInfoResponse struct {

	// **参数解释**: 查询实例已绑定的NAT网关列表。
	DnatList       *[]DNatInfoResult `json:"dnat_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListDNatInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDNatInfoResponse struct{}"
	}

	return strings.Join([]string{"ListDNatInfoResponse", string(data)}, " ")
}
