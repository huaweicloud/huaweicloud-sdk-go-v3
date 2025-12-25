package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExtendAttributeRequest Request Object
type UpdateExtendAttributeRequest struct {

	// 虚拟接口ID。
	VirtualInterfaceId string `json:"virtual_interface_id"`

	Body *UpdateExtendAttributeRequestBody `json:"body,omitempty"`
}

func (o UpdateExtendAttributeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExtendAttributeRequest struct{}"
	}

	return strings.Join([]string{"UpdateExtendAttributeRequest", string(data)}, " ")
}
