package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMqttNodeChannelConnectionInfo 更新外部推送通道请求结构体
type UpdateMqttNodeChannelConnectionInfo struct {

	// mqtt协议中的ClientId
	ClientId *string `json:"client_id,omitempty"`

	// 鉴权用户名
	Username *string `json:"username,omitempty"`

	// 鉴权密码
	Password *string `json:"password,omitempty"`

	// 客户端信任证书列表
	TrustCerts *interface{} `json:"trust_certs,omitempty"`

	// 客户端是否开启校验域名
	VerifyHostname *bool `json:"verify_hostname,omitempty"`
}

func (o UpdateMqttNodeChannelConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMqttNodeChannelConnectionInfo struct{}"
	}

	return strings.Join([]string{"UpdateMqttNodeChannelConnectionInfo", string(data)}, " ")
}
