package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangedActionDetail **参数解释**：策略 action 变更详情，包含 action 操作、action 描述和变更类型。
type ChangedActionDetail struct {

	// **参数解释**：权限操作内容。 **取值范围**：长度0~256个字符。
	Action string `json:"action"`

	// **参数解释**：权限操作的描述信息。 **取值范围**：不涉及。
	ActionDescription string `json:"action_description"`

	// **参数解释**：变更类型。added 表示需要新增的权限操作。 **取值范围**：   - added：需要新增
	ChangeType ChangedActionDetailChangeType `json:"change_type"`
}

func (o ChangedActionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangedActionDetail struct{}"
	}

	return strings.Join([]string{"ChangedActionDetail", string(data)}, " ")
}

type ChangedActionDetailChangeType struct {
	value string
}

type ChangedActionDetailChangeTypeEnum struct {
	ADDED ChangedActionDetailChangeType
}

func GetChangedActionDetailChangeTypeEnum() ChangedActionDetailChangeTypeEnum {
	return ChangedActionDetailChangeTypeEnum{
		ADDED: ChangedActionDetailChangeType{
			value: "added",
		},
	}
}

func (c ChangedActionDetailChangeType) Value() string {
	return c.value
}

func (c ChangedActionDetailChangeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangedActionDetailChangeType) UnmarshalJSON(b []byte) error {
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
