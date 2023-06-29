package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVirtualInterfaceResponse Response Object
type CreateVirtualInterfaceResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	VirtualInterface *VirtualInterface `json:"virtual_interface,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o CreateVirtualInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualInterfaceResponse struct{}"
	}

	return strings.Join([]string{"CreateVirtualInterfaceResponse", string(data)}, " ")
}
