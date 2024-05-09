package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobListResp 任务列表响应体。
type JobListResp struct {

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 任务状态。取值： - CREATING：创建中。 - CREATE_FAILED：创建失败。 - CONFIGURATION：配置中。 - STARTJOBING：启动中。 - WAITING_FOR_START：等待启动中。 - START_JOB_FAILED：任务启动失败。 - FULL_TRANSFER_STARTED：全量迁移中，灾备场景为初始化。 - FULL_TRANSFER_FAILED：全量迁移失败，灾备场景为初始化失败。 - FULL_TRANSFER_COMPLETE：全量迁移完成，灾备场景为初始化完成。 - INCRE_TRANSFER_STARTED：增量迁移中，灾备场景为灾备中。 - INCRE_TRANSFER_FAILED：增量迁移失败，灾备场景为灾备异常。 - RELEASE_RESOURCE_STARTED：结束任务中。 - RELEASE_RESOURCE_FAILED：结束任务失败。 - RELEASE_RESOURCE_COMPLETE：已结束。 - CHANGE_JOB_STARTED：任务变更中。 - CHANGE_JOB_FAILED：任务变更失败。 - CHILD_TRANSFER_STARTING：子任务启动中。 - CHILD_TRANSFER_STARTED：子任务迁移中。 - CHILD_TRANSFER_COMPLETE：子任务迁移完成。 - CHILD_TRANSFER_FAILED：子任务迁移失败。 - RELEASE_CHILD_TRANSFER_STARTED：子任务结束中。 - RELEASE_CHILD_TRANSFER_COMPLETE：子任务已结束。
	Status JobListRespStatus `json:"status"`

	// 任务描述。
	Description string `json:"description"`

	// 任务创建时间。
	CreateTime string `json:"create_time"`

	// 引擎类型。取值： - oracle-to-gaussdbv5：Oracle同步到GaussDB分布式版，实时同步场景使用。 - redis-to-gaussredis：redis同步到GeminiDB Redis，实时迁移场景使用。 - rediscluster-to-gaussredis：redis集群同步到GeminiDB Redis，实时迁移场景使用。
	EngineType JobListRespEngineType `json:"engine_type"`

	// 网络类型。取值： - eip：公网网络。 - vpc：VPC网络，灾备场景不支持选择VPC网络。 - vpn：VPN、专线网络。
	NetType JobListRespNetType `json:"net_type"`

	// 计费模式。取值： - period：包周期。 - on_demand：按需。
	ChargingMode JobListRespChargingMode `json:"charging_mode"`

	// 是否计费。
	BillingTag bool `json:"billing_tag"`

	// 任务方向。取值： - up：入云 ，灾备场景时对应本云为备。 - down：出云，灾备场景时对应本云为主。 - non-dbs：自建。
	JobDirection JobListRespJobDirection `json:"job_direction"`

	// 任务场景。取值： - migration：实时迁移。 - sync：实时同步。 - cloudDataGuard：实时灾备。
	JobType JobListRespJobType `json:"job_type"`

	// 任务模式。取值： - FULL_TRANS ：全量。 - FULL_INCR_TRANS：全量+增量。 - INCR_TRANS：增量。
	TaskType JobListRespTaskType `json:"task_type"`

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 任务模式。取值： - single：单任务。 - sync_child：同步子任务。 - multi_to_single：多对一任务。
	JobMode JobListRespJobMode `json:"job_mode"`

	// 任务角色。取值： - parent：父任务。 - child：子任务。 - master：主任务。 - slave：备任务。
	JobModeRole JobListRespJobModeRole `json:"job_mode_role"`

	// 是否跨AZ任务。
	IsMultiAz bool `json:"is_multi_az"`

	// 任务节点角色。
	NodeRole string `json:"node_role"`

	// 是否新框架。
	NodeNewFramework bool `json:"node_new_framework"`

	JobAction *JobActions `json:"job_action"`

	// 子任务列表信息体。
	Children *[]ChildrenJobListResp `json:"children,omitempty"`
}

func (o JobListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobListResp struct{}"
	}

	return strings.Join([]string{"JobListResp", string(data)}, " ")
}

type JobListRespStatus struct {
	value string
}

