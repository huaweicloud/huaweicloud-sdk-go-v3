package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MessageRole **参数解释**： 问答角色。 **约束限制**： 不涉及 **取值范围**： * user：用户问题的内容 * assistant：助手回答内容 **默认取值**： 不涉及
type MessageRole struct {
	value string
}

type MessageRoleEnum struct {
	USER      MessageRole
	ASSISTANT MessageRole
}

func GetMessageRoleEnum() MessageRoleEnum {
	return MessageRoleEnum{
		USER: MessageRole{
			value: "user",
		},
		ASSISTANT: MessageRole{
			value: "assistant",
		},
	}
}

func (c MessageRole) Value() string {
	return c.value
}

func (c MessageRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MessageRole) UnmarshalJSON(b []byte) error {
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
