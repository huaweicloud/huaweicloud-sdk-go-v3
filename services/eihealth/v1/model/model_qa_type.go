package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QaType **参数解释**： 问答类型。 **约束限制**： 不涉及 **取值范围**： * chat_call：普通问答 * tool_call：实验问答 * deepresearch_call：深度探究 **默认取值**： 不涉及
type QaType struct {
	value string
}

type QaTypeEnum struct {
	CHAT_CALL         QaType
	TOOL_CALL         QaType
	DEEPRESEARCH_CALL QaType
}

func GetQaTypeEnum() QaTypeEnum {
	return QaTypeEnum{
		CHAT_CALL: QaType{
			value: "chat_call",
		},
		TOOL_CALL: QaType{
			value: "tool_call",
		},
		DEEPRESEARCH_CALL: QaType{
			value: "deepresearch_call",
		},
	}
}

func (c QaType) Value() string {
	return c.value
}

func (c QaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QaType) UnmarshalJSON(b []byte) error {
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
