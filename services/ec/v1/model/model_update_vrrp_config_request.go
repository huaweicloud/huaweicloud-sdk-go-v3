package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVrrpConfigRequest Request Object
type UpdateVrrpConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 虚路由ID
	VirtualRouterId int32 `json:"virtual_router_id"`

	Body *CreateUpdateVrrpConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateVrrpConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVrrpConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateVrrpConfigRequest", string(data)}, " ")
}
