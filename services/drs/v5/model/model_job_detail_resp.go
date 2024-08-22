package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobDetailResp 任务详情响应体。
type JobDetailResp struct {

	// 任务ID。
	Id *string `json:"id,omitempty"`

	// 任务状态。 - CREATING：创建中 - CREATE_FAILED：创建失败 - CONFIGURATION：配置中 - STARTJOBING：启动中 - WAITING_FOR_START：等待启动中 - START_JOB_FAILED：启动失败 - PAUSING：已暂停 - FULL_TRANSFER_STARTED：全量开始，灾备场景下为初始化 - FULL_TRANSFER_FAILED：全量失败，灾备场景下为初始化失败 - FULL_TRANSFER_COMPLETE：全量完成，灾备场景下为初始化完成 - INCRE_TRANSFER_STARTED：增量开始，灾备场景下为灾备中 - INCRE_TRANSFER_FAILED：增量失败，灾备场景下为灾备异常 - RELEASE_RESOURCE_STARTED：结束任务中 - RELEASE_RESOURCE_FAILED：结束任务失败 - RELEASE_RESOURCE_COMPLETE：已结束 - REBUILD_NODE_STARTED：故障恢复中 - REBUILD_NODE_FAILED：故障恢复失败 - CHANGE_JOB_STARTED：任务变更中 - CHANGE_JOB_FAILED：任务变更失败 - DELETED：已删除 - CHILD_TRANSFER_STARTING：再编辑子任务启动中 - CHILD_TRANSFER_STARTED：再编辑子任务迁移中 - CHILD_TRANSFER_COMPLETE：再编辑子任务迁移完成 - CHILD_TRANSFER_FAILED：再编辑子任务迁移失败 - RELEASE_CHILD_TRANSFER_STARTED：再编辑子任务结束中 - RELEASE_CHILD_TRANSFER_COMPLETE：再编辑子任务已结束 - NODE_UPGRADE_START：升级开始 - NODE_UPGRADE_COMPLETE：升级完成 - NODE_UPGRADE_FAILED：升级失败
	Status *JobDetailRespStatus `json:"status,omitempty"`

	// 任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 列表中的项目总数，与分页无关。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 多任务主节点的任务ID。
	MasterJobId *string `json:"master_job_id,omitempty"`

	BaseInfo *JobBaseInfo `json:"base_info,omitempty"`

	// 任务源数据库信息体。
	SourceEndpoint *[]JobEndpointInfo `json:"source_endpoint,omitempty"`

	// 任务目标数据库信息体。
	TargetEndpoint *[]JobEndpointInfo `json:"target_endpoint,omitempty"`

	AlarmNotify *AlarmNotifyConfig `json:"alarm_notify,omitempty"`

	// 限速信息体。 - 限速：自定义的最大迁移速度，迁移过程中的迁移速度将不会超过该速度。  - 不限速：对迁移速度不进行限制，通常会最大化使用源数据库的出口带宽。该流速模式同时会对源数据库造成读消耗，消耗取决于源数据库的出口带宽。比如：源数据库的出口带宽为100MB/s，假设高速模式使用了80%带宽，则迁移对源数据库将造成80MB/s的读操作IO消耗。
	SpeedLimit *[]SpeedLimitInfo `json:"speed_limit,omitempty"`

	UserMigration *UserMigrationInfo `json:"user_migration,omitempty"`

	PolicyConfig *PolicyConfig `json:"policy_config,omitempty"`

	DbParam *DbParamInfo `json:"db_param,omitempty"`

	TuningParams *TuningParamInfo `json:"tuning_params,omitempty"`

	PeriodOrder *PeriodOrderInfo `json:"period_order,omitempty"`

	NodeInfo *JobNodeInfo `json:"node_info,omitempty"`

	// 日志查询结果信息体。
	Logs *[]TaskLogInfo `json:"logs,omitempty"`

	// 网络打通测试结果信息体。
	NetworkResults *[]QueryNetworkResult `json:"network_results,omitempty"`

	PrecheckResult *QueryPreCheckResult `json:"precheck_result,omitempty"`

	ProgressInfo *JobProgressInfo `json:"progress_info,omitempty"`

	MigrationObjectProgressInfo *QueryMigrationObjectProgressInfo `json:"migration_object_progress_info,omitempty"`

	Metrics *QueryMetricResult `json:"metrics,omitempty"`

	CompareResult *CompareResultInfo `json:"compare_result,omitempty"`

	SupportImportFileResp *SupportImportFileResult `json:"support_import_file_resp,omitempty"`

	// 由开关和版本共同控制的任务级别的功能列表。
	InstanceFeatures map[string]string `json:"instance_features,omitempty"`

	// 任务版本。
	TaskVersion *string `json:"task_version,omitempty"`

	ConnectionManagement *ConnectionManagement `json:"connection_management,omitempty"`

	// 指定公网IP的信息
	PublicIpList *[]PublicIpConfig `json:"public_ip_list,omitempty"`

	// 是否成功绑定公网IP
	BindPublicIpState *string `json:"bind_public_ip_state,omitempty"`

	// 多任务时，存在子任务绑定失败时，返回子任务的信息
	Children *[]FailedToBindEipChildInfo `json:"children,omitempty"`

	// 解除目标库只读操作后，目标库解除只读是否成功。 - pending：目标库解除操作进行中。 - success：目标库解除只读操作成功。
	IsWritable *JobDetailRespIsWritable `json:"is_writable,omitempty"`
}

