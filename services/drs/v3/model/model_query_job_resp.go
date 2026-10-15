package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QueryJobResp struct {

	// **参数解释**：  任务ID，用于唯一标识一个数据复制任务。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：  父任务ID，多任务场景下子任务所归属的父任务标识。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**：  任务名称，用户创建任务时自定义的名称，用于区分和识别不同的数据复制任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  任务当前运行状态。  **约束限制**：  不涉及。  **取值范围**：  - CREATING：创建中 - CREATE_FAILED：创建失败 - CONFIGURATION：配置中 - STARTJOBING：启动中 - WAITING_FOR_START：等待启动中 - START_JOB_FAILED：启动失败 - PAUSING：已暂停 - FULL_TRANSFER_STARTED：全量开始，灾备场景下为初始化 - FULL_TRANSFER_FAILED：全量失败，灾备场景下为初始化失败 - FULL_TRANSFER_COMPLETE：全量完成，灾备场景下为初始化完成 - INCRE_TRANSFER_STARTED：增量开始，灾备场景下为灾备中 - INCRE_TRANSFER_FAILED：增量失败，灾备场景下为灾备异常 - RELEASE_RESOURCE_STARTED：结束任务中 - RELEASE_RESOURCE_FAILED：结束任务失败 - RELEASE_RESOURCE_COMPLETE：已结束 - REBUILD_NODE_STARTED：故障恢复中 - REBUILD_NODE_FAILED：故障恢复失败 - CHANGE_JOB_STARTED：任务变更中 - CHANGE_JOB_FAILED：任务变更失败 - DELETED：已删除 - CHILD_TRANSFER_STARTING：再编辑子任务启动中 - CHILD_TRANSFER_STARTED：再编辑子任务迁移中 - CHILD_TRANSFER_COMPLETE：再编辑子任务迁移完成 - CHILD_TRANSFER_FAILED：再编辑子任务迁移失败 - RELEASE_CHILD_TRANSFER_STARTED：再编辑子任务结束中 - RELEASE_CHILD_TRANSFER_COMPLETE：再编辑子任务已结束 - NODE_UPGRADE_START：升级开始 - NODE_UPGRADE_COMPLETE：升级完成 - NODE_UPGRADE_FAILED：升级失败  **默认取值**：  不涉及
	Status *QueryJobRespStatus `json:"status,omitempty"`

	// **参数解释**：  任务描述信息，用户创建任务时自定义的备注说明。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：  任务的创建时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**：  任务迁移模式，指定数据复制任务的执行方式。  **约束限制**：  不涉及。  **取值范围**：  - FULL_TRANS：全量迁移，仅迁移历史数据 - FULL_INCR_TRANS：全量+增量迁移，先迁移历史数据再持续同步增量 - INCR_TRANS：增量迁移，仅同步数据变更  **默认取值**：  不涉及
	TaskType *QueryJobRespTaskType `json:"task_type,omitempty"`

	SourceEndpoint *Endpoint `json:"source_endpoint,omitempty"`

	DmqEndpoint *Endpoint `json:"dmq_endpoint,omitempty"`

	// **参数解释**：  物理分片源库信息列表，多对一场景下存在多个源库时返回，每个元素为Endpoint结构。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SourceSharding *[]Endpoint `json:"source_sharding,omitempty"`

	TargetEndpoint *Endpoint `json:"target_endpoint,omitempty"`

	// **参数解释**：  任务使用的网络类型。  **约束限制**：  不涉及。  **取值范围**：  - vpn：虚拟专用网络 - vpc：虚拟私有云 - eip：弹性公网IP  **默认取值**：  不涉及
	NetType *QueryJobRespNetType `json:"net_type,omitempty"`

	// **参数解释**：  任务失败原因，任务状态为失败时返回该字段。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	FailedReason *string `json:"failed_reason,omitempty"`

	InstInfo *InstInfo `json:"inst_info,omitempty"`

	// **参数解释**：  任务实际启动时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ActualStartTime *string `json:"actual_start_time,omitempty"`

	// **参数解释**：  全量迁移完成时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	FullTransferCompleteTime *string `json:"full_transfer_complete_time,omitempty"`

	// **参数解释**：  任务最近一次更新时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**：  任务的迁移方向。  **约束限制**：  不涉及。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主 - non-dbs：自建  **默认取值**：  不涉及
	JobDirection *QueryJobRespJobDirection `json:"job_direction,omitempty"`

	// **参数解释**：  迁移任务的业务场景类型。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备  **默认取值**：  不涉及
	DbUseType *QueryJobRespDbUseType `json:"db_use_type,omitempty"`

	// **参数解释**：  任务配置变更后是否需要重启迁移实例才能生效。  **约束限制**：  不涉及。  **取值范围**：  - true：需要重启 - false：不需要重启  **默认取值**：  不涉及
	NeedRestart *bool `json:"need_restart,omitempty"`

	// **参数解释**：  指定目标数据库实例在迁移过程中是否限制为只读状态。  **约束限制**：  不涉及。  **取值范围**：  - true：目标实例限制为只读 - false：目标实例不限制为只读  **默认取值**：  不涉及
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	// **参数解释**：  数据迁移过程中的冲突忽略策略，指定当目标库已存在相同数据时的处理方式。  **约束限制**：  不涉及。  **取值范围**：  - stop：冲突失败，遇到冲突时停止任务 - overwrite：冲突覆盖，用源库数据覆盖目标库已有数据 - ignore：冲突忽略，跳过冲突数据继续迁移  **默认取值**：  不涉及
	ConflictPolicy *QueryJobRespConflictPolicy `json:"conflict_policy,omitempty"`

	// **参数解释**：  过滤DDL策略，指定迁移过程中需要过滤的DDL操作类型。  **约束限制**：  不涉及。  **取值范围**：  - drop_database：过滤drop_database操作 - drop_databasefilter_all：过滤所有DDL操作 - \"\"：不过滤  **默认取值**：  不涉及
	FilterDdlPolicy *string `json:"filter_ddl_policy,omitempty"`

	// **参数解释**：  迁移速度限制信息，详情请参见SpeedLimitInfo对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SpeedLimit *[]SpeedLimitInfo `json:"speed_limit,omitempty"`

	// **参数解释**：  DRS任务的迁移方案，指定数据复制的底层实现方式。  **约束限制**：  不涉及。  **取值范围**：  - Replication：主从复制 - Tungsten：日志解析 - PGBaseBackup：PG备份  **默认取值**：  不涉及
	SchemaType *QueryJobRespSchemaType `json:"schema_type,omitempty"`

	// **参数解释**：  迁移实例节点个数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	NodeNum *int32 `json:"node_num,omitempty"`

	// **参数解释**：  对象选择开关，标识是否已开启对象选择功能。  **约束限制**：  不涉及。  **取值范围**：  - true：开启对象选择 - false：关闭对象选择  **默认取值**：  不涉及
	ObjectSwitch *bool `json:"object_switch,omitempty"`

	// **参数解释**：  主任务ID，多任务场景下子任务所归属的主任务标识。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	MasterJobId *string `json:"master_job_id,omitempty"`

	// **参数解释**：  全量快照模式。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	FullMode *string `json:"full_mode,omitempty"`

	// **参数解释**：  是否迁移结构。  **约束限制**：  不涉及。  **取值范围**：  - true：迁移结构 - false：不迁移结构  **默认取值**：  不涉及
	StructTrans *bool `json:"struct_trans,omitempty"`

	// **参数解释**：  是否迁移索引。  **约束限制**：  不涉及。  **取值范围**：  - true：迁移索引 - false：不迁移索引  **默认取值**：  不涉及
	IndexTrans *bool `json:"index_trans,omitempty"`

	// **参数解释**：  是否使用目标库的用户替换掉definer。  **约束限制**：  不涉及。  **取值范围**：  - true：使用目标库用户替换definer - false：不替换definer  **默认取值**：  不涉及
	ReplaceDefiner *bool `json:"replace_definer,omitempty"`

	// **参数解释**：  是否迁移用户。  **约束限制**：  不涉及。  **取值范围**：  - true：迁移用户 - false：不迁移用户  **默认取值**：  不涉及
	MigrateUser *bool `json:"migrate_user,omitempty"`

	// **参数解释**：  是否库级同步。  **约束限制**：  不涉及。  **取值范围**：  - true：库级同步 - false：非库级同步  **默认取值**：  不涉及
	SyncDatabase *bool `json:"sync_database,omitempty"`

	// **参数解释**：  错误码，当任务报错时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误详细信息，当任务执行失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMessage *string `json:"error_message,omitempty"`

	TargetRootDb *DefaultRootDb `json:"target_root_db,omitempty"`

	// **参数解释**：  迁移实例节点所在的可用区（AZ，Availability Zone）编码。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	AzCode *string `json:"az_code,omitempty"`

	// **参数解释**：  迁移实例节点所在的虚拟私有云（VPC，Virtual Private Cloud）ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：  迁移实例节点所在的子网ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**：  迁移实例节点所在的安全组ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数解释**：  是否多主灾备任务，双主灾备场景下该值为true。  **约束限制**：  不涉及。  **取值范围**：  - true：多主灾备任务 - false：非多主灾备任务  **默认取值**：  不涉及
	MultiWrite *bool `json:"multi_write,omitempty"`

	// **参数解释**：  是否支持IPv6。  **约束限制**：  不涉及。  **取值范围**：  - true：支持IPv6 - false：不支持IPv6  **默认取值**：  不涉及
	SupportIpV6 *bool `json:"support_ip_v6,omitempty"`

	// **参数解释**：  继承的任务ID，Oracle到Kafka链路时使用。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	InheritId *string `json:"inherit_id,omitempty"`

	// **参数解释**：  断点的GTID（Global Transaction Identifier）集合，用于记录增量同步的断点位置。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Gtid *string `json:"gtid,omitempty"`

	AlarmNotify *QuerySmnInfoResp `json:"alarm_notify,omitempty"`

	// **参数解释**：  增量任务启动位点，标识增量同步起始的数据位置。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	IncreStartPosition *string `json:"incre_start_position,omitempty"`

	// **参数解释**：  标识任务是否为跨可用区（AZ）部署的主备任务。  **约束限制**：  不涉及。  **取值范围**：  - true：跨AZ主备任务 - false：非跨AZ任务  **默认取值**：  不涉及
	IsMultiAz *bool `json:"is_multi_az,omitempty"`

	// **参数解释**：  迁移实例节点所在的可用区（AZ）名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	AzName *string `json:"az_name,omitempty"`

	// **参数解释**：  主备任务中主实例所在的可用区（AZ）。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	MasterAz *string `json:"master_az,omitempty"`

	// **参数解释**：  主备任务中备实例所在的可用区（AZ）。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SlaveAz *string `json:"slave_az,omitempty"`

	// **参数解释**：  任务主备角色，标识当前节点在主备任务中的角色。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	NodeRole *string `json:"node_role,omitempty"`

	PeriodOrder *PeriodOrderResp `json:"period_order,omitempty"`

	// **参数解释**：  已同步对象信息列表，每个元素为DatabaseObjectInfo结构。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ObjectInfos *[]DatabaseObjectInfo `json:"object_infos,omitempty"`

	// **参数解释**：  任务创建时的初始方向，表示任务首次创建时配置的迁移方向，不会随灾备倒换而改变。  **约束限制**：  不涉及。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主 - non-dbs：自建  **默认取值**：  不涉及
	OriginalJobDirection *QueryJobRespOriginalJobDirection `json:"original_job_direction,omitempty"`

	DataTransformation *GetDataTransformationResp `json:"data_transformation,omitempty"`

	// **参数解释**：  DRS任务标签列表，每个元素为Tag结构。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Tags *[]Tag `json:"tags,omitempty"`

	// **参数解释**：  指定公网IP的信息列表，每个元素为PublicIpConfig结构。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	PublicIpList *[]PublicIpConfig `json:"public_ip_list,omitempty"`

	// **参数解释**：  公网IP绑定状态，标识是否成功绑定公网IP。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	BindPublicIpState *string `json:"bind_public_ip_state,omitempty"`

	// **参数解释**：  子任务绑定失败信息列表，多任务场景下存在子任务绑定公网IP失败时返回，每个元素为FailedToBindEipChildInfo结构。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Children *[]FailedToBindEipChildInfo `json:"children,omitempty"`

	// **参数解释**：  是否开启云数据库RDS for MySQL或MariaDB的Binlog快速清理。  **约束限制**：  不涉及。  **取值范围**：  - true：开启Binlog快速清理 - false：不开启Binlog快速清理  **默认取值**：  不涉及
	IsOpenFastClean *bool `json:"is_open_fast_clean,omitempty"`

	// **参数解释**：  灾备任务内核方向。非倒换中与job_direction取值相同，倒换中与job_direction取值相反。  **约束限制**：  不涉及。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主  **默认取值**：  不涉及
	JobKernelDirection *QueryJobRespJobKernelDirection `json:"job_kernel_direction,omitempty"`

	// **参数解释**：  同步任务的同步模式，指定数据同步的粒度级别。  **约束限制**：  不涉及。  **取值范围**：  - DB_LEVEL：库级同步 - SCHEMA_LEVEL：schema级同步 - TABLE_LEVEL：表级同步 - FILE_IMPORT：文件导入 - INSTANCE_LEVEL：实例级同步  **默认取值**：  不涉及
	SyncMode *QueryJobRespSyncMode `json:"sync_mode,omitempty"`
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
	PAUSING                         QueryJobRespStatus
	FULL_TRANSFER_STARTED           QueryJobRespStatus
	FULL_TRANSFER_FAILED            QueryJobRespStatus
	FULL_TRANSFER_COMPLETE          QueryJobRespStatus
	INCRE_TRANSFER_STARTED          QueryJobRespStatus
	INCRE_TRANSFER_FAILED           QueryJobRespStatus
	RELEASE_RESOURCE_STARTED        QueryJobRespStatus
	RELEASE_RESOURCE_FAILED         QueryJobRespStatus
	RELEASE_RESOURCE_COMPLETE       QueryJobRespStatus
	REBUILD_NODE_STARTED            QueryJobRespStatus
	REBUILD_NODE_FAILED             QueryJobRespStatus
	CHANGE_JOB_STARTED              QueryJobRespStatus
	CHANGE_JOB_FAILED               QueryJobRespStatus
	DELETED                         QueryJobRespStatus
	CHILD_TRANSFER_STARTING         QueryJobRespStatus
	CHILD_TRANSFER_STARTED          QueryJobRespStatus
	CHILD_TRANSFER_COMPLETE         QueryJobRespStatus
	CHILD_TRANSFER_FAILED           QueryJobRespStatus
	RELEASE_CHILD_TRANSFER_STARTED  QueryJobRespStatus
	RELEASE_CHILD_TRANSFER_COMPLETE QueryJobRespStatus
	NODE_UPGRADE_START              QueryJobRespStatus
	NODE_UPGRADE_COMPLETE           QueryJobRespStatus
	NODE_UPGRADE_FAILED             QueryJobRespStatus
}

