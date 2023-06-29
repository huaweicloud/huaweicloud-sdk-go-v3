package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVirtualInterfaceRequestBody Specifies parameters required for updating a virtual-interface.
type UpdateVirtualInterfaceRequestBody struct {
	VirtualInterface *UpdateVirtualInterface `json:"virtual_interface"`
}

func (o UpdateVirtualInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVirtualInterfaceRequestBody", string(data)}, " ")
}
