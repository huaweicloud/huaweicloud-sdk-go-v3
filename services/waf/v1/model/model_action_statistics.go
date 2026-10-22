package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ActionStatistics 处置动作统计数据
type ActionStatistics struct {

	// 处置动作类型
	ActionType ActionStatisticsActionType `json:"action_type"`

	// 数量
	Count int64 `json:"count"`
}

func (o ActionStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionStatistics struct{}"
	}

	return strings.Join([]string{"ActionStatistics", string(data)}, " ")
}

type ActionStatisticsActionType struct {
	value string
}

type ActionStatisticsActionTypeEnum struct {
	ALLOW       ActionStatisticsActionType
	BLOCK       ActionStatisticsActionType
	OBSERVE     ActionStatisticsActionType
	MASK        ActionStatisticsActionType
	SAFE_ANSWER ActionStatisticsActionType
}

func GetActionStatisticsActionTypeEnum() ActionStatisticsActionTypeEnum {
	return ActionStatisticsActionTypeEnum{
		ALLOW: ActionStatisticsActionType{
			value: "ALLOW",
		},
		BLOCK: ActionStatisticsActionType{
			value: "BLOCK",
		},
		OBSERVE: ActionStatisticsActionType{
			value: "OBSERVE",
		},
		MASK: ActionStatisticsActionType{
			value: "MASK",
		},
		SAFE_ANSWER: ActionStatisticsActionType{
			value: "SAFE_ANSWER",
		},
	}
}

func (c ActionStatisticsActionType) Value() string {
	return c.value
}

func (c ActionStatisticsActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionStatisticsActionType) UnmarshalJSON(b []byte) error {
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
