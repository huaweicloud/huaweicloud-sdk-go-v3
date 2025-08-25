package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceDistribution struct {

	// 加解密服务实例数量
	EncryptDecrypt int32 `json:"encrypt_decrypt"`

	// 签名验签服务实例数量
	SignVerify int32 `json:"sign_verify"`

	// 密钥管理服务实例数量
	Kms int32 `json:"kms"`

	// 时间戳服务实例数量
	Timestamp int32 `json:"timestamp"`

	// 协同签名服务实例数量
	CollaSign int32 `json:"colla_sign"`

	// 动态口令服务实例数量
	Otp int32 `json:"otp"`

	// 数据库加密服务实例数量
	DbEncrypt int32 `json:"db_encrypt"`

	// 文件加密服务实例数量
	FileEncrypt int32 `json:"file_encrypt"`

	// 电子签章服务实例数量
	DigitSeal int32 `json:"digit_seal"`

	// ssl vpn服务实例数量
	SslVpn int32 `json:"ssl_vpn"`
}

func (o InstanceDistribution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDistribution struct{}"
	}

	return strings.Join([]string{"InstanceDistribution", string(data)}, " ")
}
