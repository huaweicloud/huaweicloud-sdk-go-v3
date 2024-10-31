package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpIpGroupRequestBody 创建IP地址组的请求
type CreateHttpIpGroupRequestBody struct {

	// IP地址组名称
	Name string `json:"name"`

	// IP地址/IP段，以英文逗号分隔，最多配置200个IP/IP段
	Ips string `json:"ips"`

	// IP地址组备注，最长512字符
	Description *string `json:"description,omitempty"`
}

func (o CreateHttpIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHttpIpGroupRequestBody", string(data)}, " ")
}
