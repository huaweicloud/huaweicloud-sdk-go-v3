package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantTrustedIpAddressesRequest Request Object
type ListTenantTrustedIpAddressesRequest struct {

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 返回数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTenantTrustedIpAddressesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantTrustedIpAddressesRequest struct{}"
	}

	return strings.Join([]string{"ListTenantTrustedIpAddressesRequest", string(data)}, " ")
}
