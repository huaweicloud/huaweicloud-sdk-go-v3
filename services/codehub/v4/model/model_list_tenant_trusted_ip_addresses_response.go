package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantTrustedIpAddressesResponse Response Object
type ListTenantTrustedIpAddressesResponse struct {
	Body *[]TenantTrustedIpAddressDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTenantTrustedIpAddressesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantTrustedIpAddressesResponse struct{}"
	}

	return strings.Join([]string{"ListTenantTrustedIpAddressesResponse", string(data)}, " ")
}
