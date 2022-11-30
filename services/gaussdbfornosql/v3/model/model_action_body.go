package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ActionBody struct {

	// 指定容灾实例数据同步操作。 取值pause，表示暂停容灾实例的数据同步。 取值resume，表示恢复容灾实例的数据同步。
	Action ActionBodyAction `json:"action"`
}

func (o ActionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionBody struct{}"
	}

	return strings.Join([]string{"ActionBody", string(data)}, " ")
}

type ActionBodyAction struct {
	value string
}

type ActionBodyActionEnum struct {
	PAUSE  ActionBodyAction
	RESUME ActionBodyAction
}

func GetActionBodyActionEnum() ActionBodyActionEnum {
	return ActionBodyActionEnum{
		PAUSE: ActionBodyAction{
			value: "pause",
		},
		RESUME: ActionBodyAction{
			value: "resume",
		},
	}
}

func (c ActionBodyAction) Value() string {
	return c.value
}

func (c ActionBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionBodyAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
