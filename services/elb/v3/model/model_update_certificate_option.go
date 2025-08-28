package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateCertificateOption struct {

	// **参数解释**：证书内容。 支持最大11层证书链(含证书和证书链)。  **约束限制**：不涉及  **取值范围**：PEM编码格式，最大长度65536个字符。  **默认取值**：不涉及
	Certificate *string `json:"certificate,omitempty"`

	// **参数解释**：证书的描述。  **约束限制**：不涉及  **取值范围**：0-255个字符。  **默认取值**：不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：证书的名称。  **约束限制**：不涉及  **取值范围**：0-255个字符。  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：服务器证书的私钥。  **约束限制**：不涉及  **取值范围**：PEM编码格式，最大长度8192个字符。  **默认取值**：不涉及
	PrivateKey *string `json:"private_key,omitempty"`

	// **参数解释**：服务器证书所签域名。  **约束限制**：该字段仅type为server时有效（其他类型证书，字段可传入，但不会生效）。  **取值范围**：总长度为0-10000，由若干普通域名或泛域名组成，域名之间以\",\"分隔，不超过100个域名。 - 普通域名：由若干字符串组成，字符串间以\".\"分隔，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。例：www.test.com。 - 泛域名：在普通域名的基础上仅允许首字母为\"\\*\"。例：\\*.test.com。  **默认取值**：不涉及
	Domain *string `json:"domain,omitempty"`

	// **参数解释**：服务器SM双证书的证书内容。 支持最大11层证书链(含证书和证书链)。  **约束限制**：仅当type为server_sm时，才支持传入。  **取值范围**：PEM编码格式。最大长度65536字符。  **默认取值**：不涉及
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// **参数解释**：服务器SM双证书的私钥。  **约束限制**：仅当type为server_sm时，才支持传入。  **取值范围**：PEM编码格式，最大长度8192个字符。  **默认取值**：不涉及
	EncPrivateKey *string `json:"enc_private_key,omitempty"`

	// **参数解释**：云证书管理服务（CCM）中的证书ID。  **约束限制**：仅记录证书ID，不验证其是否真实存在云证书管理服务中。并且需要将云证书管理服务中对应证书的内容手动设置到当前接口相应字段中（可能涉及字段certificate、private_key、enc_certificate和enc_private_key）  **取值范围**：不涉及  **默认取值**：不涉及
	ScmCertificateId *string `json:"scm_certificate_id,omitempty"`

	// **参数解释**：标记当前证书来源。  **约束限制**：无  **取值范围**： - scm：表示关联云证书管理服务（CCM）中的证书。 - 空值：表示自有证书。  **默认取值**：不涉及
	Source *string `json:"source,omitempty"`

	// **参数解释**：修改保护状态   **约束限制**：不涉及   **取值范围**：  - nonProtection: 不保护  - consoleProtection: 控制台修改保护   **默认取值**：不涉及
	ProtectionStatus *UpdateCertificateOptionProtectionStatus `json:"protection_status,omitempty"`

	// **参数解释**：设置修改保护的原因   **约束限制**：仅当protection_status为consoleProtection时有效   **取值范围**：不涉及   **默认取值**：空
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
