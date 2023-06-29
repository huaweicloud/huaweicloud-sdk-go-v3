package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobInfo 在线迁移任务列表信息体
type JobInfo struct {

	// 任务id
	Id string `json:"id"`

	// 任务名称
	Name string `json:"name"`

	// 任务状态。 - CREATING：创建中 - CREATE_FAILED：创建失败 - CONFIGURATION：配置中 - STARTJOBING：启动中 - WAITING_FOR_START：等待启动中 - START_JOB_FAILED：启动失败 - PAUSING：已暂停 - FULL_TRANSFER_STARTED：全量开始，灾备场景下为初始化 - FULL_TRANSFER_FAILED：全量失败，灾备场景下为初始化失败 - FULL_TRANSFER_COMPLETE：全量完成，灾备场景下为初始化完成 - INCRE_TRANSFER_STARTED：增量开始，灾备场景下为灾备中 - INCRE_TRANSFER_FAILED：增量失败，灾备场景下为灾备异常 - RELEASE_RESOURCE_STARTED：结束任务中 - RELEASE_RESOURCE_FAILED：结束任务失败 - RELEASE_RESOURCE_COMPLETE：已结束 - REBUILD_NODE_STARTED：故障恢复中 - REBUILD_NODE_FAILED：故障恢复失败 - CHANGE_JOB_STARTED：任务变更中 - CHANGE_JOB_FAILED：任务变更失败 - DELETED：已删除 - CHILD_TRANSFER_STARTING：再编辑子任务启动中 - CHILD_TRANSFER_STARTED：再编辑子任务迁移中 - CHILD_TRANSFER_COMPLETE：再编辑子任务迁移完成 - CHILD_TRANSFER_FAILED：再编辑子任务迁移失败 - RELEASE_CHILD_TRANSFER_STARTED：再编辑子任务结束中 - RELEASE_CHILD_TRANSFER_COMPLETE：再编辑子任务已结束 - NODE_UPGRADE_START：升级开始 - NODE_UPGRADE_COMPLETE：升级完成 - NODE_UPGRADE_FAILED：升级失败
	Status JobInfoStatus `json:"status"`

	// 任务描述
	Description string `json:"description"`

	// 任务创建时间
	CreateTime string `json:"create_time"`

	// 引擎类型
	EngineType JobInfoEngineType `json:"engine_type"`

	// 网络类型
	NetType JobInfoNetType `json:"net_type"`

	// 计费字段
	BillingTag bool `json:"billing_tag"`

	// 迁移方向
	JobDirection JobInfoJobDirection `json:"job_direction"`

	// 迁移场景。 - migration:实时迁移 - sync:实时同步 - cloudDataGuard:实时灾备
	DbUseType JobInfoDbUseType `json:"db_use_type"`

	// 迁移模式。 - FULL_TRANS 全量 - FULL_INCR_TRANS 全量+增量 - INCR_TRANS 增量
	TaskType JobInfoTaskType `json:"task_type"`

	// 子任务信息体
	Children *[]ChildrenJobInfo `json:"children,omitempty"`

	// 是否新框架
	NodeNewFramework bool `json:"node_newFramework"`

	JobAction *JobActionResp `json:"job_action,omitempty"`
}

func (o JobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}

type JobInfoStatus struct {
	value string
}

type JobInfoStatusEnum struct {
	CREATING                        JobInfoStatus
	CREATE_FAILED                   JobInfoStatus
	CONFIGURATION                   JobInfoStatus
	STARTJOBING                     JobInfoStatus
	WAITING_FOR_START               JobInfoStatus
	START_JOB_FAILED                JobInfoStatus
	PAUSING                         JobInfoStatus
	FULL_TRANSFER_STARTED           JobInfoStatus
	FULL_TRANSFER_FAILED            JobInfoStatus
	FULL_TRANSFER_COMPLETE          JobInfoStatus
	INCRE_TRANSFER_STARTED          JobInfoStatus
	INCRE_TRANSFER_FAILED           JobInfoStatus
	RELEASE_RESOURCE_STARTED        JobInfoStatus
	RELEASE_RESOURCE_FAILED         JobInfoStatus
	RELEASE_RESOURCE_COMPLETE       JobInfoStatus
	REBUILD_NODE_STARTED            JobInfoStatus
	REBUILD_NODE_FAILED             JobInfoStatus
	CHANGE_JOB_STARTED              JobInfoStatus
	CHANGE_JOB_FAILED               JobInfoStatus
	DELETED                         JobInfoStatus
	CHILD_TRANSFER_STARTING         JobInfoStatus
	CHILD_TRANSFER_STARTED          JobInfoStatus
	CHILD_TRANSFER_COMPLETE         JobInfoStatus
	CHILD_TRANSFER_FAILED           JobInfoStatus
	RELEASE_CHILD_TRANSFER_STARTED  JobInfoStatus
	RELEASE_CHILD_TRANSFER_COMPLETE JobInfoStatus
	NODE_UPGRADE_START              JobInfoStatus
	NODE_UPGRADE_COMPLETE           JobInfoStatus
	NODE_UPGRADE_FAILED             JobInfoStatus
}

