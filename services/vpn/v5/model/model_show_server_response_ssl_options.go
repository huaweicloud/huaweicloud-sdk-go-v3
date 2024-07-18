package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerResponseSslOptions SSL隧道协议的可选配置项
type ShowServerResponseSslOptions struct {

	// 协议
	Protocol *string `json:"protocol,omitempty"`

	// 端口
	Port *int32 `json:"port,omitempty"`

	// 加密算法
	EncryptionAlgorithm *string `json:"encryption_algorithm,omitempty"`

	// 认证算法
	AuthenticationAlgorithm *string `json:"authentication_algorithm,omitempty"`

	// 是否压缩
	IsCompressed *bool `json:"is_compressed,omitempty"`
}

func (o ShowServerResponseSslOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerResponseSslOptions struct{}"
	}

	return strings.Join([]string{"ShowServerResponseSslOptions", string(data)}, " ")
}