func (o JobDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetailResp struct{}"
	}

	return strings.Join([]string{"JobDetailResp", string(data)}, " ")
}

type JobDetailRespStatus struct {
	value string
}

type JobDetailRespStatusEnum struct {
	CREATING                        JobDetailRespStatus
	CREATE_FAILED                   JobDetailRespStatus
	CONFIGURATION                   JobDetailRespStatus
	STARTJOBING                     JobDetailRespStatus
	WAITING_FOR_START               JobDetailRespStatus
	START_JOB_FAILED                JobDetailRespStatus
	PAUSING                         JobDetailRespStatus
	FULL_TRANSFER_STARTED           JobDetailRespStatus
	FULL_TRANSFER_FAILED            JobDetailRespStatus
	FULL_TRANSFER_COMPLETE          JobDetailRespStatus
	INCRE_TRANSFER_STARTED          JobDetailRespStatus
	INCRE_TRANSFER_FAILED           JobDetailRespStatus
	RELEASE_RESOURCE_STARTED        JobDetailRespStatus
	RELEASE_RESOURCE_FAILED         JobDetailRespStatus
	RELEASE_RESOURCE_COMPLETE       JobDetailRespStatus
	REBUILD_NODE_STARTED            JobDetailRespStatus
	REBUILD_NODE_FAILED             JobDetailRespStatus
	CHANGE_JOB_STARTED              JobDetailRespStatus
	CHANGE_JOB_FAILED               JobDetailRespStatus
	DELETED                         JobDetailRespStatus
	CHILD_TRANSFER_STARTING         JobDetailRespStatus
	CHILD_TRANSFER_STARTED          JobDetailRespStatus
	CHILD_TRANSFER_COMPLETE         JobDetailRespStatus
	CHILD_TRANSFER_FAILED           JobDetailRespStatus
	RELEASE_CHILD_TRANSFER_STARTED  JobDetailRespStatus
	RELEASE_CHILD_TRANSFER_COMPLETE JobDetailRespStatus
	NODE_UPGRADE_START              JobDetailRespStatus
	NODE_UPGRADE_COMPLETE           JobDetailRespStatus
	NODE_UPGRADE_FAILED             JobDetailRespStatus
}

func GetJobDetailRespStatusEnum() JobDetailRespStatusEnum {
	return JobDetailRespStatusEnum{
		CREATING: JobDetailRespStatus{
			value: "CREATING",
		},
		CREATE_FAILED: JobDetailRespStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: JobDetailRespStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: JobDetailRespStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: JobDetailRespStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: JobDetailRespStatus{
			value: "START_JOB_FAILED",
		},
		PAUSING: JobDetailRespStatus{
			value: "PAUSING",
		},
		FULL_TRANSFER_STARTED: JobDetailRespStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: JobDetailRespStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: JobDetailRespStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: JobDetailRespStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: JobDetailRespStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: JobDetailRespStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: JobDetailRespStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: JobDetailRespStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		REBUILD_NODE_STARTED: JobDetailRespStatus{
			value: "REBUILD_NODE_STARTED",
		},
		REBUILD_NODE_FAILED: JobDetailRespStatus{
			value: "REBUILD_NODE_FAILED",
		},
		CHANGE_JOB_STARTED: JobDetailRespStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: JobDetailRespStatus{
			value: "CHANGE_JOB_FAILED",
		},
		DELETED: JobDetailRespStatus{
			value: "DELETED",
		},
		CHILD_TRANSFER_STARTING: JobDetailRespStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: JobDetailRespStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: JobDetailRespStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: JobDetailRespStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: JobDetailRespStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: JobDetailRespStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
		NODE_UPGRADE_START: JobDetailRespStatus{
			value: "NODE_UPGRADE_START",
		},
		NODE_UPGRADE_COMPLETE: JobDetailRespStatus{
			value: "NODE_UPGRADE_COMPLETE",
		},
		NODE_UPGRADE_FAILED: JobDetailRespStatus{
			value: "NODE_UPGRADE_FAILED",
		},
	}
}

func (c JobDetailRespStatus) Value() string {
	return c.value
}

func (c JobDetailRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobDetailRespStatus) UnmarshalJSON(b []byte) error {
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

type JobDetailRespIsWritable struct {
	value string
}

type JobDetailRespIsWritableEnum struct {
	PENDING JobDetailRespIsWritable
	SUCCESS JobDetailRespIsWritable
}

func GetJobDetailRespIsWritableEnum() JobDetailRespIsWritableEnum {
	return JobDetailRespIsWritableEnum{
		PENDING: JobDetailRespIsWritable{
			value: "pending",
		},
		SUCCESS: JobDetailRespIsWritable{
			value: "success",
		},
	}
}

func (c JobDetailRespIsWritable) Value() string {
	return c.value
}

func (c JobDetailRespIsWritable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobDetailRespIsWritable) UnmarshalJSON(b []byte) error {
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