type JobListRespStatusEnum struct {
	CREATING                        JobListRespStatus
	CREATE_FAILED                   JobListRespStatus
	CONFIGURATION                   JobListRespStatus
	STARTJOBING                     JobListRespStatus
	WAITING_FOR_START               JobListRespStatus
	START_JOB_FAILED                JobListRespStatus
	FULL_TRANSFER_STARTED           JobListRespStatus
	FULL_TRANSFER_FAILED            JobListRespStatus
	FULL_TRANSFER_COMPLETE          JobListRespStatus
	INCRE_TRANSFER_STARTED          JobListRespStatus
	INCRE_TRANSFER_FAILED           JobListRespStatus
	RELEASE_RESOURCE_STARTED        JobListRespStatus
	RELEASE_RESOURCE_FAILED         JobListRespStatus
	RELEASE_RESOURCE_COMPLETE       JobListRespStatus
	CHANGE_JOB_STARTED              JobListRespStatus
	CHANGE_JOB_FAILED               JobListRespStatus
	CHILD_TRANSFER_STARTING         JobListRespStatus
	CHILD_TRANSFER_STARTED          JobListRespStatus
	CHILD_TRANSFER_COMPLETE         JobListRespStatus
	CHILD_TRANSFER_FAILED           JobListRespStatus
	RELEASE_CHILD_TRANSFER_STARTED  JobListRespStatus
	RELEASE_CHILD_TRANSFER_COMPLETE JobListRespStatus
}

func GetJobListRespStatusEnum() JobListRespStatusEnum {
	return JobListRespStatusEnum{
		CREATING: JobListRespStatus{
			value: "CREATING",
		},
		CREATE_FAILED: JobListRespStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: JobListRespStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: JobListRespStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: JobListRespStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: JobListRespStatus{
			value: "START_JOB_FAILED",
		},
		FULL_TRANSFER_STARTED: JobListRespStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: JobListRespStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: JobListRespStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: JobListRespStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: JobListRespStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: JobListRespStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: JobListRespStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: JobListRespStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		CHANGE_JOB_STARTED: JobListRespStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: JobListRespStatus{
			value: "CHANGE_JOB_FAILED",
		},
		CHILD_TRANSFER_STARTING: JobListRespStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: JobListRespStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: JobListRespStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: JobListRespStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: JobListRespStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: JobListRespStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
	}
}

func (c JobListRespStatus) Value() string {
	return c.value
}

