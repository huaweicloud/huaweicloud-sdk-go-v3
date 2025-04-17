package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowNewTaskStatusResponse Response Object
type ShowNewTaskStatusResponse struct {

	// 任务状态，可选范围： - success: 表示成功 - failed: 表示失败 - waiting: 表示等待 - running: 表示运行中 - preprocess: 表示预处理 - ready: 表示准备
	TaskStatus *ShowNewTaskStatusResponseTaskStatus `json:"task_status,omitempty"`

	// 任务的附加信息
	TaskMsg        *string `json:"task_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowNewTaskStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNewTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowNewTaskStatusResponse", string(data)}, " ")
}

type ShowNewTaskStatusResponseTaskStatus struct {
	value string
}

type ShowNewTaskStatusResponseTaskStatusEnum struct {
	SUCCESS    ShowNewTaskStatusResponseTaskStatus
	FAILED     ShowNewTaskStatusResponseTaskStatus
	WAITING    ShowNewTaskStatusResponseTaskStatus
	RUNNING    ShowNewTaskStatusResponseTaskStatus
	PREPROCESS ShowNewTaskStatusResponseTaskStatus
	READY      ShowNewTaskStatusResponseTaskStatus
}

func GetShowNewTaskStatusResponseTaskStatusEnum() ShowNewTaskStatusResponseTaskStatusEnum {
	return ShowNewTaskStatusResponseTaskStatusEnum{
		SUCCESS: ShowNewTaskStatusResponseTaskStatus{
			value: "success",
		},
		FAILED: ShowNewTaskStatusResponseTaskStatus{
			value: "failed",
		},
		WAITING: ShowNewTaskStatusResponseTaskStatus{
			value: "waiting",
		},
		RUNNING: ShowNewTaskStatusResponseTaskStatus{
			value: "running",
		},
		PREPROCESS: ShowNewTaskStatusResponseTaskStatus{
			value: "preprocess",
		},
		READY: ShowNewTaskStatusResponseTaskStatus{
			value: "ready",
		},
	}
}

func (c ShowNewTaskStatusResponseTaskStatus) Value() string {
	return c.value
}

func (c ShowNewTaskStatusResponseTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowNewTaskStatusResponseTaskStatus) UnmarshalJSON(b []byte) error {
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
