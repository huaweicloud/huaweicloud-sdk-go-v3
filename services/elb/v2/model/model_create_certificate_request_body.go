package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateCertificateRequestBody This is a auto create Body Object
type CreateCertificateRequestBody struct {

	// 服务端公有密钥证书或者用于认证客户端证书的CA证书，由type字段区分。 格式：证书为PEM格式。
	Certificate string `json:"certificate"`

	// 服务端的私有密钥。 格式：私钥为PEM格式。 该字段仅type为server时有效且为必选。 该字段在type为client时无效。
	PrivateKey *string `json:"private_key,omitempty"`

	// SSL证书的描述信息。支持的最大字符长度：255
	Description *string `json:"description,omitempty"`

	// 服务端证书所签的域名。  取值：总长度为0-1024。  普通域名由若干字符串组成，字符串间以\".\"分割，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。  泛域名仅允许首段为\"*\"，其他取值限制与普通域名一致。如：*.domain.com，但不能是：*my.domain.com  该字段仅type为server时有效。
	Domain *string `json:"domain,omitempty"`

	// SSL证书的名称。支持的最大字符长度：255
	Name *string `json:"name,omitempty"`

	// SSL证书的管理状态； 取值范围： true/false。 该字段为预留字段，暂未启用。只支持设定为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// SSL证书的类型。默认值：server； 取值范围： server：服务端证书； client：客户端证书；
	Type *string `json:"type,omitempty"`

	// 企业项目ID。  传入all_granted_eps表示查询所有有权限的企业项目资源；\"0\"表示查询默认企业项目资源；或者指定的企业项目ID下的资源。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 参数解释： 证书来源  约束限制： 当scm_certificate_id不为空，且未传入source时，默认取值为“scm”。  取值范围： 无  默认取值： 当scm_certificate_id不为空，且未传入source时，默认取值为“scm”； 其他情况下默认为空。
	Source *string `json:"source,omitempty"`

	// 参数解释： 修改保护状态  约束限制： 无  取值范围：  - nonProtection: 不保护 - consoleProtection: 控制台修改保护  默认取值： nonProtection
	ProtectionStatus *CreateCertificateRequestBodyProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释： 设置修改保护的原因  约束限制： 仅当protection_status为consoleProtection时有效  取值范围： 无  默认取值： 空
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o CreateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequestBody", string(data)}, " ")
}

type CreateCertificateRequestBodyProtectionStatus struct {
	value string
}

type CreateCertificateRequestBodyProtectionStatusEnum struct {
	NON_PROTECTION     CreateCertificateRequestBodyProtectionStatus
	CONSOLE_PROTECTION CreateCertificateRequestBodyProtectionStatus
}

func GetCreateCertificateRequestBodyProtectionStatusEnum() CreateCertificateRequestBodyProtectionStatusEnum {
	return CreateCertificateRequestBodyProtectionStatusEnum{
		NON_PROTECTION: CreateCertificateRequestBodyProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreateCertificateRequestBodyProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreateCertificateRequestBodyProtectionStatus) Value() string {
	return c.value
}

func (c CreateCertificateRequestBodyProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCertificateRequestBodyProtectionStatus) UnmarshalJSON(b []byte) error {
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
