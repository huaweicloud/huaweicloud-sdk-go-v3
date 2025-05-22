package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LargestKeyPrefix **参数解释**： 前缀Key。 **取值范围**： 不涉及。
type LargestKeyPrefix struct {

	// **参数解释**： 前缀名称。 **取值范围**： 不涉及。
	KeyPrefix *string `json:"key_prefix,omitempty"`

	// **参数解释**： Key的数据类型。 **取值范围**： string list set zset hash
	Type *LargestKeyPrefixType `json:"type,omitempty"`

	// **参数解释**： Key的大小，单位：Bytes。 **取值范围**： 不涉及。
	Bytes *int32 `json:"bytes,omitempty"`

	// **参数解释**： 相同前缀key的数量。 **取值范围**： 不涉及。
	Num *int32 `json:"num,omitempty"`
}

func (o LargestKeyPrefix) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LargestKeyPrefix struct{}"
	}

	return strings.Join([]string{"LargestKeyPrefix", string(data)}, " ")
}

type LargestKeyPrefixType struct {
	value string
}

type LargestKeyPrefixTypeEnum struct {
	STRING LargestKeyPrefixType
	LIST   LargestKeyPrefixType
	SET    LargestKeyPrefixType
	ZSET   LargestKeyPrefixType
	HASH   LargestKeyPrefixType
}

func GetLargestKeyPrefixTypeEnum() LargestKeyPrefixTypeEnum {
	return LargestKeyPrefixTypeEnum{
		STRING: LargestKeyPrefixType{
			value: "string",
		},
		LIST: LargestKeyPrefixType{
			value: "list",
		},
		SET: LargestKeyPrefixType{
			value: "set",
		},
		ZSET: LargestKeyPrefixType{
			value: "zset",
		},
		HASH: LargestKeyPrefixType{
			value: "hash",
		},
	}
}

func (c LargestKeyPrefixType) Value() string {
	return c.value
}

func (c LargestKeyPrefixType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LargestKeyPrefixType) UnmarshalJSON(b []byte) error {
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
