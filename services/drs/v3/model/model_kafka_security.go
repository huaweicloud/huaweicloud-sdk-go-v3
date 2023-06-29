package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KafkaSecurity Kafka安全认证相关参数
type KafkaSecurity struct {

	// 证书名称，使用安全认证时必填。
	TrustStoreKeyName *string `json:"trust_store_key_name,omitempty"`

	// 安全证书base64转码后的值，使用安全认证时必填。
	TrustStoreKey *string `json:"trust_store_key,omitempty"`

	// 证书密码，使用安全认证时必填。
	TrustStorePassword *string `json:"trust_store_password,omitempty"`

	// 认证类型，PLAINTEXT为无认证，，使用安全认证时必填。
	Type *KafkaSecurityType `json:"type,omitempty"`
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
	PLAINTEXT KafkaSecurityType
	SASL_SSL  KafkaSecurityType
}

func GetKafkaSecurityTypeEnum() KafkaSecurityTypeEnum {
	return KafkaSecurityTypeEnum{
		PLAINTEXT: KafkaSecurityType{
			value: "PLAINTEXT",
		},
		SASL_SSL: KafkaSecurityType{
			value: "SASL_SSL",
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