func GetJobInfoStatusEnum() JobInfoStatusEnum {
	return JobInfoStatusEnum{
		CREATING: JobInfoStatus{
			value: "CREATING",
		},
		CREATE_FAILED: JobInfoStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: JobInfoStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: JobInfoStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: JobInfoStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: JobInfoStatus{
			value: "START_JOB_FAILED",
		},
		PAUSING: JobInfoStatus{
			value: "PAUSING",
		},
		FULL_TRANSFER_STARTED: JobInfoStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: JobInfoStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: JobInfoStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: JobInfoStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: JobInfoStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: JobInfoStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: JobInfoStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: JobInfoStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		REBUILD_NODE_STARTED: JobInfoStatus{
			value: "REBUILD_NODE_STARTED",
		},
		REBUILD_NODE_FAILED: JobInfoStatus{
			value: "REBUILD_NODE_FAILED",
		},
		CHANGE_JOB_STARTED: JobInfoStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: JobInfoStatus{
			value: "CHANGE_JOB_FAILED",
		},
		DELETED: JobInfoStatus{
			value: "DELETED",
		},
		CHILD_TRANSFER_STARTING: JobInfoStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: JobInfoStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: JobInfoStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: JobInfoStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: JobInfoStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: JobInfoStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
		NODE_UPGRADE_START: JobInfoStatus{
			value: "NODE_UPGRADE_START",
		},
		NODE_UPGRADE_COMPLETE: JobInfoStatus{
			value: "NODE_UPGRADE_COMPLETE",
		},
		NODE_UPGRADE_FAILED: JobInfoStatus{
			value: "NODE_UPGRADE_FAILED",
		},
	}
}

func (c JobInfoStatus) Value() string {
	return c.value
}

func (c JobInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoStatus) UnmarshalJSON(b []byte) error {
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

type JobInfoEngineType struct {
	value string
}

type JobInfoEngineTypeEnum struct {
	CLOUD_DATA_GUARD_CASSANDRA       JobInfoEngineType
	CLOUD_DATA_GUARD_DDM             JobInfoEngineType
	CLOUD_DATA_GUARD_TAURUS_TO_MYSQL JobInfoEngineType
	CLOUD_DATA_GUARD_MYSQL           JobInfoEngineType
	CLOUD_DATA_GUARD_MYSQL_TO_TAURUS JobInfoEngineType
}

func GetJobInfoEngineTypeEnum() JobInfoEngineTypeEnum {
	return JobInfoEngineTypeEnum{
		CLOUD_DATA_GUARD_CASSANDRA: JobInfoEngineType{
			value: "cloudDataGuard-cassandra",
		},
		CLOUD_DATA_GUARD_DDM: JobInfoEngineType{
			value: "cloudDataGuard-ddm",
		},
		CLOUD_DATA_GUARD_TAURUS_TO_MYSQL: JobInfoEngineType{
			value: "cloudDataGuard-taurus-to-mysql",
		},
		CLOUD_DATA_GUARD_MYSQL: JobInfoEngineType{
			value: "cloudDataGuard-mysql",
		},
		CLOUD_DATA_GUARD_MYSQL_TO_TAURUS: JobInfoEngineType{
			value: "cloudDataGuard-mysql-to-taurus",
		},
	}
}

func (c JobInfoEngineType) Value() string {
	return c.value
}

func (c JobInfoEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoEngineType) UnmarshalJSON(b []byte) error {
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

type JobInfoNetType struct {
	value string
}

type JobInfoNetTypeEnum struct {
	VPN JobInfoNetType
	VPC JobInfoNetType
	EIP JobInfoNetType
}

func GetJobInfoNetTypeEnum() JobInfoNetTypeEnum {
	return JobInfoNetTypeEnum{
		VPN: JobInfoNetType{
			value: "vpn",
		},
		VPC: JobInfoNetType{
			value: "vpc",
		},
		EIP: JobInfoNetType{
			value: "eip",
		},
	}
}

func (c JobInfoNetType) Value() string {
	return c.value
}

func (c JobInfoNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoNetType) UnmarshalJSON(b []byte) error {
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

type JobInfoJobDirection struct {
	value string
}

type JobInfoJobDirectionEnum struct {
	UP   JobInfoJobDirection
	DOWN JobInfoJobDirection
}

func GetJobInfoJobDirectionEnum() JobInfoJobDirectionEnum {
	return JobInfoJobDirectionEnum{
		UP: JobInfoJobDirection{
			value: "up",
		},
		DOWN: JobInfoJobDirection{
			value: "down",
		},
	}
}

func (c JobInfoJobDirection) Value() string {
	return c.value
}

func (c JobInfoJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoJobDirection) UnmarshalJSON(b []byte) error {
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

type JobInfoDbUseType struct {
	value string
}

type JobInfoDbUseTypeEnum struct {
	MIGRATION        JobInfoDbUseType
	SYNC             JobInfoDbUseType
	CLOUD_DATA_GUARD JobInfoDbUseType
}

func GetJobInfoDbUseTypeEnum() JobInfoDbUseTypeEnum {
	return JobInfoDbUseTypeEnum{
		MIGRATION: JobInfoDbUseType{
			value: "migration",
		},
		SYNC: JobInfoDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: JobInfoDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c JobInfoDbUseType) Value() string {
	return c.value
}

func (c JobInfoDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoDbUseType) UnmarshalJSON(b []byte) error {
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

type JobInfoTaskType struct {
	value string
}

type JobInfoTaskTypeEnum struct {
	FULL_TRANS      JobInfoTaskType
	FULL_INCR_TRANS JobInfoTaskType
	INCR_TRANS      JobInfoTaskType
}

func GetJobInfoTaskTypeEnum() JobInfoTaskTypeEnum {
	return JobInfoTaskTypeEnum{
		FULL_TRANS: JobInfoTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: JobInfoTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: JobInfoTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c JobInfoTaskType) Value() string {
	return c.value
}

func (c JobInfoTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoTaskType) UnmarshalJSON(b []byte) error {
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
