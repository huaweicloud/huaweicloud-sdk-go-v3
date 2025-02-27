package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CertificateResp 证书信息
type CertificateResp struct {

	// SSL证书id
	Id string `json:"id"`

	// SSL证书所在的项目ID
	TenantId string `json:"tenant_id"`

	// SSL证书的管理状态；暂不支持
	AdminStateUp bool `json:"admin_state_up"`

	// SSL证书的名称。
	Name string `json:"name"`

	// SSL证书的描述。
	Description string `json:"description"`

	// SSL证书的类型。分为服务器证书(server)和CA证书(client)。
	Type CertificateRespType `json:"type"`

	// 服务器证书所签域名。该字段仅type为server时有效。
	Domain string `json:"domain"`

	// 服务器证书的私钥。仅type为server时有效。type为server时必选。
	PrivateKey string `json:"private_key"`

	// 当type为server时，表示服务器证书的公钥；当type为client时，表示用于认证客户端证书的CA证书。
	Certificate string `json:"certificate"`

	// SSL证书的过期时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09
	ExpireTime string `json:"expire_time"`

	// SSL证书的创建时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09
	CreateTime string `json:"create_time"`

	// SSL证书的更新时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09
	UpdateTime string `json:"update_time"`

	// 参数解释： 证书来源  约束限制： 当scm_certificate_id不为空，且未传入source时，默认取值为“scm”。  取值范围： 无  默认取值： 当scm_certificate_id不为空，且未传入source时，默认取值为“scm”； 其他情况下默认为空。
	Source *string `json:"source,omitempty"`

	// 参数解释： 修改保护状态  约束限制： 无  取值范围：  - nonProtection: 不保护  - consoleProtection: 控制台修改保护  默认取值： nonProtection
	ProtectionStatus *CertificateRespProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释： 设置修改保护的原因 约束限制： 仅当protection_status为consoleProtection时有效 取值范围： 无 默认取值： 空
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o CertificateResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateResp struct{}"
	}

	return strings.Join([]string{"CertificateResp", string(data)}, " ")
}

type CertificateRespType struct {
	value string
}

type CertificateRespTypeEnum struct {
	SERVER CertificateRespType
	CLIENT CertificateRespType
}

func GetCertificateRespTypeEnum() CertificateRespTypeEnum {
	return CertificateRespTypeEnum{
		SERVER: CertificateRespType{
			value: "server",
		},
		CLIENT: CertificateRespType{
			value: "client",
		},
	}
}

func (c CertificateRespType) Value() string {
	return c.value
}

func (c CertificateRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateRespType) UnmarshalJSON(b []byte) error {
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

type CertificateRespProtectionStatus struct {
	value string
}

type CertificateRespProtectionStatusEnum struct {
	NON_PROTECTION     CertificateRespProtectionStatus
	CONSOLE_PROTECTION CertificateRespProtectionStatus
}

func GetCertificateRespProtectionStatusEnum() CertificateRespProtectionStatusEnum {
	return CertificateRespProtectionStatusEnum{
		NON_PROTECTION: CertificateRespProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CertificateRespProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CertificateRespProtectionStatus) Value() string {
	return c.value
}

func (c CertificateRespProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateRespProtectionStatus) UnmarshalJSON(b []byte) error {
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
