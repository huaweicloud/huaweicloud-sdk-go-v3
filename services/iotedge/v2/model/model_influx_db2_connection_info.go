package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InfluxDb2ConnectionInfo 创建外部推送通道请求结构体
type InfluxDb2ConnectionInfo struct {

	// 鉴权token
	Token string `json:"token"`

	// 客户端信任证书列表
	TrustCerts *interface{} `json:"trust_certs,omitempty"`

	// 客户端是否开启校验域名
	VerifyHostname *bool `json:"verify_hostname,omitempty"`
}

func (o InfluxDb2ConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InfluxDb2ConnectionInfo struct{}"
	}

	return strings.Join([]string{"InfluxDb2ConnectionInfo", string(data)}, " ")
}
