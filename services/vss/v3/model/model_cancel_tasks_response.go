package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CancelTasksResponse Response Object
type CancelTasksResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *CancelTasksResponseInfoCode `json:"info_code,omitempty"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务状态:   * running - 正在运行   * success - 成功   * canceled - 已取消   * waiting - 正在等待   * ready - 已就绪，排队中   * failure - 失败
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

func (c CancelTasksResponseInfoCode) Value() string {
	return c.value
}

func (c CancelTasksResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CancelTasksResponseInfoCode) UnmarshalJSON(b []byte) error {
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

type CancelTasksResponseTaskStatus struct {
	value string
}

type CancelTasksResponseTaskStatusEnum struct {
	RUNNING  CancelTasksResponseTaskStatus
	SUCCESS  CancelTasksResponseTaskStatus
	CANCELED CancelTasksResponseTaskStatus
	WAITING  CancelTasksResponseTaskStatus
	READY    CancelTasksResponseTaskStatus
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
		READY: CancelTasksResponseTaskStatus{
			value: "ready",
		},
		FAILURE: CancelTasksResponseTaskStatus{
			value: "failure",
		},
	}
}

func (c CancelTasksResponseTaskStatus) Value() string {
	return c.value
}

func (c CancelTasksResponseTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CancelTasksResponseTaskStatus) UnmarshalJSON(b []byte) error {
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
