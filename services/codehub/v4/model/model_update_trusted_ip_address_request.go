package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTrustedIpAddressRequest Request Object
type UpdateTrustedIpAddressRequest struct {

	// **参数解释：** 仓库id，代码仓首页，Repository ID后的数字Id。
	Id int32 `json:"id"`

	// **参数解释：** ip白名单id。
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
