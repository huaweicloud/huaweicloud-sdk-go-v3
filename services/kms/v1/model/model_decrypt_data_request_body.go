package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DecryptDataRequestBody struct {

	// 被加密数据密文。取值为加密数据结果中的cipher_text的值，满足正则匹配“^[0-9a-zA-Z+/=]{188,5648}$”。
	CipherText *string `json:"cipher_text,omitempty" xml:"cipher_text"`

	// 一系列key-value键值对，用于记录资源上下文信息，用于保护数据的完整性，不应包含敏感信息，最大长度为8192。 当在加密时指定了该参数时，解密密文时，需要传入相同的参数，才能正确的解密。 例如：{\"Key1\":\"Value1\",\"Key2\":\"Value2\"}
	EncryptionContext *interface{} `json:"encryption_context,omitempty" xml:"encryption_context"`

	// 数据加密算法，仅使用非对称密钥需要指定该参数，默认值为“SYMMETRIC_DEFAULT”，合法枚举值如下：  - SYMMETRIC_DEFAULT  - RSAES_OAEP_SHA_256  - RSAES_OAEP_SHA_1  - SM2_ENCRYPT
	EncryptionAlgorithm *DecryptDataRequestBodyEncryptionAlgorithm `json:"encryption_algorithm,omitempty" xml:"encryption_algorithm"`

	// 密钥ID，36字节，满足正则匹配“^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$”。仅当密文使用非对称密钥加密时才需要此参数。 例如：0d0466b0-e727-4d9c-b35d-f84bb474a37f。
	KeyId *string `json:"key_id,omitempty" xml:"key_id"`

	// 请求消息序列号，36字节序列号。 例如：919c82d4-8046-4722-9094-35c3c6524cff
	Sequence *string `json:"sequence,omitempty" xml:"sequence"`
}

func (o DecryptDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDataRequestBody struct{}"
	}

	return strings.Join([]string{"DecryptDataRequestBody", string(data)}, " ")
}

type DecryptDataRequestBodyEncryptionAlgorithm struct {
	value string
}

type DecryptDataRequestBodyEncryptionAlgorithmEnum struct {
	SYMMETRIC_DEFAULT  DecryptDataRequestBodyEncryptionAlgorithm
	RSAES_OAEP_SHA_256 DecryptDataRequestBodyEncryptionAlgorithm
	RSAES_OAEP_SHA_1   DecryptDataRequestBodyEncryptionAlgorithm
	SM2_ENCRYPT        DecryptDataRequestBodyEncryptionAlgorithm
}

func GetDecryptDataRequestBodyEncryptionAlgorithmEnum() DecryptDataRequestBodyEncryptionAlgorithmEnum {
	return DecryptDataRequestBodyEncryptionAlgorithmEnum{
		SYMMETRIC_DEFAULT: DecryptDataRequestBodyEncryptionAlgorithm{
			value: "SYMMETRIC_DEFAULT",
		},
		RSAES_OAEP_SHA_256: DecryptDataRequestBodyEncryptionAlgorithm{
			value: "RSAES_OAEP_SHA_256",
		},
		RSAES_OAEP_SHA_1: DecryptDataRequestBodyEncryptionAlgorithm{
			value: "RSAES_OAEP_SHA_1",
		},
		SM2_ENCRYPT: DecryptDataRequestBodyEncryptionAlgorithm{
			value: "SM2_ENCRYPT",
		},
	}
}

func (c DecryptDataRequestBodyEncryptionAlgorithm) Value() string {
	return c.value
}

func (c DecryptDataRequestBodyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DecryptDataRequestBodyEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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
