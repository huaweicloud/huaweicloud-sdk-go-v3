package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSamlProviderReqBody struct {

	// **参数解释**： SAML 提供商的名称。  **约束限制**： 字符串长度在 1 到 64 之间，并且只能包含：字母、数字、下划线（_）、中划线（-）。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 由支持 SAML 2.0 的身份提供商 (IdP) 生成的 XML 文档。该文档包含颁发者的名称、过期信息，以及可用于验证从 IdP 接收到的 SAML 身份验证响应（断言）的密钥。  **约束限制**： 长度范围为[1000,512000]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	SamlMetadataDocument string `json:"saml_metadata_document"`

	// **参数解释**： 指定 SAML 提供商的加密设置。  **约束限制**： 不涉及。  **取值范围**： 取值范围为[Required,Allowed]。  **默认取值**： 不涉及。
	AssertionEncryptionMode *CreateSamlProviderReqBodyAssertionEncryptionMode `json:"assertion_encryption_mode,omitempty"`

	// **参数解释**： 添加解密 SAML 断言的私钥，必须是一个 PEM 格式的 RSA 私钥。在接收到加密的 SAML 断言时，IAM 会基于 RSA-OAEP 算法使用该私钥解密得到用于加密 SAML 断言的对称密钥，然后再基于 AES-GCM 或 AES-CBC 加密算法使用对称密钥解密出 SAML 断言明文。  **约束限制**： 长度范围为[1,16384]。  **取值范围**： 字符串必须由 1 个或多个字符组成，这些字符可以是：空格、可见 ASCII 字符、Latin-1 扩展字符、Tab、换行、回车  **默认取值**： 不涉及。
	AddPrivateKey *string `json:"add_private_key,omitempty"`

	// **参数解释**： 身份提供商描述。  **约束限制**： 长度范围为[0,255]。  **取值范围**： 不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\"、\"$\"、\"^\"和\"*\"的字符串。  **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o CreateSamlProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSamlProviderReqBody struct{}"
	}

	return strings.Join([]string{"CreateSamlProviderReqBody", string(data)}, " ")
}

type CreateSamlProviderReqBodyAssertionEncryptionMode struct {
	value string
}

type CreateSamlProviderReqBodyAssertionEncryptionModeEnum struct {
	REQUIRED CreateSamlProviderReqBodyAssertionEncryptionMode
	ALLOWED  CreateSamlProviderReqBodyAssertionEncryptionMode
}

func GetCreateSamlProviderReqBodyAssertionEncryptionModeEnum() CreateSamlProviderReqBodyAssertionEncryptionModeEnum {
	return CreateSamlProviderReqBodyAssertionEncryptionModeEnum{
		REQUIRED: CreateSamlProviderReqBodyAssertionEncryptionMode{
			value: "Required",
		},
		ALLOWED: CreateSamlProviderReqBodyAssertionEncryptionMode{
			value: "Allowed",
		},
	}
}

func (c CreateSamlProviderReqBodyAssertionEncryptionMode) Value() string {
	return c.value
}

func (c CreateSamlProviderReqBodyAssertionEncryptionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSamlProviderReqBodyAssertionEncryptionMode) UnmarshalJSON(b []byte) error {
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
