package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建虚拟接口请求参数
type CreateVirtualInterfaceRequestBody struct {
	VirtualInterface *CreateVirtualInterface `json:"virtual_interface"`
}

func (o CreateVirtualInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVirtualInterfaceRequestBody", string(data)}, " ")
}
