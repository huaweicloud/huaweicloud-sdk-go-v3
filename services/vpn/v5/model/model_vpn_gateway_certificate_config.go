package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpnGatewayCertificateConfig struct {

	// VPN网关证书ID
	Id *string `json:"id,omitempty"`

	// VPN网关证书名称
	Name *string `json:"name,omitempty"`

	// VPN网关ID
	VgwId *string `json:"vgw_id,omitempty"`

	// 签名证书颁发者，国密证书时为签名证书颁发者
	Issuer *string `json:"issuer,omitempty"`

	// 签名证书签名算法，国密证书时为签名证书签名算法
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 证书序列号，国密证书时为签名证书序列号
	CertificateSerialNumber *string `json:"certificate_serial_number,omitempty"`

	// 签名证书主题，国密证书时为签名证书主题
	CertificateSubject *string `json:"certificate_subject,omitempty"`

	// 签名证书过期时间，国密证书时为签名证书过期时间
	CertificateExpireTime *sdktime.SdkTime `json:"certificate_expire_time,omitempty"`

	// CA证书序列号
	CertificateChainSerialNumber *string `json:"certificate_chain_serial_number,omitempty"`

	// CA证书主题
	CertificateChainSubject *string `json:"certificate_chain_subject,omitempty"`

	// CA证书过期时间
	CertificateChainExpireTime *sdktime.SdkTime `json:"certificate_chain_expire_time,omitempty"`

	// 国密证书的加密证书序列号，
	EncCertificateSerialNumber *string `json:"enc_certificate_serial_number,omitempty"`

	// 国密证书的加密证书主题
	EncCertificateSubject *string `json:"enc_certificate_subject,omitempty"`

	// 国密证书的加密证书过期时间
	EncCertificateExpireTime *sdktime.SdkTime `json:"enc_certificate_expire_time,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o VpnGatewayCertificateConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnGatewayCertificateConfig struct{}"
	}

	return strings.Join([]string{"VpnGatewayCertificateConfig", string(data)}, " ")
}
