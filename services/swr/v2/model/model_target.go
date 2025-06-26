package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Target struct {

	// 触发器类型，可选http
	Type string `json:"type"`

	// 触发地址，不可修改
	Address string `json:"address"`

	// 请求头，格式为key1:value1;key2:value2
	AuthHeader *string `json:"auth_header,omitempty"`

	// 是否跳过证书认证
	SkipCertVerify *bool `json:"skip_cert_verify,omitempty"`

	// 触发地址类型，可选internal(内网)，internet(公网)。internal必须为内网ip形式。
	AddressType string `json:"address_type"`
}

func (o Target) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Target struct{}"
	}

	return strings.Join([]string{"Target", string(data)}, " ")
}
