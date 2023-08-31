package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVrrpConfigRequest Request Object
type DeleteVrrpConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 虚路由ID
	VirtualRouterId int32 `json:"virtual_router_id"`
}

func (o DeleteVrrpConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVrrpConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteVrrpConfigRequest", string(data)}, " ")
}
