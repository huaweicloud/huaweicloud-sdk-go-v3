package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaCertificate 对端网关CA证书信息
type CaCertificate struct {

	// 证书序列号
	SerialNumber *string `json:"serial_number,omitempty"`

	// 签名算法
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 证书颁发者
	Issuer *string `json:"issuer,omitempty"`

	// 证书主题
	Subject *string `json:"subject,omitempty"`

	// 证书过期时间
	ExpireTime *sdktime.SdkTime `json:"expire_time,omitempty"`

	// 是否能更新内容
	IsUpdatable *bool `json:"is_updatable,omitempty"`
}

func (o CaCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaCertificate struct{}"
	}

	return strings.Join([]string{"CaCertificate", string(data)}, " ")
}
