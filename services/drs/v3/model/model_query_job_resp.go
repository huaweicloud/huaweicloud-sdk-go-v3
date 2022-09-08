package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QueryJobResp struct {

	// 任务id
	Id *string `json:"id,omitempty"`

	// 父任务id。
	ParentId *string `json:"parent_id,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 任务状态
	Status *QueryJobRespStatus `json:"status,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 创建时间，时间戳格式。
	CreateTime *string `json:"create_time,omitempty"`

	// 迁移模式
	TaskType *QueryJobRespTaskType `json:"task_type,omitempty"`

	SourceEndpoint *Endpoint `json:"source_endpoint,omitempty"`

	DmqEndpoint *Endpoint `json:"dmq_endpoint,omitempty"`

	// 物理源库信息。
	SourceSharding *[]Endpoint `json:"source_sharding,omitempty"`

	TargetEndpoint *Endpoint `json:"target_endpoint,omitempty"`

	// 网络类型
	NetType *QueryJobRespNetType `json:"net_type,omitempty"`

	// 失败原因。
	FailedReason *string `json:"failed_reason,omitempty"`

	InstInfo *InstInfo `json:"inst_info,omitempty"`

	// 实际启动时间，时间戳格式。
	ActualStartTime *string `json:"actual_start_time,omitempty"`

	// 全量完成时间，时间戳格式。
	FullTransferCompleteTime *string `json:"full_transfer_complete_time,omitempty"`

	// 更新时间，时间戳格式
	UpdateTime *string `json:"update_time,omitempty"`

	// 任务方向
	JobDirection *QueryJobRespJobDirection `json:"job_direction,omitempty"`

	// 迁移场景 - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备
	DbUseType *QueryJobRespDbUseType `json:"db_use_type,omitempty"`

	// 是否需要重启
	NeedRestart *bool `json:"need_restart,omitempty"`

	// 指定目标实例是否限制为只读
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	// 冲突忽略策略 - stop：冲突失败 - overwrite：冲突覆盖 - ignore：冲突忽略
	ConflictPolicy *QueryJobRespConflictPolicy `json:"conflict_policy,omitempty"`

	// 过滤DDL策略 - drop_database：过滤drop_database - drop_databasefilter_all：过滤所有ddl - \"\"：不过滤
	FilterDdlPolicy *string `json:"filter_ddl_policy,omitempty"`

	// 迁移速度限制。
	SpeedLimit *[]SpeedLimitInfo `json:"speed_limit,omitempty"`

	// 迁移方案 - Replication-主从复制 - Tungsten-日志解析 - PGBaseBackup-PG备份
	SchemaType *QueryJobRespSchemaType `json:"schema_type,omitempty"`

	// 节点个数。
	NodeNum *string `json:"node_num,omitempty"`

	// 对象选择开关
	ObjectSwitch *bool `json:"object_switch,omitempty"`

	// 主任务Id。
	MasterJobId *string `json:"master_job_id,omitempty"`

	// 全量快照模式。
	FullMode *string `json:"full_mode,omitempty"`

	// 是否迁移结构。
	StructTrans *bool `json:"struct_trans,omitempty"`

	// 否迁移索引。
	IndexTrans *bool `json:"index_trans,omitempty"`

	// 是否使用目标库的用户替换掉definer。
	ReplaceDefiner *bool `json:"replace_definer,omitempty"`

	// 是否迁移用户。
	MigrateUser *bool `json:"migrate_user,omitempty"`

	// 是否库级同步。
	SyncDatabase *bool `json:"sync_database,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`

	TargetRootDb *DefaultRootDb `json:"target_root_db,omitempty"`

	// node所在AZ
	AzCode *string `json:"az_code,omitempty"`

	// node所在VPC
	VpcId *string `json:"vpc_id,omitempty"`

	// node所在子网
	SubnetId *string `json:"subnet_id,omitempty"`

	// node所在安全组
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 是否多主灾备任务,双主灾备时有值为true
	MultiWrite *bool `json:"multi_write,omitempty"`

	// 是否支持IPV6
	SupportIpV6 *bool `json:"support_ip_v6,omitempty"`

	// 继承的任务ID，Oracle_Mrskafka链路时使用。
	InheritId *string `json:"inherit_id,omitempty"`

	// 断点的GTID集合
	Gtid *string `json:"gtid,omitempty"`

	AlarmNotify *QuerySmnInfoResp `json:"alarm_notify,omitempty"`

	// 增量任务启动位点
	IncreStartPosition *string `json:"incre_start_position,omitempty"`
}

func (o QueryJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryJobResp struct{}"
	}

	return strings.Join([]string{"QueryJobResp", string(data)}, " ")
}

type QueryJobRespStatus struct {
	value string
}

