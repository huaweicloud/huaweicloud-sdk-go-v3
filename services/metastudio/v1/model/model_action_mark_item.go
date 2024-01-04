package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ActionMarkItem 标注的动作信息。
type ActionMarkItem struct {

	// 选取推理数据预处理视频起始时间。格式：“HH:MM:SS.mmm”。
	ActionStartTime *string `json:"action_start_time,omitempty"`

	// 选取推理数据预处理视频结束时间。格式：“HH:MM:SS.mmm”。
	ActionEndTime *string `json:"action_end_time,omitempty"`

	// 动作类型。 * SILENCE: 静默 * ACTION：动作
	ActionType *ActionMarkItemActionType `json:"action_type,omitempty"`
}

func (o ActionMarkItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionMarkItem struct{}"
	}

	return strings.Join([]string{"ActionMarkItem", string(data)}, " ")
}

type ActionMarkItemActionType struct {
	value string
}

type ActionMarkItemActionTypeEnum struct {
	SILENCE ActionMarkItemActionType
	ACTION  ActionMarkItemActionType
}

func GetActionMarkItemActionTypeEnum() ActionMarkItemActionTypeEnum {
	return ActionMarkItemActionTypeEnum{
		SILENCE: ActionMarkItemActionType{
			value: "SILENCE",
		},
		ACTION: ActionMarkItemActionType{
			value: "ACTION",
		},
	}
}

func (c ActionMarkItemActionType) Value() string {
	return c.value
}

func (c ActionMarkItemActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionMarkItemActionType) UnmarshalJSON(b []byte) error {
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
