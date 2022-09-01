package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateTasksResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *CreateTasksResponseInfoCode `json:"info_code,omitempty" xml:"info_code"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty" xml:"info_description"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 任务状态:   * running - 正在运行   * success - 成功   * canceled - 已取消   * waiting - 正在等待   * failure - 失败
	TaskStatus     *CreateTasksResponseTaskStatus `json:"task_status,omitempty" xml:"task_status"`
	HttpStatusCode int                            `json:"-"`
}

func (o CreateTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTasksResponse struct{}"
	}

	return strings.Join([]string{"CreateTasksResponse", string(data)}, " ")
}

type CreateTasksResponseInfoCode struct {
	value string
}

type CreateTasksResponseInfoCodeEnum struct {
	SUCCESS CreateTasksResponseInfoCode
	FAILURE CreateTasksResponseInfoCode
}

func GetCreateTasksResponseInfoCodeEnum() CreateTasksResponseInfoCodeEnum {
	return CreateTasksResponseInfoCodeEnum{
		SUCCESS: CreateTasksResponseInfoCode{
			value: "success",
		},
		FAILURE: CreateTasksResponseInfoCode{
			value: "failure",
		},
	}
}

func (c CreateTasksResponseInfoCode) Value() string {
	return c.value
}

func (c CreateTasksResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTasksResponseInfoCode) UnmarshalJSON(b []byte) error {
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

type CreateTasksResponseTaskStatus struct {
	value string
}

type CreateTasksResponseTaskStatusEnum struct {
	RUNNING  CreateTasksResponseTaskStatus
	SUCCESS  CreateTasksResponseTaskStatus
	CANCELED CreateTasksResponseTaskStatus
	WAITING  CreateTasksResponseTaskStatus
	FAILURE  CreateTasksResponseTaskStatus
}

func GetCreateTasksResponseTaskStatusEnum() CreateTasksResponseTaskStatusEnum {
	return CreateTasksResponseTaskStatusEnum{
		RUNNING: CreateTasksResponseTaskStatus{
			value: "running",
		},
		SUCCESS: CreateTasksResponseTaskStatus{
			value: "success",
		},
		CANCELED: CreateTasksResponseTaskStatus{
			value: "canceled",
		},
		WAITING: CreateTasksResponseTaskStatus{
			value: "waiting",
		},
		FAILURE: CreateTasksResponseTaskStatus{
			value: "failure",
		},
	}
}

func (c CreateTasksResponseTaskStatus) Value() string {
	return c.value
}

func (c CreateTasksResponseTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTasksResponseTaskStatus) UnmarshalJSON(b []byte) error {
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
