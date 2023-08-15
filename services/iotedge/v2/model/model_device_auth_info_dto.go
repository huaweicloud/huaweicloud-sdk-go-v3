package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceAuthInfoDto 边缘节点设备认证方式信息。
type DeviceAuthInfoDto struct {

	// 边缘节点认证方式，不填默认为密钥认证接入方式(SECRET)。
	AuthType string `json:"auth_type"`

	// 证书指纹，认证类型使用证书认证接入(CERTIFICATES)需填写该字段。
	Fingerprint *string `json:"fingerprint,omitempty"`

	LocalPath *CertificateLocalPathDto `json:"local_path,omitempty"`
}

func (o DeviceAuthInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceAuthInfoDto struct{}"
	}

	return strings.Join([]string{"DeviceAuthInfoDto", string(data)}, " ")
}