type QueryJobRespStatusEnum struct {
	CREATING                        QueryJobRespStatus
	CREATE_FAILED                   QueryJobRespStatus
	CONFIGURATION                   QueryJobRespStatus
	STARTJOBING                     QueryJobRespStatus
	WAITING_FOR_START               QueryJobRespStatus
	START_JOB_FAILED                QueryJobRespStatus
	FULL_TRANSFER_STARTED           QueryJobRespStatus
	FULL_TRANSFER_FAILED            QueryJobRespStatus
	FULL_TRANSFER_COMPLETE          QueryJobRespStatus
	INCRE_TRANSFER_STARTED          QueryJobRespStatus
	INCRE_TRANSFER_FAILED           QueryJobRespStatus
	RELEASE_RESOURCE_STARTED        QueryJobRespStatus
	RELEASE_RESOURCE_FAILED         QueryJobRespStatus
	RELEASE_RESOURCE_COMPLETE       QueryJobRespStatus
	CHANGE_JOB_STARTED              QueryJobRespStatus
	CHANGE_JOB_FAILED               QueryJobRespStatus
	CHILD_TRANSFER_STARTING         QueryJobRespStatus
	CHILD_TRANSFER_STARTED          QueryJobRespStatus
	CHILD_TRANSFER_COMPLETE         QueryJobRespStatus
	CHILD_TRANSFER_FAILED           QueryJobRespStatus
	RELEASE_CHILD_TRANSFER_STARTED  QueryJobRespStatus
	RELEASE_CHILD_TRANSFER_COMPLETE QueryJobRespStatus
}

func GetQueryJobRespStatusEnum() QueryJobRespStatusEnum {
	return QueryJobRespStatusEnum{
		CREATING: QueryJobRespStatus{
			value: "CREATING：创建中",
		},
		CREATE_FAILED: QueryJobRespStatus{
			value: "CREATE_FAILED: 创建失败",
		},
		CONFIGURATION: QueryJobRespStatus{
			value: "CONFIGURATION: 配置中",
		},
		STARTJOBING: QueryJobRespStatus{
			value: "STARTJOBING: 启动中",
		},
		WAITING_FOR_START: QueryJobRespStatus{
			value: "WAITING_FOR_START：等待启动中",
		},
		START_JOB_FAILED: QueryJobRespStatus{
			value: "START_JOB_FAILED：任务启动失败",
		},
		FULL_TRANSFER_STARTED: QueryJobRespStatus{
			value: "FULL_TRANSFER_STARTED：全量迁移中，灾备场景为初始化 ",
		},
		FULL_TRANSFER_FAILED: QueryJobRespStatus{
			value: "FULL_TRANSFER_FAILED：全量迁移失败，灾备场景为初始化失败 ",
		},
		FULL_TRANSFER_COMPLETE: QueryJobRespStatus{
			value: "FULL_TRANSFER_COMPLETE：全量迁移完成，灾备场景为初始化完成",
		},
		INCRE_TRANSFER_STARTED: QueryJobRespStatus{
			value: " INCRE_TRANSFER_STARTED：增量迁移中，灾备场景为灾备中",
		},
		INCRE_TRANSFER_FAILED: QueryJobRespStatus{
			value: "INCRE_TRANSFER_FAILED：增量迁移失败，灾备场景为灾备异常",
		},
		RELEASE_RESOURCE_STARTED: QueryJobRespStatus{
			value: "RELEASE_RESOURCE_STARTED：结束任务中",
		},
		RELEASE_RESOURCE_FAILED: QueryJobRespStatus{
			value: "RELEASE_RESOURCE_FAILED：结束任务失败",
		},
		RELEASE_RESOURCE_COMPLETE: QueryJobRespStatus{
			value: "RELEASE_RESOURCE_COMPLETE：已结束",
		},
		CHANGE_JOB_STARTED: QueryJobRespStatus{
			value: "CHANGE_JOB_STARTED：任务变更中",
		},
		CHANGE_JOB_FAILED: QueryJobRespStatus{
			value: "CHANGE_JOB_FAILED：任务变更失败",
		},
		CHILD_TRANSFER_STARTING: QueryJobRespStatus{
			value: "CHILD_TRANSFER_STARTING：子任务启动中",
		},
		CHILD_TRANSFER_STARTED: QueryJobRespStatus{
			value: "CHILD_TRANSFER_STARTED：子任务迁移中",
		},
		CHILD_TRANSFER_COMPLETE: QueryJobRespStatus{
			value: "CHILD_TRANSFER_COMPLETE：子任务迁移完成",
		},
		CHILD_TRANSFER_FAILED: QueryJobRespStatus{
			value: "CHILD_TRANSFER_FAILED：子任务迁移失败",
		},
		RELEASE_CHILD_TRANSFER_STARTED: QueryJobRespStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED：子任务结束中",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: QueryJobRespStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE：子任务已结束",
		},
	}
}

func (c QueryJobRespStatus) Value() string {
	return c.value
}

