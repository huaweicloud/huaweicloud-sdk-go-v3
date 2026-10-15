package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KafkaSecurity Kafka安全认证相关参数
type KafkaSecurity struct {

	// **参数解释**：  安全协议类型，安全认证时必填，对应Kafka字段：security.protocol。  **约束限制**：  - 安全认证时必填。  **取值范围**：  - PLAINTEXT：无安全认证方式，仅需输入IP和端口进行连接 - SASL_PLAINTEXT：使用SASL机制连接Kafka，需要设置SASL相关配置 - SSL：使用SSL加密方式连接Kafka，需要设置SSL相关配置 - SASL_SSL：使用SASL及SSL加密认证方式，需要设置SSL及SASL相关参数配置信息  **默认取值**：  不涉及
	Type *KafkaSecurityType `json:"type,omitempty"`

	// **参数解释**：  证书名称。  **约束限制**：  - 安全协议为SSL、SASL_SSL时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TrustStoreKeyName *string `json:"trust_store_key_name,omitempty"`

	// **参数解释**：  安全证书base64转码后的值。  **约束限制**：  - 安全协议为SSL、SASL_SSL时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TrustStoreKey *string `json:"trust_store_key,omitempty"`

	// **参数解释**：  证书密码。  **约束限制**：  - 证书设置了密码时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TrustStorePassword *string `json:"trust_store_password,omitempty"`

	// **参数解释**：  主机名端点识别算法，指定通过服务端证书验证服务端主机名的端点识别算法，对应Kafka字段：ssl.endpoint.identification.algorithm。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不填表示禁用主机名验证
	EndpointAlgorithm *string `json:"endpoint_algorithm,omitempty"`

	// **参数解释**：  SASL机制，用于客户端连接的SASL机制，对应Kafka字段：sasl.mechanism。  **约束限制**：  - 认证类型为SASL_PLAINTEXT、SASL_SSL时必填。  **取值范围**：  - GSSAPI：GSSAPI机制 - PLAIN：PLAIN机制 - SCRAM-SHA-256：SCRAM-SHA-256机制 - SCRAM-SHA-512：SCRAM-SHA-512机制  **默认取值**：  不涉及
	SaslMechanism *string `json:"sasl_mechanism,omitempty"`

	// **参数解释**：  是否为委托令牌鉴权。  **约束限制**：  - 安全协议为SASL_SSL和SASL_PLAINTEXT时，SASL机制选择SCRAM-SHA-256或者SCRAM-SHA-512时生效。  **取值范围**：  - true：使用委托令牌鉴权 - false：不使用委托令牌鉴权  **默认取值**：  不涉及
	DelegationTokens *bool `json:"delegation_tokens,omitempty"`

	// **参数解释**：  是否开启SSL双向认证。  **约束限制**：  不涉及。  **取值范围**：  - true：开启SSL双向认证 - false：不开启SSL双向认证  **默认取值**：  不涉及
	EnableKeyStore *bool `json:"enable_key_store,omitempty"`

	// **参数解释**：  Keystore证书。  **约束限制**：  - 开启SSL双向认证时需要。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	KeyStoreKey *string `json:"key_store_key,omitempty"`

	// **参数解释**：  Keystore证书名称。  **约束限制**：  - 开启SSL双向认证时需要。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	KeyStoreKeyName *string `json:"key_store_key_name,omitempty"`

	// **参数解释**：  Keystore证书密码，对应Kafka字段：ssl.keystore.password。  **约束限制**：  - 证书设置了密码时需要。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	KeyStorePassword *string `json:"key_store_password,omitempty"`

	// **参数解释**：  是否设置Keystore私钥密码。  **约束限制**：  不涉及。  **取值范围**：  - true：设置Keystore私钥密码 - false：不设置Keystore私钥密码  **默认取值**：  false
	SetPrivateKeyPassword *bool `json:"set_private_key_password,omitempty"`

	// **参数解释**：  Keystore私钥密码，对应Kafka字段：ssl.key.password。  **约束限制**：  - 开启SSL双向认证时，set_private_key_password为true时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	KeyPassword *string `json:"key_password,omitempty"`
}

func (o KafkaSecurity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaSecurity struct{}"
	}

	return strings.Join([]string{"KafkaSecurity", string(data)}, " ")
}

type KafkaSecurityType struct {
	value string
}

type KafkaSecurityTypeEnum struct {
	PLAINTEXT      KafkaSecurityType
	SASL_PLAINTEXT KafkaSecurityType
	SASL_SSL       KafkaSecurityType
	SSL            KafkaSecurityType
}

func GetKafkaSecurityTypeEnum() KafkaSecurityTypeEnum {
	return KafkaSecurityTypeEnum{
		PLAINTEXT: KafkaSecurityType{
			value: "PLAINTEXT",
		},
		SASL_PLAINTEXT: KafkaSecurityType{
			value: "SASL_PLAINTEXT",
		},
		SASL_SSL: KafkaSecurityType{
			value: "SASL_SSL",
		},
		SSL: KafkaSecurityType{
			value: "SSL",
		},
	}
}

func (c KafkaSecurityType) Value() string {
	return c.value
}

func (c KafkaSecurityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KafkaSecurityType) UnmarshalJSON(b []byte) error {
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
