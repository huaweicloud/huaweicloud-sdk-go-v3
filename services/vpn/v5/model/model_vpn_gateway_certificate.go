package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type VpnGatewayCertificate struct {

	// VPN网关证书ID
	Id *string `json:"id,omitempty"`

	// VPN网关证书名称
	Name *string `json:"name,omitempty"`

	// 租户的项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// VPN网关ID
	VgwId *string `json:"vgw_id,omitempty"`

	// 网关证书状态
	Status *VpnGatewayCertificateStatus `json:"status,omitempty"`

	// 证书颁发者，国密证书时为签名证书颁发者
	Issuer *string `json:"issuer,omitempty"`

	// 证书签名算法，国密证书时为签名证书签名算法
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 证书序列号，国密证书时为签名证书序列号
	CertificateSerialNumber *string `json:"certificate_serial_number,omitempty"`

	// 证书主题，国密证书时为签名证书主题
	CertificateSubject *string `json:"certificate_subject,omitempty"`

	// 证书过期时间，国密证书时为签名证书过期时间
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

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o VpnGatewayCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnGatewayCertificate struct{}"
	}

	return strings.Join([]string{"VpnGatewayCertificate", string(data)}, " ")
}

type VpnGatewayCertificateStatus struct {
	value string
}

type VpnGatewayCertificateStatusEnum struct {
	BOUND   VpnGatewayCertificateStatus
	FAULT   VpnGatewayCertificateStatus
	BINDING VpnGatewayCertificateStatus
}

func GetVpnGatewayCertificateStatusEnum() VpnGatewayCertificateStatusEnum {
	return VpnGatewayCertificateStatusEnum{
		BOUND: VpnGatewayCertificateStatus{
			value: "BOUND",
		},
		FAULT: VpnGatewayCertificateStatus{
			value: "FAULT",
		},
		BINDING: VpnGatewayCertificateStatus{
			value: "BINDING",
		},
	}
}

func (c VpnGatewayCertificateStatus) Value() string {
	return c.value
}

func (c VpnGatewayCertificateStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpnGatewayCertificateStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
