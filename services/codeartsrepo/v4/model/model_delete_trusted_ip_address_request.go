package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrustedIpAddressRequest Request Object
type DeleteTrustedIpAddressRequest struct {

	// **参数解释：** 仓库id，代码仓首页，Repository ID后的数字Id。
	Id int32 `json:"id"`

	// **参数解释：** ip白名单id。
	IpId int32 `json:"ip_id"`
}

func (o DeleteTrustedIpAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrustedIpAddressRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrustedIpAddressRequest", string(data)}, " ")
}
