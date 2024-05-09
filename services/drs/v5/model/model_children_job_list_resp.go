package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChildrenJobListResp 子任务列表响应体。
type ChildrenJobListResp struct {

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 任务状态。取值： - CREATING：创建中。 - CREATE_FAILED：创建失败。 - CONFIGURATION：配置中。 - STARTJOBING：启动中。 - WAITING_FOR_START：等待启动中。 - START_JOB_FAILED：任务启动失败。 - FULL_TRANSFER_STARTED：全量迁移中，灾备场景为初始化。 - FULL_TRANSFER_FAILED：全量迁移失败，灾备场景为初始化失败。 - FULL_TRANSFER_COMPLETE：全量迁移完成，灾备场景为初始化完成。 - INCRE_TRANSFER_STARTED：增量迁移中，灾备场景为灾备中。 - INCRE_TRANSFER_FAILED：增量迁移失败，灾备场景为灾备异常。 - RELEASE_RESOURCE_STARTED：结束任务中。 - RELEASE_RESOURCE_FAILED：结束任务失败。 - RELEASE_RESOURCE_COMPLETE：已结束。 - CHANGE_JOB_STARTED：任务变更中。 - CHANGE_JOB_FAILED：任务变更失败。 - CHILD_TRANSFER_STARTING：子任务启动中。 - CHILD_TRANSFER_STARTED：子任务迁移中。 - CHILD_TRANSFER_COMPLETE：子任务迁移完成。 - CHILD_TRANSFER_FAILED：子任务迁移失败。 - RELEASE_CHILD_TRANSFER_STARTED：子任务结束中。 - RELEASE_CHILD_TRANSFER_COMPLETE：子任务已结束。
	Status ChildrenJobListRespStatus `json:"status"`

	// 任务描述。
	Description string `json:"description"`

	// 任务创建时间。
	CreateTime string `json:"create_time"`

	// 引擎类型。取值： - oracle-to-gaussdbv5：Oracle同步到GaussDB分布式版，实时同步场景使用。 - redis-to-gaussredis：Redis同步到GeminiDB Redis，实时迁移场景使用。 - rediscluster-to-gaussredis：Redis集群同步到GeminiDB Redis，实时迁移场景使用。
	EngineType ChildrenJobListRespEngineType `json:"engine_type"`

	// 网络类型。取值： - eip：公网网络。 - vpc：VPC网络，灾备场景不支持选择VPC网络。 - vpn：VPN、专线网络。
	NetType ChildrenJobListRespNetType `json:"net_type"`

	// 计费模式。取值： - period：包周期。 - on_demand：按需。
	ChargingMode ChildrenJobListRespChargingMode `json:"charging_mode"`

	// 是否计费。
	BillingTag bool `json:"billing_tag"`

	// 任务方向。取值： - up：入云 ，灾备场景时对应本云为备。 - down：出云，灾备场景时对应本云为主。 - non-dbs：自建。
	JobDirection ChildrenJobListRespJobDirection `json:"job_direction"`

	// 任务场景。取值： - migration：实时迁移。 - sync：实时同步。 - cloudDataGuard：实时灾备。
	JobType ChildrenJobListRespJobType `json:"job_type"`

	// 任务模式。取值： - FULL_TRANS ：全量。 - FULL_INCR_TRANS：全量+增量。 - INCR_TRANS：增量。
	TaskType ChildrenJobListRespTaskType `json:"task_type"`

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 任务模式。取值： - single：单任务。 - sync_child：同步子任务。 - multi_to_single：多对一任务。
	JobMode ChildrenJobListRespJobMode `json:"job_mode"`

	// 任务角色。取值： - parent：父任务。 - child：子任务。 - master：主任务。 - slave：备任务。
	JobModeRole ChildrenJobListRespJobModeRole `json:"job_mode_role"`

	// 是否跨AZ任务。
	IsMultiAz bool `json:"is_multi_az"`

	// 任务节点角色。
	NodeRole string `json:"node_role"`

	// 是否新框架。
	NodeNewFramework bool `json:"node_new_framework"`

	JobAction *JobActions `json:"job_action"`
}

func (o ChildrenJobListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChildrenJobListResp struct{}"
	}

	return strings.Join([]string{"ChildrenJobListResp", string(data)}, " ")
}

type ChildrenJobListRespStatus struct {
	value string
}

