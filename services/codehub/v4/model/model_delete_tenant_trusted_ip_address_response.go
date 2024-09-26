package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTenantTrustedIpAddressResponse Response Object
type DeleteTenantTrustedIpAddressResponse struct {

	// 状态码
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTenantTrustedIpAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantTrustedIpAddressResponse struct{}"
	}

	return strings.Join([]string{"DeleteTenantTrustedIpAddressResponse", string(data)}, " ")
}
