package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchDetailV2 struct {

	// 唯一标识符
	Id *string `json:"id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// DGC实例ID
	DgcInstanceId *string `json:"dgc_instance_id,omitempty"`

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 搜索范围: （多选） 默认为全部。 - node: 开发作业 - script: 脚本
	DocType *string `json:"doc_type,omitempty"`

	// 负责人
	Owner *string `json:"owner,omitempty"`

	// 最新修改:  默认为save: 最新保存 - save: 最新保存 - commit: 最新提交
	NewSaveOrCommit *string `json:"new_save_or_commit,omitempty"`

	// 数字版本号或草稿标识
	Version *int32 `json:"version,omitempty"`

	// 最后修改时间
	LastModifiedTime *int64 `json:"last_modified_time,omitempty"`

	// 作业名称
	JobName *string `json:"job_name,omitempty"`

	// 作业类型: （多选）样例: job_type=BATCH 默认为全部。 - BATCH: 批作业 - REAL_TIME: 流作业
	JobType *string `json:"job_type,omitempty"`

	// 作业参数
	JobParams *string `json:"job_params,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`

	// 节点类型: （多选）节点类型列表。 默认为全部。   - com.cloud.datacraft.processactivity.ExecuteHiveJob: MRS Hive SQL   - com.cloud.datacraft.activity.ExecuteSparkSQL: MRS Spark SQL   - com.cloud.datacraft.activity.MRSSparkPython: MRS Spark Python   - com.cloud.datacraft.processactivity.ExecuteImpalaJob: MRS Impala SQL   - com.cloud.datacraft.activity.DLISQL: DLI SQL   - com.cloud.datacraft.activity.DliFlinkJob: DLI Flink Job   - com.cloud.datacraft.processactivity.ExecuteDWSJob: DWS SQL   - com.cloud.datacraft.activity.ExecuteQuery: RDS SQL   - com.cloud.datacraft.activity.MRSPrestoSQL: MRS Presto SQL   - com.cloud.datacraft.processactivity.ExecuteScript: Shell   - com.cloud.datacraft.processactivity.ExecutePythonScript: Python   - com.cloud.datacraft.processactivity.ExecuteClickHouseJob: ClickHouse   - com.cloud.datacraft.processactivity.ExecuteHetuEngineJob: HetuEngine   - com.cloud.datacraft.activity.DataMigration: DataMigration
	NodeType *string `json:"node_type,omitempty"`

	// 脚本名称
	ScriptName *string `json:"script_name,omitempty"`

	// 脚本类型: （多选）样例: script_type=HIVE,DLI。 默认为全部，不过滤任何类型脚本。 - HIVE: Hive SQL - DLI: DLI SQL - DWS: DWS SQL - SparkSQL: Spark SQL - Spark Python: Spark Python - FlinkSQL: Flink SQL - RDS: RDS SQL - PRESTO: Presto SQL - HETUENGINE: HeruEngine - ClickHouse: ClickHouse - IMPALA: Impala SQL - SHELL: Shell - PYTHON: Python
	ScriptType *string `json:"script_type,omitempty"`

	// 脚本参数
	ScriptParams *string `json:"script_params,omitempty"`

	// CDM集群名称
	CdmClusterName *string `json:"cdm_cluster_name,omitempty"`

	// CDM作业名称
	CdmJobName *string `json:"cdm_job_name,omitempty"`

	// CDM参数
	CdmParams *string `json:"cdm_params,omitempty"`

	// 工作空间名称
	WorkspaceName *string `json:"workspace_name,omitempty"`

	// 作业ID
	JobId *string `json:"job_id,omitempty"`

	// 脚本ID
	ScriptId *string `json:"script_id,omitempty"`

	// 单节点作业类型
	SingleNodeJobType *string `json:"single_node_job_type,omitempty"`

	// 调度状态:  默认为全部。 - running: 已调度 - stop: 未调度
	ScheduleState *string `json:"schedule_state,omitempty"`
}

func (o SearchDetailV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDetailV2 struct{}"
	}

	return strings.Join([]string{"SearchDetailV2", string(data)}, " ")
}
