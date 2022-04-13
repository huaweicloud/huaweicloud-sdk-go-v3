package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type CancelTasksResponse struct {
	// 状态码:   * success - 成功   * failure - 失败

	InfoCode *CancelTasksResponseInfoCode `json:"info_code,omitempty"`
	// 返回的提示信息

	InfoDescription *string `json:"info_description,omitempty"`
	// 任务ID

	TaskId *string `json:"task_id,omitempty"`
	// 任务状态:   * running - 正在运行   * success - 成功   * canceled - 已取消   * waiting - 正在等待   * failure - 失败

	TaskStatus     *CancelTasksResponseTaskStatus `json:"task_status,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o CancelTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelTasksResponse struct{}"
	}

	return strings.Join([]string{"CancelTasksResponse", string(data)}, " ")
}

type CancelTasksResponseInfoCode struct {
	value string
}

type CancelTasksResponseInfoCodeEnum struct {
	SUCCESS CancelTasksResponseInfoCode
	FAILURE CancelTasksResponseInfoCode
}

func GetCancelTasksResponseInfoCodeEnum() CancelTasksResponseInfoCodeEnum {
	return CancelTasksResponseInfoCodeEnum{
		SUCCESS: CancelTasksResponseInfoCode{
			value: "success",
		},
		FAILURE: CancelTasksResponseInfoCode{
			value: "failure",
		},
	}
}

func (c CancelTasksResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CancelTasksResponseInfoCode) UnmarshalJSON(b []byte) error {
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

type CancelTasksResponseTaskStatus struct {
	value string
}

type CancelTasksResponseTaskStatusEnum struct {
	RUNNING  CancelTasksResponseTaskStatus
	SUCCESS  CancelTasksResponseTaskStatus
	CANCELED CancelTasksResponseTaskStatus
	WAITING  CancelTasksResponseTaskStatus
	FAILURE  CancelTasksResponseTaskStatus
}

func GetCancelTasksResponseTaskStatusEnum() CancelTasksResponseTaskStatusEnum {
	return CancelTasksResponseTaskStatusEnum{
		RUNNING: CancelTasksResponseTaskStatus{
			value: "running",
		},
		SUCCESS: CancelTasksResponseTaskStatus{
			value: "success",
		},
		CANCELED: CancelTasksResponseTaskStatus{
			value: "canceled",
		},
		WAITING: CancelTasksResponseTaskStatus{
			value: "waiting",
		},
		FAILURE: CancelTasksResponseTaskStatus{
			value: "failure",
		},
	}
}

func (c CancelTasksResponseTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CancelTasksResponseTaskStatus) UnmarshalJSON(b []byte) error {
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
