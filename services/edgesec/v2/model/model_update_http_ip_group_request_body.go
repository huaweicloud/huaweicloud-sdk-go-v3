package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpIpGroupRequestBody 更新IP地址组的请求
type UpdateHttpIpGroupRequestBody struct {

	// IP地址组名称
	Name string `json:"name"`

	// IP地址/地址段
	Ips string `json:"ips"`

	// IP地址组备注，最长512字符
	Description *string `json:"description,omitempty"`
}

func (o UpdateHttpIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpIpGroupRequestBody", string(data)}, " ")
}
