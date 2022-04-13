package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type TaskInfos struct {
	// 任务名称

	TaskName string `json:"task_name"`
	// 待扫描的目标网址

	Url string `json:"url"`
	// 扫描任务类型:   * normal - 普通任务   * monitor - 监测任务

	TaskType *TaskInfosTaskType `json:"task_type,omitempty"`
}

func (o TaskInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInfos struct{}"
	}

	return strings.Join([]string{"TaskInfos", string(data)}, " ")
}

type TaskInfosTaskType struct {
	value string
}

type TaskInfosTaskTypeEnum struct {
	NORMAL  TaskInfosTaskType
	MONITOR TaskInfosTaskType
}

func GetTaskInfosTaskTypeEnum() TaskInfosTaskTypeEnum {
	return TaskInfosTaskTypeEnum{
		NORMAL: TaskInfosTaskType{
			value: "normal",
		},
		MONITOR: TaskInfosTaskType{
			value: "monitor",
		},
	}
}

func (c TaskInfosTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskInfosTaskType) UnmarshalJSON(b []byte) error {
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
