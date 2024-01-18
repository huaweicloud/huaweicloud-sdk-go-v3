package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VgwIpsecPolicy struct {

	// 认证算法
	AuthenticationAlgorithm *string `json:"authentication_algorithm,omitempty"`

	// 加密算法
	EncryptionAlgorithm *string `json:"encryption_algorithm,omitempty"`

	// DH密钥组
	Pfs *string `json:"pfs,omitempty"`

	// 表示配置IPSec连接建立的隧道以时间为基准的生存周期
	LifetimeSeconds *int32 `json:"lifetime_seconds,omitempty"`
}

func (o VgwIpsecPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VgwIpsecPolicy struct{}"
	}

	return strings.Join([]string{"VgwIpsecPolicy", string(data)}, " ")
}
