package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVirtualInterfaceRequest Request Object
type UpdateVirtualInterfaceRequest struct {

	// 虚拟接口ID。
	VirtualInterfaceId string `json:"virtual_interface_id"`

	Body *UpdateVirtualInterfaceRequestBody `json:"body,omitempty"`
}

func (o UpdateVirtualInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualInterfaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateVirtualInterfaceRequest", string(data)}, " ")
}
