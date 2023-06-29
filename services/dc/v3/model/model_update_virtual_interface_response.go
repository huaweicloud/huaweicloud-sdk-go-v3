package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVirtualInterfaceResponse Response Object
type UpdateVirtualInterfaceResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	VirtualInterface *VirtualInterface `json:"virtual_interface,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o UpdateVirtualInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualInterfaceResponse struct{}"
	}

	return strings.Join([]string{"UpdateVirtualInterfaceResponse", string(data)}, " ")
}