func (c JobListRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobListRespStatus) UnmarshalJSON(b []byte) error {
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

type JobListRespEngineType struct {
	value string
}

type JobListRespEngineTypeEnum struct {
	ORACLE_TO_GAUSSDBV5        JobListRespEngineType
	REDIS_TO_GAUSSREDIS        JobListRespEngineType
	REDISCLUSTER_TO_GAUSSREDIS JobListRespEngineType
}

func GetJobListRespEngineTypeEnum() JobListRespEngineTypeEnum {
	return JobListRespEngineTypeEnum{
		ORACLE_TO_GAUSSDBV5: JobListRespEngineType{
			value: "oracle-to-gaussdbv5",
		},
		REDIS_TO_GAUSSREDIS: JobListRespEngineType{
			value: "redis-to-gaussredis",
		},
		REDISCLUSTER_TO_GAUSSREDIS: JobListRespEngineType{
			value: "rediscluster-to-gaussredis",
		},
	}
}

func (c JobListRespEngineType) Value() string {
	return c.value
}

func (c JobListRespEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobListRespEngineType) UnmarshalJSON(b []byte) error {
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

type JobListRespNetType struct {
	value string
}

type JobListRespNetTypeEnum struct {
	EIP JobListRespNetType
	VPC JobListRespNetType
	VPN JobListRespNetType
}

func GetJobListRespNetTypeEnum() JobListRespNetTypeEnum {
	return JobListRespNetTypeEnum{
		EIP: JobListRespNetType{
			value: "eip",
		},
		VPC: JobListRespNetType{
			value: "vpc",
		},
		VPN: JobListRespNetType{
			value: "vpn",
		},
	}
}

func (c JobListRespNetType) Value() string {
	return c.value
}

func (c JobListRespNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobListRespNetType) UnmarshalJSON(b []byte) error {
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

type JobListRespChargingMode struct {
	value string
}

type JobListRespChargingModeEnum struct {
	PERIOD    JobListRespChargingMode
	ON_DEMAND JobListRespChargingMode
}

func GetJobListRespChargingModeEnum() JobListRespChargingModeEnum {
	return JobListRespChargingModeEnum{
		PERIOD: JobListRespChargingMode{
			value: "period",
		},
		ON_DEMAND: JobListRespChargingMode{
			value: "on_demand",
		},
	}
}

func (c JobListRespChargingMode) Value() string {
	return c.value
}

func (c JobListRespChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobListRespChargingMode) UnmarshalJSON(b []byte) error {
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

type JobListRespJobDirection struct {
	value string
}

type JobListRespJobDirectionEnum struct {
	UP      JobListRespJobDirection
	DOWN    JobListRespJobDirection
	NON_DBS JobListRespJobDirection
}

func GetJobListRespJobDirectionEnum() JobListRespJobDirectionEnum {
	return JobListRespJobDirectionEnum{
		UP: JobListRespJobDirection{
			value: "up",
		},
		DOWN: JobListRespJobDirection{
			value: "down",
		},
		NON_DBS: JobListRespJobDirection{
			value: "non-dbs",
		},
	}
}

func (c JobListRespJobDirection) Value() string {
	return c.value
}

func (c JobListRespJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobListRespJobDirection) UnmarshalJSON(b []byte) error {
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

type JobListRespJobType struct {
	value string
}

type JobListRespJobTypeEnum struct {
	MIGRATION        JobListRespJobType
	SYNC             JobListRespJobType
	CLOUD_DATA_GUARD JobListRespJobType
}

func GetJobListRespJobTypeEnum() JobListRespJobTypeEnum {
	return JobListRespJobTypeEnum{
		MIGRATION: JobListRespJobType{
			value: "migration",
		},
		SYNC: JobListRespJobType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: JobListRespJobType{
			value: "cloudDataGuard",
		},
	}
}

func (c JobListRespJobType) Value() string {
	return c.value
}

func (c JobListRespJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobListRespJobType) UnmarshalJSON(b []byte) error {
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

type JobListRespTaskType struct {
	value string
}

type JobListRespTaskTypeEnum struct {
	FULL_TRANS      JobListRespTaskType
	FULL_INCR_TRANS JobListRespTaskType
	INCR_TRANS      JobListRespTaskType
}

func GetJobListRespTaskTypeEnum() JobListRespTaskTypeEnum {
	return JobListRespTaskTypeEnum{
		FULL_TRANS: JobListRespTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: JobListRespTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: JobListRespTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c JobListRespTaskType) Value() string {
	return c.value
}

func (c JobListRespTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobListRespTaskType) UnmarshalJSON(b []byte) error {
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

type JobListRespJobMode struct {
	value string
}

type JobListRespJobModeEnum struct {
	SINGLE          JobListRespJobMode
	SYNC_CHILD      JobListRespJobMode
	MULTI_TO_SINGLE JobListRespJobMode
}

func GetJobListRespJobModeEnum() JobListRespJobModeEnum {
	return JobListRespJobModeEnum{
		SINGLE: JobListRespJobMode{
			value: "single",
		},
		SYNC_CHILD: JobListRespJobMode{
			value: "sync_child",
		},
		MULTI_TO_SINGLE: JobListRespJobMode{
			value: "multi_to_single",
		},
	}
}

func (c JobListRespJobMode) Value() string {
	return c.value
}

func (c JobListRespJobMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobListRespJobMode) UnmarshalJSON(b []byte) error {
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

type JobListRespJobModeRole struct {
	value string
}

type JobListRespJobModeRoleEnum struct {
	PARENT JobListRespJobModeRole
	CHILD  JobListRespJobModeRole
	MASTER JobListRespJobModeRole
	SLAVE  JobListRespJobModeRole
}

func GetJobListRespJobModeRoleEnum() JobListRespJobModeRoleEnum {
	return JobListRespJobModeRoleEnum{
		PARENT: JobListRespJobModeRole{
			value: "parent",
		},
		CHILD: JobListRespJobModeRole{
			value: "child",
		},
		MASTER: JobListRespJobModeRole{
			value: "master",
		},
		SLAVE: JobListRespJobModeRole{
			value: "slave",
		},
	}
}

func (c JobListRespJobModeRole) Value() string {
	return c.value
}

func (c JobListRespJobModeRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobListRespJobModeRole) UnmarshalJSON(b []byte) error {
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
