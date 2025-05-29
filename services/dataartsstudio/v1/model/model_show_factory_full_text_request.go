package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryFullTextRequest Request Object
type ShowFactoryFullTextRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 搜索空间范围: 默认不传参，在全部空间内搜索。 - 当前工作空间ID: 当前工作空间的ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 全局搜索关键字，输入至少两位字符。
	SearchText *string `json:"search_text,omitempty"`

	// 作业类型: （多选）样例: job_type=BATCH 默认为全部。 - BATCH: 批作业 - REAL_TIME: 流作业
	JobType *string `json:"job_type,omitempty"`

	// 脚本类型: （多选）样例: script_type=HIVE,DLI。 默认为全部，不过滤任何类型脚本。 - HIVE: Hive SQL - [DLI: DLI SQL](tag:nohcs) - DWS: DWS SQL - SparkSQL: Spark SQL - SparkPython: Spark Python - FlinkSQL: Flink SQL - [RDS: RDS SQL](tag:nohcs) - PRESTO: Presto SQL - HETUENGINE: HeruEngine - ClickHouse: ClickHouse - IMPALA: Impala SQL - SHELL: Shell - PYTHON: Python
	ScriptType *string `json:"script_type,omitempty"`

	// 节点类型: （多选）节点类型列表。样例: node_type=com.cloud.datacraft.processactivity.ExecuteHiveJob 默认为全部。 - com.cloud.datacraft.processactivity.ExecuteHiveJob: MRS Hive SQL - com.cloud.datacraft.activity.ExecuteSparkSQL: MRS Spark SQL - com.cloud.datacraft.activity.MRSSparkPython: MRS Spark Python - com.cloud.datacraft.processactivity.ExecuteImpalaJob: MRS Impala SQL - [com.cloud.datacraft.activity.DLISQL: DLI SQL](tag:nohcs) - [com.cloud.datacraft.activity.DliFlinkJob: DLI Flink Job](tag:nohcs) - com.cloud.datacraft.processactivity.ExecuteDWSJob: DWS SQL - com.cloud.datacraft.activity.ExecuteQuery: RDS SQL - com.cloud.datacraft.activity.MRSPrestoSQL: MRS Presto SQL - com.cloud.datacraft.processactivity.ExecuteScript: Shell - com.cloud.datacraft.processactivity.ExecutePythonScript: Python - com.cloud.datacraft.processactivity.ExecuteClickHouseJob: ClickHouse - com.cloud.datacraft.processactivity.ExecuteHetuEngineJob: HetuEngine - com.cloud.datacraft.activity.DataMigration: DataMigration
	NodeType *string `json:"node_type,omitempty"`

	// 最新修改: 样例: new_save_or_commit=save 默认为save: 最新保存 - save: 最新保存 - commit: 最新提交
	NewSaveOrCommit *string `json:"new_save_or_commit,omitempty"`

	// 责任人名称: （多选）人员列表或我的节点。样例: owners=dayu_wm 默认不过滤责任人。
	Owners *string `json:"owners,omitempty"`

	// 搜索范围: （多选）样例: doc_types=script 默认为全部。 - node: 开发作业 - script: 脚本
	DocTypes *string `json:"doc_types,omitempty"`

	// 开始时间，配合结束时间参数使用，默认没有时间范围。样例: begin_time=1746633600000
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 结束时间，配合开始时间参数使用，默认没有时间范围。样例: endTime=1746806399999
	EndTime *int64 `json:"end_time,omitempty"`

	// 分页返回结果，指定每页最大记录数，范围[1,100]。样例: limit=10 默认值: 10。
	Limit *int32 `json:"limit,omitempty"`

	// 分页的起始页，取值范围大于等于0。样例: offset=0 默认值: 0。
	Offset *int32 `json:"offset,omitempty"`

	// 是否搜索配置参数部分的内容: 样例: if_query_parameters=false 默认为false: 不搜索配置参数部分的内容 - true: 是 - false: 否
	IfQueryParameters *string `json:"if_query_parameters,omitempty"`

	// 匹配方式: 样例: match_type=0 默认为0: 通用。 - 0: 通用 - 1: 模糊
	MatchType *int32 `json:"match_type,omitempty"`

	// 调度状态: 仅支持作业查找场景，需配置new_save_or_commit参数为commit 默认为全部。 - running: 已调度 - stop: 未调度
	ScheduleState *string `json:"schedule_state,omitempty"`

	// 是否精确搜索: 开启后配合exact_field参数使用。 默认为false: 非精确搜索 - true: 精确搜索 - false: 非精确搜索
	IsExact *string `json:"is_exact,omitempty"`

	// 精确查询的字段, 开启精确搜索时生效: - jobName: 作业名 - scriptName: 脚本名 - jobId: 作业ID - scriptId: 脚本ID
	ExactField *string `json:"exact_field,omitempty"`
}

func (o ShowFactoryFullTextRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryFullTextRequest struct{}"
	}

	return strings.Join([]string{"ShowFactoryFullTextRequest", string(data)}, " ")
}
