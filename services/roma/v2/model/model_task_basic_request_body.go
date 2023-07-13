package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TaskBasicRequestBody struct {

	// 任务名称，只能以字母、数字为开头，包含字母、数字和 . _ -  3~100个字符
	TaskName string `json:"task_name"`

	// 任务类型 - REALTIME (实时) - TIMING (定时)
	TaskType TaskBasicRequestBodyTaskType `json:"task_type"`

	// 源端数据源ID
	SourceDatasourceId string `json:"source_datasource_id"`

	// 目标端数据源ID
	TargetDatasourceId string `json:"target_datasource_id"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 任务标签,只能包含字母、数字、中划线、下划线
	TaskTag *string `json:"task_tag,omitempty"`
}

func (o TaskBasicRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBasicRequestBody struct{}"
	}

	return strings.Join([]string{"TaskBasicRequestBody", string(data)}, " ")
}

type TaskBasicRequestBodyTaskType struct {
	value string
}

type TaskBasicRequestBodyTaskTypeEnum struct {
	REALTIME TaskBasicRequestBodyTaskType
	TIMING   TaskBasicRequestBodyTaskType
}

func GetTaskBasicRequestBodyTaskTypeEnum() TaskBasicRequestBodyTaskTypeEnum {
	return TaskBasicRequestBodyTaskTypeEnum{
		REALTIME: TaskBasicRequestBodyTaskType{
			value: "REALTIME",
		},
		TIMING: TaskBasicRequestBodyTaskType{
			value: "TIMING",
		},
	}
}

func (c TaskBasicRequestBodyTaskType) Value() string {
	return c.value
}

func (c TaskBasicRequestBodyTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskBasicRequestBodyTaskType) UnmarshalJSON(b []byte) error {
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
