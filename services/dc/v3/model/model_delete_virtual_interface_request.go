package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVirtualInterfaceRequest Request Object
type DeleteVirtualInterfaceRequest struct {

	// 虚拟接口ID。
	VirtualInterfaceId string `json:"virtual_interface_id"`
}

func (o DeleteVirtualInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualInterfaceRequest struct{}"
	}

	return strings.Join([]string{"DeleteVirtualInterfaceRequest", string(data)}, " ")
}
