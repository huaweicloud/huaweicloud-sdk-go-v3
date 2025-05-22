package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KeyTypeByte Key类型总大小。
type KeyTypeByte struct {

	// **参数解释**： Key的数据类型。 **取值范围**： string list set zset hash
	Type *KeyTypeByteType `json:"type,omitempty"`

	// **参数解释**： 每种数据类型Key的总大小，单位：Bytes。 **取值范围**： 不涉及。
	Bytes *int32 `json:"bytes,omitempty"`
}

func (o KeyTypeByte) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyTypeByte struct{}"
	}

	return strings.Join([]string{"KeyTypeByte", string(data)}, " ")
}

type KeyTypeByteType struct {
	value string
}

type KeyTypeByteTypeEnum struct {
	STRING KeyTypeByteType
	LIST   KeyTypeByteType
	SET    KeyTypeByteType
	ZSET   KeyTypeByteType
	HASH   KeyTypeByteType
}

func GetKeyTypeByteTypeEnum() KeyTypeByteTypeEnum {
	return KeyTypeByteTypeEnum{
		STRING: KeyTypeByteType{
			value: "string",
		},
		LIST: KeyTypeByteType{
			value: "list",
		},
		SET: KeyTypeByteType{
			value: "set",
		},
		ZSET: KeyTypeByteType{
			value: "zset",
		},
		HASH: KeyTypeByteType{
			value: "hash",
		},
	}
}

func (c KeyTypeByteType) Value() string {
	return c.value
}

func (c KeyTypeByteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeyTypeByteType) UnmarshalJSON(b []byte) error {
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
