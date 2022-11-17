package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVirtualInterfaceRequest struct {

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// 虚拟接口ID。
	VirtualInterfaceId string `json:"virtual_interface_id"`
}

func (o ShowVirtualInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirtualInterfaceRequest struct{}"
	}

	return strings.Join([]string{"ShowVirtualInterfaceRequest", string(data)}, " ")
}
