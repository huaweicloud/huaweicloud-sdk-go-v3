package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// TaskListItemVo 任务列表项视图
type TaskListItemVo struct {

	// 任务编号
	Id *string `json:"id,omitempty"`

	// 状态
	Status *TaskListItemVoStatus `json:"status,omitempty"`

	// 开始时间（UTC）
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 结束时间（UTC）
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 创建时间（UTC）
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o TaskListItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskListItemVo struct{}"
	}

	return strings.Join([]string{"TaskListItemVo", string(data)}, " ")
}

type TaskListItemVoStatus struct {
	value string
}

type TaskListItemVoStatusEnum struct {
	RUNNING   TaskListItemVoStatus
	FAILED    TaskListItemVoStatus
	SUCCESSED TaskListItemVoStatus
}

func GetTaskListItemVoStatusEnum() TaskListItemVoStatusEnum {
	return TaskListItemVoStatusEnum{
		RUNNING: TaskListItemVoStatus{
			value: "Running",
		},
		FAILED: TaskListItemVoStatus{
			value: "Failed",
		},
		SUCCESSED: TaskListItemVoStatus{
			value: "Successed",
		},
	}
}

func (c TaskListItemVoStatus) Value() string {
	return c.value
}

func (c TaskListItemVoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskListItemVoStatus) UnmarshalJSON(b []byte) error {
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
