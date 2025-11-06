package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTenantTrustedIpAddressRequest Request Object
type DeleteTenantTrustedIpAddressRequest struct {

	// **参数解释：** ip白名单id。
	IpId int32 `json:"ip_id"`
}

func (o DeleteTenantTrustedIpAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantTrustedIpAddressRequest struct{}"
	}

	return strings.Join([]string{"DeleteTenantTrustedIpAddressRequest", string(data)}, " ")
}
