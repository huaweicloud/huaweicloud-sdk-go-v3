package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateCertificateOption struct {

	// 证书的内容。PEM编码格式。 最大长度65536字符。 支持证书链，最大11层(含证书和证书链)。
	Certificate *string `json:"certificate,omitempty"`

	// 证书的描述。
	Description *string `json:"description,omitempty"`

	// 证书的名称。
	Name *string `json:"name,omitempty"`

	// 服务器证书的私钥。PEM编码格式。 当type为client时，该参数被忽略，不影响证书的创建和使用。若传入不符合格式值，则会报错。 当type为server时，该字段必须符合格式要求，且私钥必须是有效的。 最大长度8192字符。
	PrivateKey *string `json:"private_key,omitempty"`

	// 服务器证书所签域名。该字段仅type为server时有效。  总长度为0-10000，由若干普通域名或泛域名组成，域名之间以\",\"分隔，不超过100个域名。  普通域名：由若干字符串组成，字符串间以\".\"分隔，单个字符串长度不超过63个字符， 只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。例：www.test.com；  泛域名：在普通域名的基础上仅允许首字母为\"\\*\"。例：\\*.test.com
	Domain *string `json:"domain,omitempty"`

	// HTTPS协议使用的SM加密证书内容。支持证书链，最大11层(含证书和证书链)。  取值：PEM编码格式。最大长度65536字符。  使用说明：仅type为server_sm时有效。
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// HTTPS协议使用的SM加密证书内容。  取值：PEM编码格式。最大长度8192字符。  使用说明：仅type为server_sm时有效。
	EncPrivateKey *string `json:"enc_private_key,omitempty"`

	// scm证书id
	ScmCertificateId *string `json:"scm_certificate_id,omitempty"`

	// 参数解释：证书来源 取值范围：无  默认取值：当scm_certificate_id不为空，且未传入source时，默认取值为“scm”； 其他情况下默认为空。
	Source *string `json:"source,omitempty"`

	// 参数解释：修改保护状态  约束限制：无  取值范围：  - nonProtection: 不保护  - consoleProtection: 控制台修改保护  默认取值：nonProtection
	ProtectionStatus *UpdateCertificateOptionProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释：设置修改保护的原因  约束限制：仅当protection_status为consoleProtection时有效  取值范围：无  默认取值：空
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o UpdateCertificateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateOption struct{}"
	}

	return strings.Join([]string{"UpdateCertificateOption", string(data)}, " ")
}

type UpdateCertificateOptionProtectionStatus struct {
	value string
}

type UpdateCertificateOptionProtectionStatusEnum struct {
	NON_PROTECTION     UpdateCertificateOptionProtectionStatus
	CONSOLE_PROTECTION UpdateCertificateOptionProtectionStatus
}

func GetUpdateCertificateOptionProtectionStatusEnum() UpdateCertificateOptionProtectionStatusEnum {
	return UpdateCertificateOptionProtectionStatusEnum{
		NON_PROTECTION: UpdateCertificateOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdateCertificateOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdateCertificateOptionProtectionStatus) Value() string {
	return c.value
}

func (c UpdateCertificateOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCertificateOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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
