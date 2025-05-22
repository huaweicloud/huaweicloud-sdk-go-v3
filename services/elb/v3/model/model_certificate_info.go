package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CertificateInfo 证书信息。
type CertificateInfo struct {

	// 证书的管理状态。  不支持该字段，请勿使用。
	AdminStateUp bool `json:"admin_state_up"`

	// 证书的内容。PEM编码格式。
	Certificate string `json:"certificate"`

	// 证书的描述。
	Description string `json:"description"`

	// 服务器证书所签域名。该字段仅type为server时有效。  总长度为0-10000，由若干普通域名或泛域名组成，域名之间以\",\"分隔，不超过100个域名。  普通域名：由若干字符串组成，字符串间以\".\"分隔，单个字符串长度不超过63个字符， 只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。例：www.test.com。  泛域名：在普通域名的基础上仅允许首字母为\"\\*\"。例：\\*.test.com
	Domain string `json:"domain"`

	// 证书ID。
	Id string `json:"id"`

	// 证书的名称。
	Name string `json:"name"`

	// 服务器证书的私钥。PEM编码格式。  当type为client时，该参数被忽略，不影响证书的创建和使用。  当type为server时，该字段必须符合格式要求，且私钥必须是有效的。
	PrivateKey string `json:"private_key"`

	// SSL证书的类型。分为服务器证书(server)、CA证书(client)。默认值：server。
	Type string `json:"type"`

	// 证书创建时间。
	CreatedAt string `json:"created_at"`

	// 证书更新时间。
	UpdatedAt string `json:"updated_at"`

	// 证书使用截止时间。
	ExpireTime string `json:"expire_time"`

	// 证书所在项目ID。
	ProjectId string `json:"project_id"`

	// HTTPS协议使用的SM加密证书内容。  取值：PEM编码格式。  注意：仅在当前局点的SM加密证书特性开启才会返回该字段。
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// HTTPS协议使用的SM加密证书私钥。  取值：PEM编码格式。  注意：仅在当前局点的SM加密证书特性开启才会返回该字段。
	EncPrivateKey *string `json:"enc_private_key,omitempty"`

	// SCM证书ID。
	ScmCertificateId *string `json:"scm_certificate_id,omitempty"`

	// 证书绑定的主域名。
	CommonName *string `json:"common_name,omitempty"`

	// 证书指纹。
	Fingerprint *string `json:"fingerprint,omitempty"`

	// 证书绑定的所有域名。
	SubjectAlternativeNames *[]string `json:"subject_alternative_names,omitempty"`

	// 参数解释：证书来源 取值范围：无  默认取值：当scm_certificate_id不为空，且未传入source时，默认取值为“scm”； 其他情况下默认为空。
	Source *string `json:"source,omitempty"`

	// 参数解释：修改保护状态  约束限制：无  取值范围：  - nonProtection: 不保护  - consoleProtection: 控制台修改保护  默认取值：nonProtection
	ProtectionStatus *CertificateInfoProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释：设置修改保护的原因  约束限制：仅当protection_status为consoleProtection时有效  取值范围：无  默认取值：空
	ProtectionReason *string `json:"protection_reason,omitempty"`
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
