package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateCertificateOption **参数解释**：创建证书请求参数。  **约束限制**：不涉及
type CreateCertificateOption struct {

	// **参数解释**：证书的管理状态。该字段当前无用，设置为true或者false都不影响证书使用。  **约束限制**：不涉及  **取值范围**： - true：表示证书可用。 - false：表示证书不可用。  **默认取值**：true
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**：证书内容。支持最大11层证书链(含证书和证书链)。  **约束限制**：不涉及  **取值范围**：PEM编码格式，最大长度65536个字符。  **默认取值**：不涉及
	Certificate *string `json:"certificate,omitempty"`

	// **参数解释**：证书的描述。  **约束限制**：不涉及  **取值范围**：0-255个字符。  **默认取值**：不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：服务器证书所签域名。  **约束限制**：该字段仅type为server时有效（其他类型证书，字段可传入，但不会生效）。  **取值范围**：总长度为0-10000，由若干普通域名或泛域名组成，域名之间以\",\"分隔，不超过100个域名。 - 普通域名：由若干字符串组成，字符串间以\".\"分隔，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。例：www.test.com。 - 泛域名：在普通域名的基础上仅允许首字母为\"\\*\"。例：\\*.test.com。  **默认取值**：不涉及
	Domain *string `json:"domain,omitempty"`

	// **参数解释**：证书的名称。  **约束限制**：不涉及  **取值范围**：0-255个字符。  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：服务器证书的私钥。  **约束限制**： - 当type为server和server_sm时，创建时必须传入。 - 当type为其他值时，字段无用，可以不传入；若传入则必须符合PEM格式。  **取值范围**：PEM编码格式，最大长度8192个字符。  **默认取值**：不涉及
	PrivateKey *string `json:"private_key,omitempty"`

	// **参数解释**：项目ID。获取方式请参见[获取项目ID](elb_fl_0008.xml)。  **约束限制**：不涉及  **取值范围**：长度为32个字符，由小写字母和数字组成。  **默认取值**：不涉及  > 该字段实际无效，最终使用url中的project_id。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：证书的类型。  **约束限制**：不涉及  **取值范围**： - server：服务器证书。 - client：CA证书。 - server_sm：服务器SM双证书。  **默认取值**：server
	Type *CreateCertificateOptionType `json:"type,omitempty"`

	// **参数解释**：资源所属的企业项目ID。创建时不传则资源属于default企业项目，返回enterprise_project_id=\"0\"。  **约束限制**：不能传入空字符串\"\"、\"0\"或不存在的企业项目ID。  **取值范围**：不涉及  **默认取值**：\"0\"  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：服务器SM双证书的证书内容。支持最大11层证书链(含证书和证书链)。  **约束限制**：仅当type为server_sm时，才支持且必须传入。  **取值范围**：PEM编码格式。最大长度65536字符。  **默认取值**：不涉及
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// **参数解释**：服务器SM双证书的私钥。  **约束限制**：仅当type为server_sm时，才支持且必须传入。  **取值范围**：PEM编码格式，最大长度8192个字符。  **默认取值**：不涉及
	EncPrivateKey *string `json:"enc_private_key,omitempty"`

	// **参数解释**：云证书与管理服务（CCM）中的证书ID。  **约束限制**：仅记录证书ID，不验证其是否真实存在云证书与管理服务中。并且需要将云证书与管理服务中对应证书的内容手动设置到当前接口相应字段中（可能涉及字段certificate、private_key、enc_certificate和enc_private_key）  **取值范围**：不涉及  **默认取值**：不涉及
	ScmCertificateId *string `json:"scm_certificate_id,omitempty"`

	// **参数解释**：标记当前证书来源。  **约束限制**：无  **取值范围**： - scm：表示关联云证书与管理服务（CCM）中的证书。 - 空值：表示自有证书。  **默认取值**：当scm_certificate_id不为空，默认取值为\"scm\"。否则默认为空值。
	Source *string `json:"source,omitempty"`

	// **参数解释**：修改保护状态。  **约束限制**：不涉及  **取值范围**： - nonProtection: 不保护 - consoleProtection: 控制台修改保护，即禁止通过控制台修改。  **默认取值**：nonProtection
	ProtectionStatus *CreateCertificateOptionProtectionStatus `json:"protection_status,omitempty"`

	// **参数解释**：修改保护的原因。  **约束限制**：仅当protection_status为consoleProtection时有效。  **取值范围**：不涉及  **默认取值**：空
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o CreateCertificateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateOption struct{}"
	}

	return strings.Join([]string{"CreateCertificateOption", string(data)}, " ")
}

type CreateCertificateOptionType struct {
	value string
}

type CreateCertificateOptionTypeEnum struct {
	SERVER    CreateCertificateOptionType
	CLIENT    CreateCertificateOptionType
	SERVER_SM CreateCertificateOptionType
}

func GetCreateCertificateOptionTypeEnum() CreateCertificateOptionTypeEnum {
	return CreateCertificateOptionTypeEnum{
		SERVER: CreateCertificateOptionType{
			value: "server",
		},
		CLIENT: CreateCertificateOptionType{
			value: "client",
		},
		SERVER_SM: CreateCertificateOptionType{
			value: "server_sm",
		},
	}
}

func (c CreateCertificateOptionType) Value() string {
	return c.value
}

func (c CreateCertificateOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCertificateOptionType) UnmarshalJSON(b []byte) error {
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

type CreateCertificateOptionProtectionStatus struct {
	value string
}

type CreateCertificateOptionProtectionStatusEnum struct {
	NON_PROTECTION     CreateCertificateOptionProtectionStatus
	CONSOLE_PROTECTION CreateCertificateOptionProtectionStatus
}

func GetCreateCertificateOptionProtectionStatusEnum() CreateCertificateOptionProtectionStatusEnum {
	return CreateCertificateOptionProtectionStatusEnum{
		NON_PROTECTION: CreateCertificateOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreateCertificateOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreateCertificateOptionProtectionStatus) Value() string {
	return c.value
}

func (c CreateCertificateOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCertificateOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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
