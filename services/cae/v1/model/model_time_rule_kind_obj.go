package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TimeRuleKindObj API类型，固定值“TimerRule”，该值不可修改。
type TimeRuleKindObj struct {
	value string
}

type TimeRuleKindObjEnum struct {
	TIMER_RULE TimeRuleKindObj
}

func GetTimeRuleKindObjEnum() TimeRuleKindObjEnum {
	return TimeRuleKindObjEnum{
		TIMER_RULE: TimeRuleKindObj{
			value: "TimerRule",
		},
	}
}

func (c TimeRuleKindObj) Value() string {
	return c.value
}

func (c TimeRuleKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TimeRuleKindObj) UnmarshalJSON(b []byte) error {
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
