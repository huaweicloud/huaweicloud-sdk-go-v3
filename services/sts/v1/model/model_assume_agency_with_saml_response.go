package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyWithSamlResponse Response Object
type AssumeAgencyWithSamlResponse struct {

	// **参数解释**： SAML断言中SourceIdentity属性值所申明的身份。  **约束限制**： 不涉及。
	SourceIdentity *string `json:"source_identity,omitempty"`

	AssumedAgency *AssumedAgencyWithFederationDto `json:"assumed_agency,omitempty"`

	Credentials *CredentialsDto `json:"credentials,omitempty"`

	// **参数解释**：  SAML断言中SubjectConfirmationData元素的Recipient属性值。  **约束限制**： 不涉及。
	Audience *string `json:"audience,omitempty"`

	// **参数解释**： SAML断言中Issuer元素的值。  **约束限制**： 不涉及。
	Issuer *string `json:"issuer,omitempty"`

	// **参数解释**： 以下三部分的哈希值：issuer、华为云账号的Account ID以及IAM中SAML提供商的名称（URN的最后一部分）。name_qualifier和subject的组合可用于唯一标识用户。下面的伪代码展示了哈希值的计算方式：BASE64 ( SHA1 ( \"https://example.com/saml\" + \"8c1eef3a241945f69c3d3axxxxxxxxxx\" + \"/MySAMLIdPName\" ) )  **约束限制**： 不涉及。
	NameQualifier *string `json:"name_qualifier,omitempty"`

	// **参数解释**：  SAML断言中Subject元素的NameID元素的值。  **约束限制**： 不涉及。
	Subject *string `json:"subject,omitempty"`

	// **参数解释**：  NameID的格式，由SAML断言中NameID元素的Format属性定义。格式的典型示例是transient（临时）或persistent（持久）。 如果该格式包含前缀urn:oasis:names:tc:SAML:2.0:nameid-format，该前缀将被移除。例如，urn:oasis:names:tc:SAML:2.0:nameid-format:transient将作为transient返回。如果格式包含任何其他前缀，则直接返回该格式而不作任何修改。  **约束限制**： 不涉及。
	SubjectType    *string `json:"subject_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssumeAgencyWithSamlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyWithSamlResponse struct{}"
	}

	return strings.Join([]string{"AssumeAgencyWithSamlResponse", string(data)}, " ")
}