type ChildrenJobListRespStatusEnum struct {
	CREATING                        ChildrenJobListRespStatus
	CREATE_FAILED                   ChildrenJobListRespStatus
	CONFIGURATION                   ChildrenJobListRespStatus
	STARTJOBING                     ChildrenJobListRespStatus
	WAITING_FOR_START               ChildrenJobListRespStatus
	START_JOB_FAILED                ChildrenJobListRespStatus
	FULL_TRANSFER_STARTED           ChildrenJobListRespStatus
	FULL_TRANSFER_FAILED            ChildrenJobListRespStatus
	FULL_TRANSFER_COMPLETE          ChildrenJobListRespStatus
	INCRE_TRANSFER_STARTED          ChildrenJobListRespStatus
	INCRE_TRANSFER_FAILED           ChildrenJobListRespStatus
	RELEASE_RESOURCE_STARTED        ChildrenJobListRespStatus
	RELEASE_RESOURCE_FAILED         ChildrenJobListRespStatus
	RELEASE_RESOURCE_COMPLETE       ChildrenJobListRespStatus
	CHANGE_JOB_STARTED              ChildrenJobListRespStatus
	CHANGE_JOB_FAILED               ChildrenJobListRespStatus
	CHILD_TRANSFER_STARTING         ChildrenJobListRespStatus
	CHILD_TRANSFER_STARTED          ChildrenJobListRespStatus
	CHILD_TRANSFER_COMPLETE         ChildrenJobListRespStatus
	CHILD_TRANSFER_FAILED           ChildrenJobListRespStatus
	RELEASE_CHILD_TRANSFER_STARTED  ChildrenJobListRespStatus
	RELEASE_CHILD_TRANSFER_COMPLETE ChildrenJobListRespStatus
}

func GetChildrenJobListRespStatusEnum() ChildrenJobListRespStatusEnum {
	return ChildrenJobListRespStatusEnum{
		CREATING: ChildrenJobListRespStatus{
			value: "CREATING",
		},
		CREATE_FAILED: ChildrenJobListRespStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: ChildrenJobListRespStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: ChildrenJobListRespStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: ChildrenJobListRespStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: ChildrenJobListRespStatus{
			value: "START_JOB_FAILED",
		},
		FULL_TRANSFER_STARTED: ChildrenJobListRespStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: ChildrenJobListRespStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: ChildrenJobListRespStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: ChildrenJobListRespStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: ChildrenJobListRespStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: ChildrenJobListRespStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: ChildrenJobListRespStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: ChildrenJobListRespStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		CHANGE_JOB_STARTED: ChildrenJobListRespStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: ChildrenJobListRespStatus{
			value: "CHANGE_JOB_FAILED",
		},
		CHILD_TRANSFER_STARTING: ChildrenJobListRespStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: ChildrenJobListRespStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: ChildrenJobListRespStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: ChildrenJobListRespStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: ChildrenJobListRespStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: ChildrenJobListRespStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
	}
}

func (c ChildrenJobListRespStatus) Value() string {
	return c.value
}

