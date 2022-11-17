package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVirtualInterfaceRequestBody struct {

	// 操作请求ID
	RequestId string `json:"request_id"`

	VirtualInterface *CreateVirtualInterface `json:"virtual_interface"`
}

func (o CreateVirtualInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVirtualInterfaceRequestBody", string(data)}, " ")
}
