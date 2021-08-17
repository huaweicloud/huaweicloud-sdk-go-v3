package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 批量查询任务时返回体重返回的任务信息。
type TasksResponseBody struct {
	// 迁移任务id

	Id *string `json:"id,omitempty"`
	// 任务名称（用户自定义）

	Name *string `json:"name,omitempty"`
	// 任务类型，创建迁移任务时必选，更新迁移任务时可选

	Type *TasksResponseBodyType `json:"type,omitempty"`
	// 操作系统类型，分为WINDOWS和LINUX，创建时必选，更新时可选

	OsType *TasksResponseBodyOsType `json:"os_type,omitempty"`
	// 任务状态

	State *string `json:"state,omitempty"`
	// 预估完成时间

	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty"`
	// 任务创建时间

	CreateDate *int64 `json:"create_date,omitempty"`
	// 进程优先级 0：低 1：标准 2：高

	Priority *int32 `json:"priority,omitempty"`
	// 迁移限速

	SpeedLimit *int32 `json:"speed_limit,omitempty"`
	// 迁移速率，单位：MB/S

	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`
	// 压缩率

	CompressRate *float64 `json:"compress_rate,omitempty"`
	// 迁移完成后是否启动目的端服务器 true：启动 false：停止

	StartTargetServer *bool `json:"start_target_server,omitempty"`
	// 错误信息

	ErrorJson *string `json:"error_json,omitempty"`
	// 任务总耗时

	TotalTime *int64 `json:"total_time,omitempty"`
	// 目的端服务器的IP地址。 公网迁移时请填写弹性IP地址 专线迁移时请填写私有IP地址

	MigrationIp *string `json:"migration_ip,omitempty"`
	// 任务关联的子任务信息

	SubTasks *[]SubTaskAssociatedWithTask `json:"sub_tasks,omitempty"`

	SourceServer *SourceServerAssociatedWithTask `json:"source_server,omitempty"`
	// 迁移项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	TargetServer *TargetServerAssociatedWithTask `json:"target_server,omitempty"`
	// 日志收集状态

	LogCollectStatus *TasksResponseBodyLogCollectStatus `json:"log_collect_status,omitempty"`

	CloneServer *CloneServerBrief `json:"clone_server,omitempty"`
	// 是否同步

	Syncing *bool `json:"syncing,omitempty"`
}

func (o TasksResponseBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TasksResponseBody struct{}"
	}

	return strings.Join([]string{"TasksResponseBody", string(data)}, " ")
}

type TasksResponseBodyType struct {
	value string
}

type TasksResponseBodyTypeEnum struct {
	MIGRATE_FILE  TasksResponseBodyType
	MIGRATE_BLOCK TasksResponseBodyType
}

func GetTasksResponseBodyTypeEnum() TasksResponseBodyTypeEnum {
	return TasksResponseBodyTypeEnum{
		MIGRATE_FILE: TasksResponseBodyType{
			value: "MIGRATE_FILE：文件级迁移",
		},
		MIGRATE_BLOCK: TasksResponseBodyType{
			value: "MIGRATE_BLOCK：块级迁移",
		},
	}
}

func (c TasksResponseBodyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *TasksResponseBodyType) UnmarshalJSON(b []byte) error {
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

type TasksResponseBodyOsType struct {
	value string
}

type TasksResponseBodyOsTypeEnum struct {
	WINDOWS TasksResponseBodyOsType
	LINUX   TasksResponseBodyOsType
}

func GetTasksResponseBodyOsTypeEnum() TasksResponseBodyOsTypeEnum {
	return TasksResponseBodyOsTypeEnum{
		WINDOWS: TasksResponseBodyOsType{
			value: "WINDOWS",
		},
		LINUX: TasksResponseBodyOsType{
			value: "LINUX",
		},
	}
}

func (c TasksResponseBodyOsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *TasksResponseBodyOsType) UnmarshalJSON(b []byte) error {
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

type TasksResponseBodyLogCollectStatus struct {
	value string
}

type TasksResponseBodyLogCollectStatusEnum struct {
	INIT                          TasksResponseBodyLogCollectStatus
	TELL_AGENT_TO_COLLECTAGENT    TasksResponseBodyLogCollectStatus
	WAIT_AGENT_COLLECT_ACKAGENT   TasksResponseBodyLogCollectStatus
	AGENT_COLLECT_FAILAGENT       TasksResponseBodyLogCollectStatus
	AGENT_COLLECT_SUCCESSAGENT    TasksResponseBodyLogCollectStatus
	WAIT_SERVER_COLLECTSERVER     TasksResponseBodyLogCollectStatus
	SERVER_COLLECT_FAILSERVER     TasksResponseBodyLogCollectStatus
	SERVER_COLLECT_SUCCESSSERVER  TasksResponseBodyLogCollectStatus
	TELL_AGENT_RESET_ACLAGENT     TasksResponseBodyLogCollectStatus
	WAIT_AGENT_RESET_ACL_ACKAGENT TasksResponseBodyLogCollectStatus
}

func GetTasksResponseBodyLogCollectStatusEnum() TasksResponseBodyLogCollectStatusEnum {
	return TasksResponseBodyLogCollectStatusEnum{
		INIT: TasksResponseBodyLogCollectStatus{
			value: "INIT:等待搜集状态",
		},
		TELL_AGENT_TO_COLLECTAGENT: TasksResponseBodyLogCollectStatus{
			value: "TELL_AGENT_TO_COLLECT:通知agent搜集日志",
		},
		WAIT_AGENT_COLLECT_ACKAGENT: TasksResponseBodyLogCollectStatus{
			value: "WAIT_AGENT_COLLECT_ACK:等待Agent确认搜集确认",
		},
		AGENT_COLLECT_FAILAGENT: TasksResponseBodyLogCollectStatus{
			value: "AGENT_COLLECT_FAIL:Agent搜集失败",
		},
		AGENT_COLLECT_SUCCESSAGENT: TasksResponseBodyLogCollectStatus{
			value: "AGENT_COLLECT_SUCCESS：Agent搜集成功",
		},
		WAIT_SERVER_COLLECTSERVER: TasksResponseBodyLogCollectStatus{
			value: "WAIT_SERVER_COLLECT：等待Server端日志搜集",
		},
		SERVER_COLLECT_FAILSERVER: TasksResponseBodyLogCollectStatus{
			value: "SERVER_COLLECT_FAIL：Server搜集失败",
		},
		SERVER_COLLECT_SUCCESSSERVER: TasksResponseBodyLogCollectStatus{
			value: "SERVER_COLLECT_SUCCESS：Server搜集成功",
		},
		TELL_AGENT_RESET_ACLAGENT: TasksResponseBodyLogCollectStatus{
			value: "TELL_AGENT_RESET_ACL：通知Agent取消日志授权",
		},
		WAIT_AGENT_RESET_ACL_ACKAGENT: TasksResponseBodyLogCollectStatus{
			value: "WAIT_AGENT_RESET_ACL_ACK：等待Agent确认",
		},
	}
}

func (c TasksResponseBodyLogCollectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *TasksResponseBodyLogCollectStatus) UnmarshalJSON(b []byte) error {
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
