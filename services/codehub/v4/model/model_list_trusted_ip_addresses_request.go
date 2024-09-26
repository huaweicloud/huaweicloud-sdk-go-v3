package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrustedIpAddressesRequest Request Object
type ListTrustedIpAddressesRequest struct {

	// 仓库id
	Id int32 `json:"id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 返回数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTrustedIpAddressesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrustedIpAddressesRequest struct{}"
	}

	return strings.Join([]string{"ListTrustedIpAddressesRequest", string(data)}, " ")
}
