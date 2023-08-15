package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Address 服务器开发端口信息
type Address struct {

	// 云手机服务器的内网IP，过期字段
	IntranetIp *string `json:"intranet_ip,omitempty"`

	// 云手机服务器的公网IP，过期字段
	AccessIp *string `json:"access_ip,omitempty"`

	// 云手机服务器的内网IP，新增字段
	ServerIp *string `json:"server_ip,omitempty"`

	// 云手机服务器的公网IP，新增字段
	PublicIp *string `json:"public_ip,omitempty"`
}

func (o Address) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Address struct{}"
	}

	return strings.Join([]string{"Address", string(data)}, " ")
}
