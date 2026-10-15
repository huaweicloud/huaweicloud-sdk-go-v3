package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobListResp 任务列表响应体。
type JobListResp struct {

	// **参数解释**：  任务ID，用于唯一标识一个数据复制任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Id string `json:"id"`

	// **参数解释**：  任务名称，用于标识数据复制任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Name string `json:"name"`

	// **参数解释**：  任务当前运行状态。  **约束限制**：  不涉及。  **取值范围**：  - CREATING：创建中 - CREATE_FAILED：创建失败 - CONFIGURATION：配置中 - STARTJOBING：启动中 - WAITING_FOR_START：等待启动中 - START_JOB_FAILED：启动失败 - FULL_TRANSFER_STARTED：全量开始，灾备场景下为初始化 - FULL_TRANSFER_FAILED：全量失败，灾备场景下为初始化失败 - FULL_TRANSFER_COMPLETE：全量完成，灾备场景下为初始化完成 - INCRE_TRANSFER_STARTED：增量开始，灾备场景下为灾备中 - INCRE_TRANSFER_FAILED：增量失败，灾备场景下为灾备异常 - RELEASE_RESOURCE_STARTED：结束任务中 - RELEASE_RESOURCE_FAILED：结束任务失败 - RELEASE_RESOURCE_COMPLETE：已结束 - CHANGE_JOB_STARTED：任务变更中 - CHANGE_JOB_FAILED：任务变更失败 - CHILD_TRANSFER_STARTING：子任务启动中 - CHILD_TRANSFER_STARTED：子任务迁移中 - CHILD_TRANSFER_COMPLETE：子任务迁移完成 - CHILD_TRANSFER_FAILED：子任务迁移失败 - RELEASE_CHILD_TRANSFER_STARTED：子任务结束中 - RELEASE_CHILD_TRANSFER_COMPLETE：子任务已结束  **默认取值**：  不涉及
	Status JobListRespStatus `json:"status"`

	// **参数解释**：  任务的描述信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Description string `json:"description"`

	// **参数解释**：  任务的创建时间，表示任务首次被创建的时间戳。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	CreateTime string `json:"create_time"`

	// **参数解释**：  DRS任务引擎类型。  **约束限制**：  不涉及。  **取值范围**：  - oracle-to-gaussdbv5：Oracle同步到GaussDB分布式版，实时同步场景使用 - redis-to-gaussredis：Redis迁移到GeminiDB Redis，实时迁移场景使用 - rediscluster-to-gaussredis：Redis集群迁移到GeminiDB Redis，实时迁移场景使用  **默认取值**：  不涉及
	EngineType JobListRespEngineType `json:"engine_type"`

	// **参数解释**：  任务使用的网络类型。  **约束限制**：  不涉及。  **取值范围**：  - eip：公网网络 - vpc：VPC网络，灾备场景不支持选择VPC网络 - vpn：VPN、专线网络  **默认取值**：  不涉及
	NetType JobListRespNetType `json:"net_type"`

	// **参数解释**：  任务的计费模式，指定数据复制任务的付费方式。  **约束限制**：  不涉及。  **取值范围**：  - period：包周期 - on_demand：按需  **默认取值**：  不涉及
	ChargingMode JobListRespChargingMode `json:"charging_mode"`

	// **参数解释**：  任务的计费模式标识，用于标识该任务当前是否处于计费状态。  **约束限制**：  不涉及。  **取值范围**：  - true：计费中 - false：不计费  **默认取值**：  不涉及
	BillingTag bool `json:"billing_tag"`

	// **参数解释**：  任务的迁移方向。  **约束限制**：  不涉及。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主 - non-dbs：自建  **默认取值**：  不涉及
	JobDirection JobListRespJobDirection `json:"job_direction"`

	// **参数解释**：  DRS任务场景类型，用于区分不同的数据复制场景。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备  **默认取值**：  不涉及
	JobType JobListRespJobType `json:"job_type"`

	// **参数解释**：  任务迁移模式，指定数据复制任务的执行方式。  **约束限制**：  不涉及。  **取值范围**：  - FULL_TRANS：全量迁移，仅迁移历史数据 - FULL_INCR_TRANS：全量+增量迁移，先迁移历史数据再持续同步增量 - INCR_TRANS：增量迁移，仅同步数据变更  **默认取值**：  不涉及
	TaskType JobListRespTaskType `json:"task_type"`

	// **参数解释**：  企业项目ID，用于按企业项目过滤查询结果。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释**：  任务模式，指定数据复制任务的执行组织方式。  **约束限制**：  不涉及。  **取值范围**：  - single：单任务 - sync_child：同步子任务 - multi_to_single：多对一任务  **默认取值**：  不涉及
	JobMode JobListRespJobMode `json:"job_mode"`

	// **参数解释**：  任务角色，指定数据复制任务在多任务场景中的角色。  **约束限制**：  不涉及。  **取值范围**：  - parent：父任务 - child：子任务 - master：主任务 - slave：备任务  **默认取值**：  不涉及
	JobModeRole JobListRespJobModeRole `json:"job_mode_role"`

	// **参数解释**：  标识任务是否为跨可用区（AZ）部署的主备任务。  **约束限制**：  不涉及。  **取值范围**：  - true：跨AZ主备任务 - false：非跨AZ任务  **默认取值**：  不涉及
	IsMultiAz bool `json:"is_multi_az"`

	// **参数解释**：  任务节点角色，标识任务在集群部署中的角色。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	NodeRole string `json:"node_role"`

	// **参数解释**：  标识任务是否使用新框架执行。  **约束限制**：  不涉及。  **取值范围**：  - true：使用新框架 - false：使用旧框架  **默认取值**：  不涉及
	NodeNewFramework bool `json:"node_new_framework"`

	JobAction *JobActions `json:"job_action"`

	// **参数解释**：  子任务列表信息体，包含当前任务下的所有子任务详情。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
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
