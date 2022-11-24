package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListVirtualInterfacesResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 虚拟接口对象
	VirtualInterfaces *[]VirtualInterface `json:"virtual_interfaces,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVirtualInterfacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirtualInterfacesResponse struct{}"
	}

	return strings.Join([]string{"ListVirtualInterfacesResponse", string(data)}, " ")
}
