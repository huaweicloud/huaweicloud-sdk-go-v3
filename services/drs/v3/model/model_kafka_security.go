package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Kafka安全认证相关参数
type KafkaSecurity struct {
	// 证书名称。

	TrustStoreKeyName *string `json:"trust_store_key_name,omitempty"`
	// 安全证书base64转码后的值。

	TrustStoreKey *string `json:"trust_store_key,omitempty"`
	// 证书密码。

	TrustStorePassword *string `json:"trust_store_password,omitempty"`
	// 认证类型，PLAINTEXT为无认证：

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

func (c KafkaSecurityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KafkaSecurityType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
