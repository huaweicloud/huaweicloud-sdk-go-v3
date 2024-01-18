package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VgwIkePolicy struct {

	// 加密算法
	EncryptionAlgorithm *string `json:"encryption_algorithm,omitempty"`

	// DH密钥组
	DhGroup *string `json:"dh_group,omitempty"`

	// 认证算法
	AuthenticationAlgorithm *string `json:"authentication_algorithm,omitempty"`

	// 表示SA的生存周期，当该生存周期超时后IKE SA将自动更新
	LifetimeSeconds *int32 `json:"lifetime_seconds,omitempty"`
}

func (o VgwIkePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VgwIkePolicy struct{}"
	}

	return strings.Join([]string{"VgwIkePolicy", string(data)}, " ")
}