func (c ChildrenJobListRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobListRespStatus) UnmarshalJSON(b []byte) error {
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

type ChildrenJobListRespEngineType struct {
	value string
}

type ChildrenJobListRespEngineTypeEnum struct {
	ORACLE_TO_GAUSSDBV5        ChildrenJobListRespEngineType
	REDIS_TO_GAUSSREDIS        ChildrenJobListRespEngineType
	REDISCLUSTER_TO_GAUSSREDIS ChildrenJobListRespEngineType
}

func GetChildrenJobListRespEngineTypeEnum() ChildrenJobListRespEngineTypeEnum {
	return ChildrenJobListRespEngineTypeEnum{
		ORACLE_TO_GAUSSDBV5: ChildrenJobListRespEngineType{
			value: "oracle-to-gaussdbv5",
		},
		REDIS_TO_GAUSSREDIS: ChildrenJobListRespEngineType{
			value: "redis-to-gaussredis",
		},
		REDISCLUSTER_TO_GAUSSREDIS: ChildrenJobListRespEngineType{
			value: "rediscluster-to-gaussredis",
		},
	}
}

func (c ChildrenJobListRespEngineType) Value() string {
	return c.value
}

func (c ChildrenJobListRespEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobListRespEngineType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobListRespNetType struct {
	value string
}

type ChildrenJobListRespNetTypeEnum struct {
	EIP ChildrenJobListRespNetType
	VPC ChildrenJobListRespNetType
	VPN ChildrenJobListRespNetType
}

func GetChildrenJobListRespNetTypeEnum() ChildrenJobListRespNetTypeEnum {
	return ChildrenJobListRespNetTypeEnum{
		EIP: ChildrenJobListRespNetType{
			value: "eip",
		},
		VPC: ChildrenJobListRespNetType{
			value: "vpc",
		},
		VPN: ChildrenJobListRespNetType{
			value: "vpn",
		},
	}
}

func (c ChildrenJobListRespNetType) Value() string {
	return c.value
}

func (c ChildrenJobListRespNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobListRespNetType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobListRespChargingMode struct {
	value string
}

type ChildrenJobListRespChargingModeEnum struct {
	PERIOD    ChildrenJobListRespChargingMode
	ON_DEMAND ChildrenJobListRespChargingMode
}

func GetChildrenJobListRespChargingModeEnum() ChildrenJobListRespChargingModeEnum {
	return ChildrenJobListRespChargingModeEnum{
		PERIOD: ChildrenJobListRespChargingMode{
			value: "period",
		},
		ON_DEMAND: ChildrenJobListRespChargingMode{
			value: "on_demand",
		},
	}
}

func (c ChildrenJobListRespChargingMode) Value() string {
	return c.value
}

func (c ChildrenJobListRespChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobListRespChargingMode) UnmarshalJSON(b []byte) error {
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

type ChildrenJobListRespJobDirection struct {
	value string
}

type ChildrenJobListRespJobDirectionEnum struct {
	UP      ChildrenJobListRespJobDirection
	DOWN    ChildrenJobListRespJobDirection
	NON_DBS ChildrenJobListRespJobDirection
}

func GetChildrenJobListRespJobDirectionEnum() ChildrenJobListRespJobDirectionEnum {
	return ChildrenJobListRespJobDirectionEnum{
		UP: ChildrenJobListRespJobDirection{
			value: "up",
		},
		DOWN: ChildrenJobListRespJobDirection{
			value: "down",
		},
		NON_DBS: ChildrenJobListRespJobDirection{
			value: "non-dbs",
		},
	}
}

func (c ChildrenJobListRespJobDirection) Value() string {
	return c.value
}

func (c ChildrenJobListRespJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobListRespJobDirection) UnmarshalJSON(b []byte) error {
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

type ChildrenJobListRespJobType struct {
	value string
}

type ChildrenJobListRespJobTypeEnum struct {
	MIGRATION        ChildrenJobListRespJobType
	SYNC             ChildrenJobListRespJobType
	CLOUD_DATA_GUARD ChildrenJobListRespJobType
}

func GetChildrenJobListRespJobTypeEnum() ChildrenJobListRespJobTypeEnum {
	return ChildrenJobListRespJobTypeEnum{
		MIGRATION: ChildrenJobListRespJobType{
			value: "migration",
		},
		SYNC: ChildrenJobListRespJobType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ChildrenJobListRespJobType{
			value: "cloudDataGuard",
		},
	}
}

func (c ChildrenJobListRespJobType) Value() string {
	return c.value
}

func (c ChildrenJobListRespJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobListRespJobType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobListRespTaskType struct {
	value string
}

type ChildrenJobListRespTaskTypeEnum struct {
	FULL_TRANS      ChildrenJobListRespTaskType
	FULL_INCR_TRANS ChildrenJobListRespTaskType
	INCR_TRANS      ChildrenJobListRespTaskType
}

func GetChildrenJobListRespTaskTypeEnum() ChildrenJobListRespTaskTypeEnum {
	return ChildrenJobListRespTaskTypeEnum{
		FULL_TRANS: ChildrenJobListRespTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: ChildrenJobListRespTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: ChildrenJobListRespTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c ChildrenJobListRespTaskType) Value() string {
	return c.value
}

func (c ChildrenJobListRespTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobListRespTaskType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobListRespJobMode struct {
	value string
}

type ChildrenJobListRespJobModeEnum struct {
	SINGLE          ChildrenJobListRespJobMode
	SYNC_CHILD      ChildrenJobListRespJobMode
	MULTI_TO_SINGLE ChildrenJobListRespJobMode
}

func GetChildrenJobListRespJobModeEnum() ChildrenJobListRespJobModeEnum {
	return ChildrenJobListRespJobModeEnum{
		SINGLE: ChildrenJobListRespJobMode{
			value: "single",
		},
		SYNC_CHILD: ChildrenJobListRespJobMode{
			value: "sync_child",
		},
		MULTI_TO_SINGLE: ChildrenJobListRespJobMode{
			value: "multi_to_single",
		},
	}
}

func (c ChildrenJobListRespJobMode) Value() string {
	return c.value
}

func (c ChildrenJobListRespJobMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobListRespJobMode) UnmarshalJSON(b []byte) error {
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

type ChildrenJobListRespJobModeRole struct {
	value string
}

type ChildrenJobListRespJobModeRoleEnum struct {
	PARENT ChildrenJobListRespJobModeRole
	CHILD  ChildrenJobListRespJobModeRole
	MASTER ChildrenJobListRespJobModeRole
	SLAVE  ChildrenJobListRespJobModeRole
}

func GetChildrenJobListRespJobModeRoleEnum() ChildrenJobListRespJobModeRoleEnum {
	return ChildrenJobListRespJobModeRoleEnum{
		PARENT: ChildrenJobListRespJobModeRole{
			value: "parent",
		},
		CHILD: ChildrenJobListRespJobModeRole{
			value: "child",
		},
		MASTER: ChildrenJobListRespJobModeRole{
			value: "master",
		},
		SLAVE: ChildrenJobListRespJobModeRole{
			value: "slave",
		},
	}
}

func (c ChildrenJobListRespJobModeRole) Value() string {
	return c.value
}

func (c ChildrenJobListRespJobModeRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobListRespJobModeRole) UnmarshalJSON(b []byte) error {
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
