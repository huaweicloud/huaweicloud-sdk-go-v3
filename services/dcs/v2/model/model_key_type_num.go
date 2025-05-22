package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KeyTypeNum **参数解释**： Key类型总数量。 **取值范围**： 不涉及。
type KeyTypeNum struct {

	// **参数解释**： Key的数据类型。 **取值范围**： string list set zset hash
	Type *KeyTypeNumType `json:"type,omitempty"`

	// **参数解释**： Key类型的总数量。 **取值范围**： 不涉及。
	Num *int32 `json:"num,omitempty"`
}

func (o KeyTypeNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyTypeNum struct{}"
	}

	return strings.Join([]string{"KeyTypeNum", string(data)}, " ")
}

type KeyTypeNumType struct {
	value string
}

type KeyTypeNumTypeEnum struct {
	STRING KeyTypeNumType
	LIST   KeyTypeNumType
	SET    KeyTypeNumType
	ZSET   KeyTypeNumType
	HASH   KeyTypeNumType
}

func GetKeyTypeNumTypeEnum() KeyTypeNumTypeEnum {
	return KeyTypeNumTypeEnum{
		STRING: KeyTypeNumType{
			value: "string",
		},
		LIST: KeyTypeNumType{
			value: "list",
		},
		SET: KeyTypeNumType{
			value: "set",
		},
		ZSET: KeyTypeNumType{
			value: "zset",
		},
		HASH: KeyTypeNumType{
			value: "hash",
		},
	}
}

func (c KeyTypeNumType) Value() string {
	return c.value
}

func (c KeyTypeNumType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeyTypeNumType) UnmarshalJSON(b []byte) error {
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
