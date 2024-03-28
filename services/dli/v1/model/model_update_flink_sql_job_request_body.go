package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkSqlJobRequestBody 更新SQL作业的请求参数。
type UpdateFlinkSqlJobRequestBody struct {

	// 作业名称。长度限制：0-57个字符。
	Name *string `json:"name,omitempty"`

	// 作业描述。长度限制：0-2048个字符。
	Desc *string `json:"desc,omitempty"`

	// 队列名称。长度限制：1-128个字符。
	QueueName *string `json:"queue_name,omitempty"`

	// Stream SQL语句。长度限制：0-1024*1024个字符。
	SqlBody *string `json:"sql_body,omitempty"`

	// 作业运行模式： shared_cluster：共享。 exclusive_cluster：独享。 edge_node：边缘节点。 默认值为shared_cluster。
	RunMode *string `json:"run_mode,omitempty"`

	// 用户为作业选择的CU数量。默认值为2。
	CuNumber *int32 `json:"cu_number,omitempty"`

	// 用户设置的作业并行数目。默认值为1。
	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	// 是否开启作业自动快照功能。 开启：true; 关闭：false; 默认：false.
	CheckpointEnabled *bool `json:"checkpoint_enabled,omitempty"`

	// 快照模式： 1表示ExactlyOnce：数据只被消费一次。 2表示AtLeastOnce：数据至少被消费一次。 默认值为1。
	CheckpointMode *int32 `json:"checkpoint_mode,omitempty"`

	// 快照时间间隔, 单位为秒，默认值为10。
	CheckpointInterval *int32 `json:"checkpoint_interval,omitempty"`

	// 当checkpoint_enabled为true时，该参数是用户授权保存快照的OBS路径。当log_enabled 为true时，该参数是用户授权保存作业日志的OBS路径。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// 是否开启作业的日志上传到用户的OBS功能，默认为false。
	LogEnabled *bool `json:"log_enabled,omitempty"`

	// 当作业异常时，向该SMN主题推送告警信息。
	SmnTopic *string `json:"smn_topic,omitempty"`

	// 是否开启作业异常自动重启，默认为false。
	RestartWhenException *bool `json:"restart_when_exception,omitempty"`

	// 空闲状态过期周期，单位为秒，默认值为3600。
	IdleStateRetention *int32 `json:"idle_state_retention,omitempty"`

	// 边缘计算组ID列表。
	EdgeGroupIds *[]string `json:"edge_group_ids,omitempty"`

	// 作业脏数据策略。 “2”：保存； “1”：抛出异常； “0”：忽略； 默认值为“0”。
	DirtyDataStrategy *string `json:"dirty_data_strategy,omitempty"`

	// 用户已上传到DLI资源管理系统的资源包名，用户sql作业的udf jar通过该参数传入。
	UdfJarUrl *string `json:"udf_jar_url,omitempty"`

	// 用户为作业选择的管理单元（jobmanager）CU数量，默认值为“1”。
	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	// 每个taskmanager的CU数，默认值为“1”。
	TmCus *int32 `json:"tm_cus,omitempty"`

	// 每个taskmanager的slot数，默认值为“(parallel_number*tm_cus)/(cu_number-manager_cu_number)”。
	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	// 异常重启是否从checkpoint恢复。
	ResumeCheckpoint *bool `json:"resume_checkpoint,omitempty"`

	// 异常重试最大次数，单位：次/小时。取值范围：-1或大于0。默认值为“-1”，表示无限次数。
	ResumeMaxNum *int32 `json:"resume_max_num,omitempty"`

	// Flink作业运行时自定义优化参数。
	RuntimeConfig *string `json:"runtime_config,omitempty"`

	// 算子的并行度配置。
	OperatorConfig *string `json:"operator_config,omitempty"`

	// 每个算子的流量/命中率配置，json格式的字符串。例如： {\"operator_list\":[   {\"id\":\"0a448493b4782967b150582570326227\",\"rate_factor\":0.55},   {\"id\":\"6d2677a0ecc3fd8df0b72ec675edf8f4\",\"rate_factor\":1},   {\"id\":\"ea632d67b7d595e5b851708ae9ad79d6\",\"rate_factor\":0.55},   {\"id\":\"bc764cd8ddf7a0cff126f51c16239658\",\"output_rate\":2000} ]}
	StaticEstimatorConfig *string `json:"static_estimator_config,omitempty"`

	// Flink版本。当前只支持1.10和1.12。
	FlinkVersion *string `json:"flink_version,omitempty"`
}

func (o UpdateFlinkSqlJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkSqlJobRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFlinkSqlJobRequestBody", string(data)}, " ")
}
