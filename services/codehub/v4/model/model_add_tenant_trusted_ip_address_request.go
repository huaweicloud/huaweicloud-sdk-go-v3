package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTenantTrustedIpAddressRequest Request Object
type AddTenantTrustedIpAddressRequest struct {
	Body *AddTrustedIpAddressRequestBody `json:"body,omitempty"`
}

func (o AddTenantTrustedIpAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTenantTrustedIpAddressRequest struct{}"
	}

	return strings.Join([]string{"AddTenantTrustedIpAddressRequest", string(data)}, " ")
}
