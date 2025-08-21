package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrustedIpAddressesRequest Request Object
type ListTrustedIpAddressesRequest struct {

	// **参数解释：** 仓库id，代码仓首页，Repository ID后的数字Id。
	Id int32 `json:"id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTrustedIpAddressesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrustedIpAddressesRequest struct{}"
	}

	return strings.Join([]string{"ListTrustedIpAddressesRequest", string(data)}, " ")
}
