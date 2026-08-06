package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateDatakeyCapsuleRequestBody struct {

	// **参数解释：** 密钥ID **约束限制：** UUID格式，满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyId string `json:"key_id"`

	// **参数解释：** 待创建的数据密钥长度 **约束限制：** 256或者128二选一 **取值范围：** - 128 - 256 **默认取值：** 不涉及
	DatakeyLength CreateDatakeyCapsuleRequestBodyDatakeyLength `json:"datakey_length"`

	// **参数解释：** 公钥信息，使用RSAES_OAEP_SHA_256算法加密；如果传递了public_key，KMS会使用该公钥对明文数据密钥进行加密，并返回加密后的数据密钥 **约束限制：** 仅支持RSA公钥 **取值范围：** 不涉及 **默认取值：** 不涉及
	PublicKey *string `json:"public_key,omitempty"`

	// **参数解释：** 密钥策略ID和内联的密钥策略二选一 **约束限制：** 仅支持RSA公钥 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释：** 密钥策略ID和内联的密钥策略二选一 **约束限制：** 仅支持RSA公钥 **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyPolicy *string `json:"key_policy,omitempty"`
}

func (o CreateDatakeyCapsuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyCapsuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDatakeyCapsuleRequestBody", string(data)}, " ")
}

type CreateDatakeyCapsuleRequestBodyDatakeyLength struct {
	value string
}

type CreateDatakeyCapsuleRequestBodyDatakeyLengthEnum struct {
	AES_256 CreateDatakeyCapsuleRequestBodyDatakeyLength
	SM4     CreateDatakeyCapsuleRequestBodyDatakeyLength
}

func GetCreateDatakeyCapsuleRequestBodyDatakeyLengthEnum() CreateDatakeyCapsuleRequestBodyDatakeyLengthEnum {
	return CreateDatakeyCapsuleRequestBodyDatakeyLengthEnum{
		AES_256: CreateDatakeyCapsuleRequestBodyDatakeyLength{
			value: "AES_256",
		},
		SM4: CreateDatakeyCapsuleRequestBodyDatakeyLength{
			value: "SM4",
		},
	}
}

func (c CreateDatakeyCapsuleRequestBodyDatakeyLength) Value() string {
	return c.value
}

func (c CreateDatakeyCapsuleRequestBodyDatakeyLength) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDatakeyCapsuleRequestBodyDatakeyLength) UnmarshalJSON(b []byte) error {
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
