package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RadiusHost Radius主机配置信息
type RadiusHost struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// IP地址
	RadiusIp *string `json:"radius_ip,omitempty"`

	// 认证端口
	AuthPort *int32 `json:"auth_port,omitempty"`

	// 接收端口
	AcceptPort *int32 `json:"accept_port,omitempty"`

	// NAS ID
	NasId *string `json:"nas_id,omitempty"`
}

func (o RadiusHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RadiusHost struct{}"
	}

	return strings.Join([]string{"RadiusHost", string(data)}, " ")
}
