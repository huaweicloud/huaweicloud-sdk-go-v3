package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ItMqttConnectionInfo 创建外部推送通道请求结构体
type ItMqttConnectionInfo struct {

	// 鉴权用户名
	Username string `json:"username"`

	// 鉴权密码
	Password string `json:"password"`

	// 客户端信任证书列表
	TrustCerts *interface{} `json:"trust_certs,omitempty"`

	// 客户端是否开启校验域名
	VerifyHostname *bool `json:"verify_hostname,omitempty"`
}

func (o ItMqttConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItMqttConnectionInfo struct{}"
	}

	return strings.Join([]string{"ItMqttConnectionInfo", string(data)}, " ")
}
