package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MessageAttribute 消息属性
type MessageAttribute struct {

	// 属性名称。属性名称只能包含小写英文字母([a-z])、数字([0-9])、下划线(_)，下划线不能出现在开始或结尾，下划线不能连续出现，长度为1到32个字符
	Name string `json:"name"`

	// 属性类型 STRING：字符串（String）类型 STRING_ARRAY：字符串数组（String.Array）类型 PROTOCOL：协议类型
	Type MessageAttributeType `json:"type"`

	// 属性值。 当属性类型为“STRING”时，属性值只能包含中英文、数字、下划线，长度为1到32个字符。 当属性类型为“STRING_ARRAY”时，属性值为字符串数组，数组长度为1到10，数组中的元素内容不能重复，数组中的每个字符串都只能包含中英文、数字、下划线，且长度为1到32个字符。 当属性类型为“PROTOCOL”时，属性值为支持协议类型的字符串数组。
	Value *interface{} `json:"value"`
}

func (o MessageAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessageAttribute struct{}"
	}

	return strings.Join([]string{"MessageAttribute", string(data)}, " ")
}

type MessageAttributeType struct {
	value string
}

type MessageAttributeTypeEnum struct {
	STRING       MessageAttributeType
	STRING_ARRAY MessageAttributeType
	PROTOCOL     MessageAttributeType
}

func GetMessageAttributeTypeEnum() MessageAttributeTypeEnum {
	return MessageAttributeTypeEnum{
		STRING: MessageAttributeType{
			value: "STRING",
		},
		STRING_ARRAY: MessageAttributeType{
			value: "STRING_ARRAY",
		},
		PROTOCOL: MessageAttributeType{
			value: "PROTOCOL",
		},
	}
}

func (c MessageAttributeType) Value() string {
	return c.value
}

func (c MessageAttributeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MessageAttributeType) UnmarshalJSON(b []byte) error {
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
