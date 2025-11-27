package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TasksResponseBody 批量查询任务时返回体重返回的任务信息。
type TasksResponseBody struct {

	// 迁移任务ID
	Id *string `json:"id,omitempty"`

	// 任务名称（用户自定义）
	Name *string `json:"name,omitempty"`

	// 任务类型，创建时必选，更新时可选 MIGRATE_FILE:文件级迁移 MIGRATE_BLOCK:块级迁移
	Type *TasksResponseBodyType `json:"type,omitempty"`

	// 操作系统类型，分为WINDOWS和LINUX，创建时必选，更新时可选
	OsType *TasksResponseBodyOsType `json:"os_type,omitempty"`

	// 迁移任务状态 READY: 准备就绪 RUNNING: 迁移中 SYNCING: 同步中 MIGRATE_SUCCESS: 迁移成功 SYNC_SUCCESS: 同步成功 MIGRATE_FAIL: 失败 SYNC_FAIL: 同步失败 ABORTING: 中止中 ABORT: 中止 SKIPPING: 跳过中 DELETING: 删除中 RESETING: 回滚中
	State *TasksResponseBodyState `json:"state,omitempty"`

	// 预估完成时间
	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty"`

	// 任务创建时间
	CreateDate *int64 `json:"create_date,omitempty"`

	// 进程优先级 0：低 1：标准 2：高
	Priority *TasksResponseBodyPriority `json:"priority,omitempty"`

	// 迁移限速
	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	// 迁移速率，单位：Mbit/s
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

	// 迁移项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	TargetServer *TargetServerAssociatedWithTask `json:"target_server,omitempty"`

	// 日志收集状态 INIT:就绪 UPLOADING:上传中 UPLOAD_FAIL:上传失败 UPLOADED:已上传
	LogCollectStatus *TasksResponseBodyLogCollectStatus `json:"log_collect_status,omitempty"`

	CloneServer *CloneServerBrief `json:"clone_server,omitempty"`

	// 是否同步
	Syncing *bool `json:"syncing,omitempty"`

	NetworkCheckInfo *NetworkCheckInfoRequestBody `json:"network_check_info,omitempty"`

	// 特殊配置项配置信息
	SpecialConfig *[]ConfigBody `json:"special_config,omitempty"`

	// 主机的CPU使用率，单位是百分比
	TotalCpuUsage *float64 `json:"total_cpu_usage,omitempty"`

	// Agent的CPU使用率，单位是百分比
	AgentCpuUsage *float64 `json:"agent_cpu_usage,omitempty"`

	// 主机的内存使用值，单位是MB
	TotalMemUsage *float64 `json:"total_mem_usage,omitempty"`

	// Agent的内存使用值，单位是MB
	AgentMemUsage *float64 `json:"agent_mem_usage,omitempty"`

	// 主机的磁盘I/O值，单位是Mbit/s
	TotalDiskIo *float64 `json:"total_disk_io,omitempty"`

	// Agent的磁盘I/O值，单位是Mbit/s
	AgentDiskIo *float64 `json:"agent_disk_io,omitempty"`

	// 是否开启迁移演练
	NeedMigrationTest *bool `json:"need_migration_test,omitempty"`

	// 当前子任务及进度
	SubtaskInfo *string `json:"subtask_info,omitempty"`
}

func (o TasksResponseBody) String() string {
	data, err := utils.Marshal(o)
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
			value: "MIGRATE_FILE",
		},
		MIGRATE_BLOCK: TasksResponseBodyType{
			value: "MIGRATE_BLOCK",
		},
	}
}

func (c TasksResponseBodyType) Value() string {
	return c.value
}

