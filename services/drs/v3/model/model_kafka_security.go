package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KafkaSecurity Kafka安全认证相关参数
type KafkaSecurity struct {

	// 安全协议，安全认证时必填，对应Kafka字段：security.protocol。 - PLAINTEXT：无安全认证方式，仅需输入IP和端口进行连接。 - SASL_PLAINTEXT：使用SASL机制连接Kafka，需要设置SASL相关配置。 - SSL：使用SSL加密方式连接Kafka，需要设置SSL相关配置。 - SASL_SSL：使用SASL及SSL加密认证方式，需要设置SSL及SASL相关参数配置信息。
	Type *KafkaSecurityType `json:"type,omitempty"`

	// 证书名称，安全协议为SSL、SASL_SSL时必填。
	TrustStoreKeyName *string `json:"trust_store_key_name,omitempty"`

	// 安全证书base64转码后的值，安全协议为SSL、SASL_SSL时必填。
	TrustStoreKey *string `json:"trust_store_key,omitempty"`

	// 证书密码，证书设置了密码时必填。
	TrustStorePassword *string `json:"trust_store_password,omitempty"`

	// 主机名端点识别算法，指定通过服务端证书验证服务端主机名的端点识别算法，不填表示禁用主机名验证。对应Kafka字段：ssl.endpoint.identification.algorithm。
	EndpointAlgorithm *string `json:"endpoint_algorithm,omitempty"`

	// SASL机制，用于客户端连接的SASL机制，对应Kafka字段：sasl.mechanism，支持以下四项，取值： - GSSAPI - PLAIN - SCRAM-SHA-256 - SCRAM-SHA-512
	SaslMechanism *string `json:"sasl_mechanism,omitempty"`

	// 是否为委托令牌鉴权，安全协议为SASL_SSL和SASL_PLAINTEXT时，SASL机制选择“SCRAM-SHA-256”或者“SCRAM-SHA-512”时生效。
	DelegationTokens *bool `json:"delegation_tokens,omitempty"`

	// 是否开启SSL双向认证。
	EnableKeyStore *bool `json:"enable_key_store,omitempty"`

	// Keystore证书，开启SSL双向认证时需要。
	KeyStoreKey *string `json:"key_store_key,omitempty"`

	// Keystore证书名称，开启SSL双向认证时需要。
	KeyStoreKeyName *string `json:"key_store_key_name,omitempty"`

	// Keystore证书密码，证书设置了密码时需要。对应Kafka字段：ssl.keystore.password。
	KeyStorePassword *string `json:"key_store_password,omitempty"`

	// 是否设置Keystore私钥密码，默认为false。
	SetPrivateKeyPassword *bool `json:"set_private_key_password,omitempty"`

	// Keystore私钥密码，开启SSL双向认证时，set_private_key_password为true时必填。对应Kafka字段：ssl.key.password。
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
