package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LargestKey **参数解释**： 大key记录。 **取值范围**： 不涉及。
type LargestKey struct {

	// **参数解释**： Key名称。 **取值范围**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： Key的数据类型。 **取值范围**： string list set zset hash
	Type *LargestKeyType `json:"type,omitempty"`

	// **参数解释**： Key的大小，单位：Bytes。 **取值范围**： 不涉及。
	Bytes *int32 `json:"bytes,omitempty"`

	// **参数解释**： 元素数量或元素大小（String类型为元素大小，单位：Bytes。非String类型为元素数量）。 **取值范围**： 不涉及。
	NumOfElem *int32 `json:"num_of_elem,omitempty"`
}

func (o LargestKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LargestKey struct{}"
	}

	return strings.Join([]string{"LargestKey", string(data)}, " ")
}

type LargestKeyType struct {
	value string
}

type LargestKeyTypeEnum struct {
	STRING LargestKeyType
	LIST   LargestKeyType
	SET    LargestKeyType
	ZSET   LargestKeyType
	HASH   LargestKeyType
}

func GetLargestKeyTypeEnum() LargestKeyTypeEnum {
	return LargestKeyTypeEnum{
		STRING: LargestKeyType{
			value: "string",
		},
		LIST: LargestKeyType{
			value: "list",
		},
		SET: LargestKeyType{
			value: "set",
		},
		ZSET: LargestKeyType{
			value: "zset",
		},
		HASH: LargestKeyType{
			value: "hash",
		},
	}
}

func (c LargestKeyType) Value() string {
	return c.value
}

func (c LargestKeyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LargestKeyType) UnmarshalJSON(b []byte) error {
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