func (c TasksResponseBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksResponseBodyType) UnmarshalJSON(b []byte) error {
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

func (c TasksResponseBodyOsType) Value() string {
	return c.value
}

func (c TasksResponseBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksResponseBodyOsType) UnmarshalJSON(b []byte) error {
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

type TasksResponseBodyState struct {
	value string
}

type TasksResponseBodyStateEnum struct {
	READY           TasksResponseBodyState
	RUNNING         TasksResponseBodyState
	SYNCING         TasksResponseBodyState
	MIGRATE_SUCCESS TasksResponseBodyState
	SYNC_SUCCESS    TasksResponseBodyState
	MIGRATE_FAIL    TasksResponseBodyState
	SYNC_FAIL       TasksResponseBodyState
	ABORTING        TasksResponseBodyState
	ABORT           TasksResponseBodyState
	SKIPPING        TasksResponseBodyState
	DELETING        TasksResponseBodyState
	RESETING        TasksResponseBodyState
}

func GetTasksResponseBodyStateEnum() TasksResponseBodyStateEnum {
	return TasksResponseBodyStateEnum{
		READY: TasksResponseBodyState{
			value: "READY",
		},
		RUNNING: TasksResponseBodyState{
			value: "RUNNING",
		},
		SYNCING: TasksResponseBodyState{
			value: "SYNCING",
		},
		MIGRATE_SUCCESS: TasksResponseBodyState{
			value: "MIGRATE_SUCCESS",
		},
		SYNC_SUCCESS: TasksResponseBodyState{
			value: "SYNC_SUCCESS",
		},
		MIGRATE_FAIL: TasksResponseBodyState{
			value: "MIGRATE_FAIL",
		},
		SYNC_FAIL: TasksResponseBodyState{
			value: "SYNC_FAIL",
		},
		ABORTING: TasksResponseBodyState{
			value: "ABORTING",
		},
		ABORT: TasksResponseBodyState{
			value: "ABORT",
		},
		SKIPPING: TasksResponseBodyState{
			value: "SKIPPING",
		},
		DELETING: TasksResponseBodyState{
			value: "DELETING",
		},
		RESETING: TasksResponseBodyState{
			value: "RESETING",
		},
	}
}

func (c TasksResponseBodyState) Value() string {
	return c.value
}

func (c TasksResponseBodyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksResponseBodyState) UnmarshalJSON(b []byte) error {
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

type TasksResponseBodyPriority struct {
	value int32
}

type TasksResponseBodyPriorityEnum struct {
	E_0 TasksResponseBodyPriority
	E_1 TasksResponseBodyPriority
	E_2 TasksResponseBodyPriority
}

func GetTasksResponseBodyPriorityEnum() TasksResponseBodyPriorityEnum {
	return TasksResponseBodyPriorityEnum{
		E_0: TasksResponseBodyPriority{
			value: 0,
		}, E_1: TasksResponseBodyPriority{
			value: 1,
		}, E_2: TasksResponseBodyPriority{
			value: 2,
		},
	}
}

func (c TasksResponseBodyPriority) Value() int32 {
	return c.value
}

func (c TasksResponseBodyPriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksResponseBodyPriority) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type TasksResponseBodyLogCollectStatus struct {
	value string
}

type TasksResponseBodyLogCollectStatusEnum struct {
	INIT        TasksResponseBodyLogCollectStatus
	UPLOADING   TasksResponseBodyLogCollectStatus
	UPLOAD_FAIL TasksResponseBodyLogCollectStatus
	UPLOADED    TasksResponseBodyLogCollectStatus
}

func GetTasksResponseBodyLogCollectStatusEnum() TasksResponseBodyLogCollectStatusEnum {
	return TasksResponseBodyLogCollectStatusEnum{
		INIT: TasksResponseBodyLogCollectStatus{
			value: "INIT",
		},
		UPLOADING: TasksResponseBodyLogCollectStatus{
			value: "UPLOADING",
		},
		UPLOAD_FAIL: TasksResponseBodyLogCollectStatus{
			value: "UPLOAD_FAIL",
		},
		UPLOADED: TasksResponseBodyLogCollectStatus{
			value: "UPLOADED",
		},
	}
}

func (c TasksResponseBodyLogCollectStatus) Value() string {
	return c.value
}

func (c TasksResponseBodyLogCollectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksResponseBodyLogCollectStatus) UnmarshalJSON(b []byte) error {
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
