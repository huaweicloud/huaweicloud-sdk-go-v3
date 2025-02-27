package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateCertificateResponse Response Object
type UpdateCertificateResponse struct {

	// SSL证书id
	Id *string `json:"id,omitempty"`

	// SSL证书所在的项目ID
	TenantId *string `json:"tenant_id,omitempty"`

	// SSL证书的管理状态；暂不支持
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// SSL证书的名称。
	Name *string `json:"name,omitempty"`

	// SSL证书的描述。
	Description *string `json:"description,omitempty"`

	// SSL证书的类型。分为服务器证书(server)和CA证书(client)。
	Type *UpdateCertificateResponseType `json:"type,omitempty"`

	// 服务器证书所签域名。该字段仅type为server时有效。
	Domain *string `json:"domain,omitempty"`

	// 服务器证书的私钥。仅type为server时有效。type为server时必选。
	PrivateKey *string `json:"private_key,omitempty"`

	// 当type为server时，表示服务器证书的公钥；当type为client时，表示用于认证客户端证书的CA证书。
	Certificate *string `json:"certificate,omitempty"`

	// SSL证书的过期时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09
	ExpireTime *string `json:"expire_time,omitempty"`

	// SSL证书的创建时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09
	CreateTime *string `json:"create_time,omitempty"`

	// SSL证书的更新时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09
	UpdateTime *string `json:"update_time,omitempty"`

	// 参数解释： 证书来源  约束限制： 当scm_certificate_id不为空，且未传入source时，默认取值为“scm”。  取值范围： 无  默认取值： 当scm_certificate_id不为空，且未传入source时，默认取值为“scm”； 其他情况下默认为空。
	Source *string `json:"source,omitempty"`

	// 参数解释： 修改保护状态  约束限制： 无  取值范围：  - nonProtection: 不保护  - consoleProtection: 控制台修改保护  默认取值： nonProtection
	ProtectionStatus *UpdateCertificateResponseProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释： 设置修改保护的原因 约束限制： 仅当protection_status为consoleProtection时有效 取值范围： 无 默认取值： 空
	ProtectionReason *string `json:"protection_reason,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o UpdateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateResponse struct{}"
	}

	return strings.Join([]string{"UpdateCertificateResponse", string(data)}, " ")
}

type UpdateCertificateResponseType struct {
	value string
}

type UpdateCertificateResponseTypeEnum struct {
	SERVER UpdateCertificateResponseType
	CLIENT UpdateCertificateResponseType
}

func GetUpdateCertificateResponseTypeEnum() UpdateCertificateResponseTypeEnum {
	return UpdateCertificateResponseTypeEnum{
		SERVER: UpdateCertificateResponseType{
			value: "server",
		},
		CLIENT: UpdateCertificateResponseType{
			value: "client",
		},
	}
}

func (c UpdateCertificateResponseType) Value() string {
	return c.value
}

func (c UpdateCertificateResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCertificateResponseType) UnmarshalJSON(b []byte) error {
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

type UpdateCertificateResponseProtectionStatus struct {
	value string
}

type UpdateCertificateResponseProtectionStatusEnum struct {
	NON_PROTECTION     UpdateCertificateResponseProtectionStatus
	CONSOLE_PROTECTION UpdateCertificateResponseProtectionStatus
}

func GetUpdateCertificateResponseProtectionStatusEnum() UpdateCertificateResponseProtectionStatusEnum {
	return UpdateCertificateResponseProtectionStatusEnum{
		NON_PROTECTION: UpdateCertificateResponseProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdateCertificateResponseProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdateCertificateResponseProtectionStatus) Value() string {
	return c.value
}

func (c UpdateCertificateResponseProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCertificateResponseProtectionStatus) UnmarshalJSON(b []byte) error {
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
