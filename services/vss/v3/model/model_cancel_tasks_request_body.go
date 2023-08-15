package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CancelTasksRequestBody struct {

	// 任务ID
	TaskId string `json:"task_id"`

	// 对扫描任务的操作:   * cancel - 取消扫描任务   * restart - 重启扫描任务
	Action CancelTasksRequestBodyAction `json:"action"`
}

func (o CancelTasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelTasksRequestBody struct{}"
	}

	return strings.Join([]string{"CancelTasksRequestBody", string(data)}, " ")
}

type CancelTasksRequestBodyAction struct {
	value string
}

type CancelTasksRequestBodyActionEnum struct {
	CANCEL  CancelTasksRequestBodyAction
	RESTART CancelTasksRequestBodyAction
}

func GetCancelTasksRequestBodyActionEnum() CancelTasksRequestBodyActionEnum {
	return CancelTasksRequestBodyActionEnum{
		CANCEL: CancelTasksRequestBodyAction{
			value: "cancel",
		},
		RESTART: CancelTasksRequestBodyAction{
			value: "restart",
		},
	}
}

func (c CancelTasksRequestBodyAction) Value() string {
	return c.value
}

func (c CancelTasksRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CancelTasksRequestBodyAction) UnmarshalJSON(b []byte) error {
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
