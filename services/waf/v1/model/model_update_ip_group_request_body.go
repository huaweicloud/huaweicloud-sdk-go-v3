package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 地址组名称
type UpdateIpGroupRequestBody struct {

	// 地址组名称
	Name string `json:"name"`

	// 以逗号分隔的ip或ip段
	Ips string `json:"ips"`

	// 地址组描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequestBody", string(data)}, " ")
}
