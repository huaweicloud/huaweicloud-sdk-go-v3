package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CertificateInfo **参数解释**：ELB证书信息。注意：真正的证书在内层字段中。
type CertificateInfo struct {

	// **参数解释**：证书的管理状态。该字段当前无用，设置为true或者false都不影响证书使用。  **取值范围**： - true：表示证书可用。 - false：表示证书不可用。
	AdminStateUp bool `json:"admin_state_up"`

	// **参数解释**：证书内容。支持最大11层证书链(含证书和证书链)。  **取值范围**：PEM编码格式，最大长度65536个字符。
	Certificate string `json:"certificate"`

	// **参数解释**：证书的描述。  **取值范围**：0-255个字符。
	Description string `json:"description"`

	// **参数解释**：服务器证书所签域名。  **取值范围**：总长度为0-10000，由若干普通域名或泛域名组成，域名之间以\",\"分隔，不超过100个域名。 - 普通域名：由若干字符串组成，字符串间以\".\"分隔，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。例：www.test.com。 - 泛域名：在普通域名的基础上仅允许首字母为\"\\*\"。例：\\*.test.com。
	Domain string `json:"domain"`

	// **参数解释**：ELB证书管理对象ID。  **取值范围**：由32位数字和小写字母组成。
	Id string `json:"id"`

	// **参数解释**：证书的名称。  **取值范围**：0-255个字符。
	Name string `json:"name"`

	// **参数解释**：服务器证书的私钥。  **取值范围**：PEM编码格式，最大长度8192个字符。
	PrivateKey string `json:"private_key"`

	// **参数解释**：证书的类型。  **取值范围**： - server：服务器证书。 - client：CA证书。 - server_sm：服务器SM双证书。
	Type string `json:"type"`

	// **参数解释**：创建时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。
	CreatedAt string `json:"created_at"`

	// **参数解释**：更新时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。
	UpdatedAt string `json:"updated_at"`

	// **参数解释**：证书有效期的截止时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。
	ExpireTime string `json:"expire_time"`

	// **参数解释**：项目ID。获取方式请参见[获取项目ID](elb_fl_0008.xml)。  **取值范围**：长度为32个字符，由小写字母和数字组成。
	ProjectId string `json:"project_id"`

	// **参数解释**：服务器SM双证书的证书内容。 支持最大11层证书链(含证书和证书链)。  **取值范围**：PEM编码格式。最大长度65536字符。
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// **参数解释**：服务器SM双证书的私钥。  **取值范围**：PEM编码格式，最大长度8192个字符。
	EncPrivateKey *string `json:"enc_private_key,omitempty"`

	// **参数解释**：云证书与管理服务（CCM）中的证书ID。  **取值范围**：不涉及
	ScmCertificateId *string `json:"scm_certificate_id,omitempty"`

	// **参数解释**：证书绑定的主域名。  **取值范围**：不涉及
	CommonName *string `json:"common_name,omitempty"`

	// **参数解释**：证书指纹。  **取值范围**：不涉及
	Fingerprint *string `json:"fingerprint,omitempty"`

	// **参数解释**：证书绑定的域名列表。  **取值范围**：不涉及
	SubjectAlternativeNames *[]string `json:"subject_alternative_names,omitempty"`

	// **参数解释**：标记当前证书来源。  **取值范围**： - scm：表示关联云证书与管理服务（CCM）中的证书。 - 空值：表示自有证书。
	Source *string `json:"source,omitempty"`

	// **参数解释**：修改保护状态。  **取值范围**：  - nonProtection: 不保护 - consoleProtection: 控制台修改保护，即禁止通过控制台修改。
	ProtectionStatus *CertificateInfoProtectionStatus `json:"protection_status,omitempty"`

	// **参数解释**：修改保护的原因。  **取值范围**：不涉及
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// **参数解释**：资源所属的企业项目ID。  **取值范围**： - \"0\"：表示资源属于default企业项目。 - UUID格式的字符串，表示非默认企业项目。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CertificateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateInfo struct{}"
	}

	return strings.Join([]string{"CertificateInfo", string(data)}, " ")
}

type CertificateInfoProtectionStatus struct {
	value string
}

type CertificateInfoProtectionStatusEnum struct {
	NON_PROTECTION     CertificateInfoProtectionStatus
	CONSOLE_PROTECTION CertificateInfoProtectionStatus
}

func GetCertificateInfoProtectionStatusEnum() CertificateInfoProtectionStatusEnum {
	return CertificateInfoProtectionStatusEnum{
		NON_PROTECTION: CertificateInfoProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CertificateInfoProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CertificateInfoProtectionStatus) Value() string {
	return c.value
}

func (c CertificateInfoProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateInfoProtectionStatus) UnmarshalJSON(b []byte) error {
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
