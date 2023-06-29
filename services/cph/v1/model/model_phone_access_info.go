package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhoneAccessInfo 云手机访问信息。
type PhoneAccessInfo struct {

	// 自定义端口类型，不超过16个字节
	Type *string `json:"type,omitempty"`

	// 云手机IP（过期）
	DeviceIp *string `json:"device_ip,omitempty"`

	// 云手机IP
	PhoneIp *string `json:"phone_ip,omitempty"`

	// 服务监听端口
	ListenPort *int32 `json:"listen_port,omitempty"`

	// 云手机服务器的访问IP（过期）
	AccessIp *string `json:"access_ip,omitempty"`

	// 云手机服务器的公网IP，如果端口设置了非公网访问，该字段返回空字符串
	PublicIp *string `json:"public_ip,omitempty"`

	// 云手机服务器的内网IP（过期）
	IntranetIp *string `json:"intranet_ip,omitempty"`

	// 云手机服务器的内网IP
	ServerIp *string `json:"server_ip,omitempty"`

	// 服务映射到公网的访问端口
	AccessPort *int32 `json:"access_port,omitempty"`
}

func (o PhoneAccessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneAccessInfo struct{}"
	}

	return strings.Join([]string{"PhoneAccessInfo", string(data)}, " ")
}
