package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TaskBasicRsp struct {

	// 任务ID, 可为空
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 任务名称，只能以字母、数字为开头，包含字母、数字和 . _ -  3~100个字符
	TaskName *string `json:"task_name,omitempty" xml:"task_name"`

	// 任务类型 - realtime (实时) - timing (定时)
	TaskType *TaskBasicRspTaskType `json:"task_type,omitempty" xml:"task_type"`

	// 任务状态, - stop (0停止\\未启动) - running (1运行中)
	Status *TaskBasicRspStatus `json:"status,omitempty" xml:"status"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 源端数据源ID
	SourceDatasourceId *string `json:"source_datasource_id,omitempty" xml:"source_datasource_id"`

	// 目标端数据源ID
	TargetDatasourceId *string `json:"target_datasource_id,omitempty" xml:"target_datasource_id"`

	// 源端数据源的名称
	SourceDatasourceName *string `json:"source_datasource_name,omitempty" xml:"source_datasource_name"`

	// 目标端数据源的名称
	TargetDatasourceName *string `json:"target_datasource_name,omitempty" xml:"target_datasource_name"`

	// 源端数据源所属集成应用ID
	SourceAppId *string `json:"source_app_id,omitempty" xml:"source_app_id"`

	// 目标端数据源所属集成应用ID
	TargetAppId *string `json:"target_app_id,omitempty" xml:"target_app_id"`

	// 源端数据源所属集成应用名称
	SourceAppName *string `json:"source_app_name,omitempty" xml:"source_app_name"`

	// 目标端数据源所属集成应用名称
	TargetAppName *string `json:"target_app_name,omitempty" xml:"target_app_name"`

	// 创建时间
	CreatedDate *int64 `json:"created_date,omitempty" xml:"created_date"`

	// 最近一次的修改时间
	LastModifiedDate *int64 `json:"last_modified_date,omitempty" xml:"last_modified_date"`

	// 描述信息
	Description *string `json:"description,omitempty" xml:"description"`

	// 任务标签,只能包含字母、数字、中划线、下划线
	TaskTag *string `json:"task_tag,omitempty" xml:"task_tag"`

	// 任务的创建者
	CreatedBy *string `json:"created_by,omitempty" xml:"created_by"`
}

func (o TaskBasicRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBasicRsp struct{}"
	}

	return strings.Join([]string{"TaskBasicRsp", string(data)}, " ")
}

type TaskBasicRspTaskType struct {
	value string
}

type TaskBasicRspTaskTypeEnum struct {
	REALTIME TaskBasicRspTaskType
	TIMING   TaskBasicRspTaskType
}

func GetTaskBasicRspTaskTypeEnum() TaskBasicRspTaskTypeEnum {
	return TaskBasicRspTaskTypeEnum{
		REALTIME: TaskBasicRspTaskType{
			value: "realtime",
		},
		TIMING: TaskBasicRspTaskType{
			value: "timing",
		},
	}
}

func (c TaskBasicRspTaskType) Value() string {
	return c.value
}

func (c TaskBasicRspTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskBasicRspTaskType) UnmarshalJSON(b []byte) error {
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

type TaskBasicRspStatus struct {
	value string
}

type TaskBasicRspStatusEnum struct {
	STOP    TaskBasicRspStatus
	RUNNING TaskBasicRspStatus
}

func GetTaskBasicRspStatusEnum() TaskBasicRspStatusEnum {
	return TaskBasicRspStatusEnum{
		STOP: TaskBasicRspStatus{
			value: "stop",
		},
		RUNNING: TaskBasicRspStatus{
			value: "running",
		},
	}
}

func (c TaskBasicRspStatus) Value() string {
	return c.value
}

func (c TaskBasicRspStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskBasicRspStatus) UnmarshalJSON(b []byte) error {
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
