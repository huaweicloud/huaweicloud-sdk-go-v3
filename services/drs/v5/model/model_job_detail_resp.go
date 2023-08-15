package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobDetailResp 任务详情响应体。
type JobDetailResp struct {

	// 任务ID。
	Id *string `json:"id,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

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

	// 限速信息体。 - 限速：自定义的最大迁移速度，迁移过程中的迁移速度将不会超过该速度。 - 不限速：对迁移速度不进行限制，通常会最大化使用源数据库的出口带宽。该流速模式同时会对源数据库造成读消耗，消耗取决于源数据库的出口带宽。比如：源数据库的出口带宽为100MB/s，假设高速模式使用了80%带宽，则迁移对源数据库将造成80MB/s的读操作IO消耗。
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
}

func (o JobDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetailResp struct{}"
	}

	return strings.Join([]string{"JobDetailResp", string(data)}, " ")
}
