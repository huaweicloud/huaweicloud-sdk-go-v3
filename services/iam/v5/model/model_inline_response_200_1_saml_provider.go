package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// InlineResponse2001SamlProvider **参数解释**： SAML 提供商。  **取值范围**： 不涉及。
type InlineResponse2001SamlProvider struct {

	// **参数解释**： SAML 身份提供商的ID。  **取值范围**： 字符串长度在 1 到 64 之间，并且只能包含：字母、数字、中划线（-）。
	ProviderId string `json:"provider_id"`

	// **参数解释**： SAML 身份提供商的名称。  **取值范围**： 字符串长度在 1 到 64 之间，并且只能包含：字母、数字、下划线（_）、中划线（-）。
	Name string `json:"name"`

	// **参数解释**： 身份提供商描述。  **取值范围**： 字符串长度在 0 到 255 之间，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\"、\"$\"、\"^\"和\"*\"的字符串。
	Description string `json:"description"`

	// **参数解释**： 统一资源名称。  **取值范围**： 字符串长度在 16 到 1500 之间，可以包含：字母、数字、斜杠（/）、等号（=）、下划线（_）、冒号（:）、中划线（-）
	Urn string `json:"urn"`

	// **参数解释**： 支持 SAML 2.0 的身份提供商 (IdP) 的 XML 文档。  **取值范围**： 长度范围为[1000,512000]。
	SamlMetadataDocument string `json:"saml_metadata_document"`

	// **参数解释**： 指定 SAML 身份提供商的加密设置。  **取值范围**： 取值范围为[Required,Allowed]。
	AssertionEncryptionMode InlineResponse2001SamlProviderAssertionEncryptionMode `json:"assertion_encryption_mode"`

	// **参数解释**： 解密 SAML 断言的私钥。  **取值范围**： 不涉及。
	PrivateKeys []InlineResponse200PrivateKeys `json:"private_keys"`

	// **参数解释**： 自定义标签列表。  **取值范围**： 数组长度不涉及。
	Tags []InlineResponse2001SamlProviderTags `json:"tags"`

	// **参数解释**： SAML 身份提供商创建时间。  **取值范围**： 不涉及。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// **参数解释**： SAML 身份提供商过期时间。  **取值范围**： 不涉及。
	ExpiresAt *sdktime.SdkTime `json:"expires_at"`
}

func (o InlineResponse2001SamlProvider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InlineResponse2001SamlProvider struct{}"
	}

	return strings.Join([]string{"InlineResponse2001SamlProvider", string(data)}, " ")
}

type InlineResponse2001SamlProviderAssertionEncryptionMode struct {
	value string
}

type InlineResponse2001SamlProviderAssertionEncryptionModeEnum struct {
	REQUIRED InlineResponse2001SamlProviderAssertionEncryptionMode
	ALLOWED  InlineResponse2001SamlProviderAssertionEncryptionMode
}

func GetInlineResponse2001SamlProviderAssertionEncryptionModeEnum() InlineResponse2001SamlProviderAssertionEncryptionModeEnum {
	return InlineResponse2001SamlProviderAssertionEncryptionModeEnum{
		REQUIRED: InlineResponse2001SamlProviderAssertionEncryptionMode{
			value: "Required",
		},
		ALLOWED: InlineResponse2001SamlProviderAssertionEncryptionMode{
			value: "Allowed",
		},
	}
}

func (c InlineResponse2001SamlProviderAssertionEncryptionMode) Value() string {
	return c.value
}

func (c InlineResponse2001SamlProviderAssertionEncryptionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InlineResponse2001SamlProviderAssertionEncryptionMode) UnmarshalJSON(b []byte) error {
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
