package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTrustedIpAddressRequest Request Object
type AddTrustedIpAddressRequest struct {

	// **参数解释：** 仓库id，代码仓首页，Repository ID后的数字Id。
	Id int32 `json:"id"`

	Body *AddTrustedIpAddressRequestBody `json:"body,omitempty"`
}

func (o AddTrustedIpAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTrustedIpAddressRequest struct{}"
	}

	return strings.Join([]string{"AddTrustedIpAddressRequest", string(data)}, " ")
}
