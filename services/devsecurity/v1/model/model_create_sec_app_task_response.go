package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSecAppTaskResponse Response Object
type CreateSecAppTaskResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *CreateSecAppTaskResponseInfoCode `json:"info_code,omitempty"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务状态:   * WAITING - 等待   * RUNNING - 进行   * SUCCESS - 完成   * FAILURE - 失败   * STOP - 停止   * DELETED - 删除
	TaskStatus     *CreateSecAppTaskResponseTaskStatus `json:"task_status,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o CreateSecAppTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecAppTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateSecAppTaskResponse", string(data)}, " ")
}

type CreateSecAppTaskResponseInfoCode struct {
	value string
}

type CreateSecAppTaskResponseInfoCodeEnum struct {
	SUCCESS CreateSecAppTaskResponseInfoCode
	FAILURE CreateSecAppTaskResponseInfoCode
}

func GetCreateSecAppTaskResponseInfoCodeEnum() CreateSecAppTaskResponseInfoCodeEnum {
	return CreateSecAppTaskResponseInfoCodeEnum{
		SUCCESS: CreateSecAppTaskResponseInfoCode{
			value: "success",
		},
		FAILURE: CreateSecAppTaskResponseInfoCode{
			value: "failure",
		},
	}
}

func (c CreateSecAppTaskResponseInfoCode) Value() string {
	return c.value
}

func (c CreateSecAppTaskResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecAppTaskResponseInfoCode) UnmarshalJSON(b []byte) error {
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

type CreateSecAppTaskResponseTaskStatus struct {
	value string
}

type CreateSecAppTaskResponseTaskStatusEnum struct {
	WAITING CreateSecAppTaskResponseTaskStatus
	RUNNING CreateSecAppTaskResponseTaskStatus
	SUCCESS CreateSecAppTaskResponseTaskStatus
	FAILURE CreateSecAppTaskResponseTaskStatus
	STOP    CreateSecAppTaskResponseTaskStatus
	DELETED CreateSecAppTaskResponseTaskStatus
}

func GetCreateSecAppTaskResponseTaskStatusEnum() CreateSecAppTaskResponseTaskStatusEnum {
	return CreateSecAppTaskResponseTaskStatusEnum{
		WAITING: CreateSecAppTaskResponseTaskStatus{
			value: "WAITING",
		},
		RUNNING: CreateSecAppTaskResponseTaskStatus{
			value: "RUNNING",
		},
		SUCCESS: CreateSecAppTaskResponseTaskStatus{
			value: "SUCCESS",
		},
		FAILURE: CreateSecAppTaskResponseTaskStatus{
			value: "FAILURE",
		},
		STOP: CreateSecAppTaskResponseTaskStatus{
			value: "STOP",
		},
		DELETED: CreateSecAppTaskResponseTaskStatus{
			value: "DELETED",
		},
	}
}

func (c CreateSecAppTaskResponseTaskStatus) Value() string {
	return c.value
}

func (c CreateSecAppTaskResponseTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecAppTaskResponseTaskStatus) UnmarshalJSON(b []byte) error {
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
