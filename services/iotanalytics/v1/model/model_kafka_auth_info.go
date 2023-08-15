package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KafkaAuthInfo Kafka 认证信息
type KafkaAuthInfo struct {

	// 安全协议
	SecurityProtocol string `json:"security_protocol"`

	SaslPlainAuthInfo *SaslPlainAuthInfo `json:"sasl_plain_auth_info,omitempty"`
}

func (o KafkaAuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaAuthInfo struct{}"
	}

	return strings.Join([]string{"KafkaAuthInfo", string(data)}, " ")
}
