package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTaskResponse Response Object
type ShowTaskResponse struct {

	// 任务名称（用户自定义）
	Name *string `json:"name,omitempty"`

	// 任务类型，创建时必选，更新时可选 MIGRATE_FILE:文件级迁移 MIGRATE_BLOCK:块级迁移
	Type *ShowTaskResponseType `json:"type,omitempty"`

	// 操作系统类型，分为WINDOWS和LINUX，创建时必选，更新时可选
	OsType *ShowTaskResponseOsType `json:"os_type,omitempty"`

	// 迁移任务ID
	Id *string `json:"id,omitempty"`

	// 进程优先级  0：低  1：标准（默认）  2：高
	Priority *ShowTaskResponsePriority `json:"priority,omitempty"`

	// 迁移限速
	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	// 目的端服务器的区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 迁移完成后是否启动目的端服务器  true：启动  false：停止
	StartTargetServer *bool `json:"start_target_server,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 目的端服务器的IP地址。  公网迁移时请填写弹性IP地址  专线迁移时请填写私有IP地址
	MigrationIp *string `json:"migration_ip,omitempty"`

	// 目的端服务器的区域名称
	RegionName *string `json:"region_name,omitempty"`

	// 目的端服务器所在项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 目的端服务器所在项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 模板ID
	VmTemplateId *string `json:"vm_template_id,omitempty"`

	SourceServer *SourceServerResponse `json:"source_server,omitempty"`

	TargetServer *TaskTargetServer `json:"target_server,omitempty"`

	// 迁移任务状态 READY: 准备就绪 RUNNING: 迁移中 SYNCING: 同步中 MIGRATE_SUCCESS: 迁移成功 SYNC_SUCCESS: 同步成功 MIGRATE_FAIL: 失败 SYNC_FAIL: 同步失败 ABORTING: 中止中 ABORT: 中止 SKIPPING: 跳过中 DELETING: 删除中 RESETING: 回滚中
	State *ShowTaskResponseState `json:"state,omitempty"`

	// 预估完成时间
	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty"`

	// 连接状态
	Connected *bool `json:"connected,omitempty"`

	// 任务创建时间
	CreateDate *int64 `json:"create_date,omitempty"`

	// 任务开始时间
	StartDate *int64 `json:"start_date,omitempty"`

	// 任务结束时间
	FinishDate *int64 `json:"finish_date,omitempty"`

	// 迁移速率，单位：Mbit/s
	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	// 压缩率
	CompressRate *float64 `json:"compress_rate,omitempty"`

	// 错误信息
	ErrorJson *string `json:"error_json,omitempty"`

	// 任务总耗时
	TotalTime *int64 `json:"total_time,omitempty"`

	// 暂时保留float,兼容现网老版本的SMS-Agent
	FloatIp *string `json:"float_ip,omitempty"`

	// 迁移剩余时间（秒）
	RemainSeconds *int64 `json:"remain_seconds,omitempty"`

	// 目的端的快照ID
	TargetSnapshotId *string `json:"target_snapshot_id,omitempty"`

	CloneServer *CloneServer `json:"clone_server,omitempty"`

	// 任务包含的子任务列表
	SubTasks *[]SubTask `json:"sub_tasks,omitempty"`

	NetworkCheckInfo *NetworkCheckInfoRequestBody `json:"network_check_info,omitempty"`

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
	SubtaskInfo    *string `json:"subtask_info,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResponse", string(data)}, " ")
}

type ShowTaskResponseType struct {
	value string
}

type ShowTaskResponseTypeEnum struct {
	MIGRATE_FILE  ShowTaskResponseType
	MIGRATE_BLOCK ShowTaskResponseType
}

func GetShowTaskResponseTypeEnum() ShowTaskResponseTypeEnum {
	return ShowTaskResponseTypeEnum{
		MIGRATE_FILE: ShowTaskResponseType{
			value: "MIGRATE_FILE",
		},
		MIGRATE_BLOCK: ShowTaskResponseType{
			value: "MIGRATE_BLOCK",
		},
	}
}

func (c ShowTaskResponseType) Value() string {
	return c.value
}

func (c ShowTaskResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponseType) UnmarshalJSON(b []byte) error {
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

type ShowTaskResponseOsType struct {
	value string
}

type ShowTaskResponseOsTypeEnum struct {
	WINDOWS ShowTaskResponseOsType
	LINUX   ShowTaskResponseOsType
}

func GetShowTaskResponseOsTypeEnum() ShowTaskResponseOsTypeEnum {
	return ShowTaskResponseOsTypeEnum{
		WINDOWS: ShowTaskResponseOsType{
			value: "WINDOWS",
		},
		LINUX: ShowTaskResponseOsType{
			value: "LINUX",
		},
	}
}

func (c ShowTaskResponseOsType) Value() string {
	return c.value
}

func (c ShowTaskResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponseOsType) UnmarshalJSON(b []byte) error {
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

type ShowTaskResponsePriority struct {
	value int32
}

type ShowTaskResponsePriorityEnum struct {
	E_0 ShowTaskResponsePriority
	E_1 ShowTaskResponsePriority
	E_2 ShowTaskResponsePriority
}

func GetShowTaskResponsePriorityEnum() ShowTaskResponsePriorityEnum {
	return ShowTaskResponsePriorityEnum{
		E_0: ShowTaskResponsePriority{
			value: 0,
		}, E_1: ShowTaskResponsePriority{
			value: 1,
		}, E_2: ShowTaskResponsePriority{
			value: 2,
		},
	}
}

func (c ShowTaskResponsePriority) Value() int32 {
	return c.value
}

func (c ShowTaskResponsePriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponsePriority) UnmarshalJSON(b []byte) error {
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

type ShowTaskResponseState struct {
	value string
}

type ShowTaskResponseStateEnum struct {
	READY           ShowTaskResponseState
	RUNNING         ShowTaskResponseState
	SYNCING         ShowTaskResponseState
	MIGRATE_SUCCESS ShowTaskResponseState
	SYNC_SUCCESS    ShowTaskResponseState
	MIGRATE_FAIL    ShowTaskResponseState
	SYNC_FAIL       ShowTaskResponseState
	ABORTING        ShowTaskResponseState
	ABORT           ShowTaskResponseState
	SKIPPING        ShowTaskResponseState
	DELETING        ShowTaskResponseState
	RESETING        ShowTaskResponseState
}

func GetShowTaskResponseStateEnum() ShowTaskResponseStateEnum {
	return ShowTaskResponseStateEnum{
		READY: ShowTaskResponseState{
			value: "READY",
		},
		RUNNING: ShowTaskResponseState{
			value: "RUNNING",
		},
		SYNCING: ShowTaskResponseState{
			value: "SYNCING",
		},
		MIGRATE_SUCCESS: ShowTaskResponseState{
			value: "MIGRATE_SUCCESS",
		},
		SYNC_SUCCESS: ShowTaskResponseState{
			value: "SYNC_SUCCESS",
		},
		MIGRATE_FAIL: ShowTaskResponseState{
			value: "MIGRATE_FAIL",
		},
		SYNC_FAIL: ShowTaskResponseState{
			value: "SYNC_FAIL",
		},
		ABORTING: ShowTaskResponseState{
			value: "ABORTING",
		},
		ABORT: ShowTaskResponseState{
			value: "ABORT",
		},
		SKIPPING: ShowTaskResponseState{
			value: "SKIPPING",
		},
		DELETING: ShowTaskResponseState{
			value: "DELETING",
		},
		RESETING: ShowTaskResponseState{
			value: "RESETING",
		},
	}
}

func (c ShowTaskResponseState) Value() string {
	return c.value
}

func (c ShowTaskResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponseState) UnmarshalJSON(b []byte) error {
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