func GetQueryJobRespStatusEnum() QueryJobRespStatusEnum {
	return QueryJobRespStatusEnum{
		CREATING: QueryJobRespStatus{
			value: "CREATING",
		},
		CREATE_FAILED: QueryJobRespStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: QueryJobRespStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: QueryJobRespStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: QueryJobRespStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: QueryJobRespStatus{
			value: "START_JOB_FAILED",
		},
		PAUSING: QueryJobRespStatus{
			value: "PAUSING",
		},
		FULL_TRANSFER_STARTED: QueryJobRespStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: QueryJobRespStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: QueryJobRespStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: QueryJobRespStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: QueryJobRespStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: QueryJobRespStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: QueryJobRespStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: QueryJobRespStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		REBUILD_NODE_STARTED: QueryJobRespStatus{
			value: "REBUILD_NODE_STARTED",
		},
		REBUILD_NODE_FAILED: QueryJobRespStatus{
			value: "REBUILD_NODE_FAILED",
		},
		CHANGE_JOB_STARTED: QueryJobRespStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: QueryJobRespStatus{
			value: "CHANGE_JOB_FAILED",
		},
		DELETED: QueryJobRespStatus{
			value: "DELETED",
		},
		CHILD_TRANSFER_STARTING: QueryJobRespStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: QueryJobRespStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: QueryJobRespStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: QueryJobRespStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: QueryJobRespStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: QueryJobRespStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
		NODE_UPGRADE_START: QueryJobRespStatus{
			value: "NODE_UPGRADE_START",
		},
		NODE_UPGRADE_COMPLETE: QueryJobRespStatus{
			value: "NODE_UPGRADE_COMPLETE",
		},
		NODE_UPGRADE_FAILED: QueryJobRespStatus{
			value: "NODE_UPGRADE_FAILED",
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

type QueryJobRespOriginalJobDirection struct {
	value string
}

type QueryJobRespOriginalJobDirectionEnum struct {
	UP      QueryJobRespOriginalJobDirection
	DOWN    QueryJobRespOriginalJobDirection
	NON_DBS QueryJobRespOriginalJobDirection
}

func GetQueryJobRespOriginalJobDirectionEnum() QueryJobRespOriginalJobDirectionEnum {
	return QueryJobRespOriginalJobDirectionEnum{
		UP: QueryJobRespOriginalJobDirection{
			value: "up",
		},
		DOWN: QueryJobRespOriginalJobDirection{
			value: "down",
		},
		NON_DBS: QueryJobRespOriginalJobDirection{
			value: "non-dbs",
		},
	}
}

func (c QueryJobRespOriginalJobDirection) Value() string {
	return c.value
}

func (c QueryJobRespOriginalJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespOriginalJobDirection) UnmarshalJSON(b []byte) error {
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

type QueryJobRespJobKernelDirection struct {
	value string
}

type QueryJobRespJobKernelDirectionEnum struct {
	UP   QueryJobRespJobKernelDirection
	DOWN QueryJobRespJobKernelDirection
}

func GetQueryJobRespJobKernelDirectionEnum() QueryJobRespJobKernelDirectionEnum {
	return QueryJobRespJobKernelDirectionEnum{
		UP: QueryJobRespJobKernelDirection{
			value: "up",
		},
		DOWN: QueryJobRespJobKernelDirection{
			value: "down",
		},
	}
}

func (c QueryJobRespJobKernelDirection) Value() string {
	return c.value
}

func (c QueryJobRespJobKernelDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespJobKernelDirection) UnmarshalJSON(b []byte) error {
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

type QueryJobRespSyncMode struct {
	value string
}

type QueryJobRespSyncModeEnum struct {
	DB_LEVEL       QueryJobRespSyncMode
	SCHEMA_LEVEL   QueryJobRespSyncMode
	TABLE_LEVEL    QueryJobRespSyncMode
	FILE_IMPORT    QueryJobRespSyncMode
	INSTANCE_LEVEL QueryJobRespSyncMode
}

func GetQueryJobRespSyncModeEnum() QueryJobRespSyncModeEnum {
	return QueryJobRespSyncModeEnum{
		DB_LEVEL: QueryJobRespSyncMode{
			value: "DB_LEVEL",
		},
		SCHEMA_LEVEL: QueryJobRespSyncMode{
			value: "SCHEMA_LEVEL",
		},
		TABLE_LEVEL: QueryJobRespSyncMode{
			value: "TABLE_LEVEL",
		},
		FILE_IMPORT: QueryJobRespSyncMode{
			value: "FILE_IMPORT",
		},
		INSTANCE_LEVEL: QueryJobRespSyncMode{
			value: "INSTANCE_LEVEL",
		},
	}
}

func (c QueryJobRespSyncMode) Value() string {
	return c.value
}

func (c QueryJobRespSyncMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespSyncMode) UnmarshalJSON(b []byte) error {
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
