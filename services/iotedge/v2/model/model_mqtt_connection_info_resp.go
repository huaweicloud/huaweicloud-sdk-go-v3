package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MqttConnectionInfoResp 外部推送通道返回详情
type MqttConnectionInfoResp struct {

	// 鉴权用户名
	Username *string `json:"username,omitempty"`

	// 客户端信任证书列表
	TrustCerts *interface{} `json:"trust_certs,omitempty"`

	// 客户端是否开启校验域名
	VerifyHostname *bool `json:"verify_hostname,omitempty"`
}

func (o MqttConnectionInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MqttConnectionInfoResp struct{}"
	}

	return strings.Join([]string{"MqttConnectionInfoResp", string(data)}, " ")
}
