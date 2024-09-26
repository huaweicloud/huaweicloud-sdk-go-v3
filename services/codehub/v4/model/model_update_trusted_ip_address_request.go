package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTrustedIpAddressRequest Request Object
type UpdateTrustedIpAddressRequest struct {

	// 仓库id
	Id int32 `json:"id"`

	// ip的id
	IpId int32 `json:"ip_id"`

	Body *AddTrustedIpAddressRequestBody `json:"body,omitempty"`
}

func (o UpdateTrustedIpAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrustedIpAddressRequest struct{}"
	}

	return strings.Join([]string{"UpdateTrustedIpAddressRequest", string(data)}, " ")
}
