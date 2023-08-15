package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVirtualInterfaceResponse Response Object
type ShowVirtualInterfaceResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	VirtualInterface *VirtualInterface `json:"virtual_interface,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o ShowVirtualInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirtualInterfaceResponse struct{}"
	}

	return strings.Join([]string{"ShowVirtualInterfaceResponse", string(data)}, " ")
}
