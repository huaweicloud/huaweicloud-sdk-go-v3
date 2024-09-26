package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrustedIpAddressRequest Request Object
type DeleteTrustedIpAddressRequest struct {

	// ip的id
	Id int32 `json:"id"`

	// ip的id
	IpId int32 `json:"ip_id"`
}

func (o DeleteTrustedIpAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrustedIpAddressRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrustedIpAddressRequest", string(data)}, " ")
}
