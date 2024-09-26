package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantTrustedIpAddressRequest Request Object
type UpdateTenantTrustedIpAddressRequest struct {

	// ip的ip
	IpId int32 `json:"ip_id"`

	Body *AddTrustedIpAddressRequestBody `json:"body,omitempty"`
}

func (o UpdateTenantTrustedIpAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantTrustedIpAddressRequest struct{}"
	}

	return strings.Join([]string{"UpdateTenantTrustedIpAddressRequest", string(data)}, " ")
}