func (c QueryJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespStatus) UnmarshalJSON(b []byte) error {
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

type QueryJobRespTaskType struct {
	value string
}

type QueryJobRespTaskTypeEnum struct {
	FULL_TRANS      QueryJobRespTaskType
	FULL_INCR_TRANS QueryJobRespTaskType
	INCR_TRANS      QueryJobRespTaskType
}

func GetQueryJobRespTaskTypeEnum() QueryJobRespTaskTypeEnum {
	return QueryJobRespTaskTypeEnum{
		FULL_TRANS: QueryJobRespTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: QueryJobRespTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: QueryJobRespTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c QueryJobRespTaskType) Value() string {
	return c.value
}

func (c QueryJobRespTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespTaskType) UnmarshalJSON(b []byte) error {
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

type QueryJobRespNetType struct {
	value string
}

type QueryJobRespNetTypeEnum struct {
	VPN QueryJobRespNetType
	VPC QueryJobRespNetType
	EIP QueryJobRespNetType
}

func GetQueryJobRespNetTypeEnum() QueryJobRespNetTypeEnum {
	return QueryJobRespNetTypeEnum{
		VPN: QueryJobRespNetType{
			value: "vpn",
		},
		VPC: QueryJobRespNetType{
			value: "vpc",
		},
		EIP: QueryJobRespNetType{
			value: "eip",
		},
	}
}

func (c QueryJobRespNetType) Value() string {
	return c.value
}

func (c QueryJobRespNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespNetType) UnmarshalJSON(b []byte) error {
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

type QueryJobRespJobDirection struct {
	value string
}

type QueryJobRespJobDirectionEnum struct {
	UP      QueryJobRespJobDirection
	DOWN    QueryJobRespJobDirection
	NON_DBS QueryJobRespJobDirection
}

func GetQueryJobRespJobDirectionEnum() QueryJobRespJobDirectionEnum {
	return QueryJobRespJobDirectionEnum{
		UP: QueryJobRespJobDirection{
			value: "up",
		},
		DOWN: QueryJobRespJobDirection{
			value: "down",
		},
		NON_DBS: QueryJobRespJobDirection{
			value: "non-dbs",
		},
	}
}

func (c QueryJobRespJobDirection) Value() string {
	return c.value
}

func (c QueryJobRespJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespJobDirection) UnmarshalJSON(b []byte) error {
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

type QueryJobRespDbUseType struct {
	value string
}

type QueryJobRespDbUseTypeEnum struct {
	MIGRATION        QueryJobRespDbUseType
	SYNC             QueryJobRespDbUseType
	CLOUD_DATA_GUARD QueryJobRespDbUseType
}

func GetQueryJobRespDbUseTypeEnum() QueryJobRespDbUseTypeEnum {
	return QueryJobRespDbUseTypeEnum{
		MIGRATION: QueryJobRespDbUseType{
			value: "migration",
		},
		SYNC: QueryJobRespDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: QueryJobRespDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c QueryJobRespDbUseType) Value() string {
	return c.value
}

func (c QueryJobRespDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespDbUseType) UnmarshalJSON(b []byte) error {
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

type QueryJobRespConflictPolicy struct {
	value string
}

type QueryJobRespConflictPolicyEnum struct {
	STOP      QueryJobRespConflictPolicy
	OVERWRITE QueryJobRespConflictPolicy
	IGNORE    QueryJobRespConflictPolicy
}

func GetQueryJobRespConflictPolicyEnum() QueryJobRespConflictPolicyEnum {
	return QueryJobRespConflictPolicyEnum{
		STOP: QueryJobRespConflictPolicy{
			value: "stop",
		},
		OVERWRITE: QueryJobRespConflictPolicy{
			value: "overwrite",
		},
		IGNORE: QueryJobRespConflictPolicy{
			value: "ignore",
		},
	}
}

func (c QueryJobRespConflictPolicy) Value() string {
	return c.value
}

func (c QueryJobRespConflictPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespConflictPolicy) UnmarshalJSON(b []byte) error {
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

type QueryJobRespSchemaType struct {
	value string
}

type QueryJobRespSchemaTypeEnum struct {
	REPLICATION    QueryJobRespSchemaType
	TUNGSTEN       QueryJobRespSchemaType
	PG_BASE_BACKUP QueryJobRespSchemaType
}

func GetQueryJobRespSchemaTypeEnum() QueryJobRespSchemaTypeEnum {
	return QueryJobRespSchemaTypeEnum{
		REPLICATION: QueryJobRespSchemaType{
			value: "Replication",
		},
		TUNGSTEN: QueryJobRespSchemaType{
			value: "Tungsten",
		},
		PG_BASE_BACKUP: QueryJobRespSchemaType{
			value: "PGBaseBackup",
		},
	}
}

func (c QueryJobRespSchemaType) Value() string {
	return c.value
}

func (c QueryJobRespSchemaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespSchemaType) UnmarshalJSON(b []byte) error {
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
