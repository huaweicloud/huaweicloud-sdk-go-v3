package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkSqlJobRequestBody 创建SQL作业的请求参数。
type CreateFlinkSqlJobRequestBody struct {

	// 作业名称。长度限制：0-57个字符。
	Name string `json:"name"`

	// 作业描述。长度限制：0-2048个字符。
	Desc *string `json:"desc,omitempty"`

	// 模板Id。  如果template_id和sql_body都不为空，优先sql_body；如果template_id不空，sql_body为空，以template_id内容填充sql_body。
	TemplateId *int64 `json:"template_id,omitempty"`

	// 队列名称。长度限制：1-128个字符。
	QueueName *string `json:"queue_name,omitempty"`

	// Stream SQL语句。长度限制：1024*1024个字符。
	SqlBody *string `json:"sql_body,omitempty"`

	// 作业运行模式： shared_cluster：共享。 exclusive_cluster：独享。 edge_node：边缘节点。 默认值为：shared_cluster
	RunMode *string `json:"run_mode,omitempty"`

	// 用户为作业选择的CU数量，默认值为2。
	CuNumber *int32 `json:"cu_number,omitempty"`

	// 用户设置的作业并行数目。默认值为1。
	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	// 是否开启作业自动快照功能。 开启：true； 关闭：false； 默认：false。
	CheckpointEnabled *bool `json:"checkpoint_enabled,omitempty"`

	// 快照模式： 1表示ExactlyOnce：数据只被消费一次。 2表示AtLeastOnce：数据至少被消费一次。 默认值为1。
	CheckpointMode *int32 `json:"checkpoint_mode,omitempty"`

	// 快照时间间隔, 单位为秒。默认值为10。
	CheckpointInterval *int32 `json:"checkpoint_interval,omitempty"`

	// 当checkpoint_enabled==true时，该参数是用户授权保存快照的OBS路径。当log_enabled ==true时，该参数是用户授权保存作业日志的OBS路径。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// 是否开启作业的日志上传到用户的OBS功能。默认为false。
	LogEnabled *bool `json:"log_enabled,omitempty"`

	// 当作业异常时，向该SMN主题推送告警信息。
	SmnTopic *string `json:"smn_topic,omitempty"`

	// 是否开启作业异常自动重启。默认为false。
	RestartWhenException *bool `json:"restart_when_exception,omitempty"`

	// 空闲状态过期周期，单位为秒，默认值为3600。
	IdleStateRetention *int32 `json:"idle_state_retention,omitempty"`

	// 作业类型：flink_sql_job、flink_opensource_sql_job。 默认值：“flink_opensource_sql_job”。  “run_mode”为“exclusive_cluster”时，作业类型须为“flink_sql_job”或“flink_opensource_sql_job”。  “run_mode””为“shared_cluster”时作业类型必须为”flink_sql_job“。
	JobType *string `json:"job_type,omitempty"`

	// 边缘计算组ID列表。
	EdgeGroupIds *[]string `json:"edge_group_ids,omitempty"`

	// 作业脏数据策略。 “2”：保存； “1”：抛出异常； “0”：忽略； 默认值为“0”。
	DirtyDataStrategy *string `json:"dirty_data_strategy,omitempty"`

	// 用户自定义UDF文件，在后续作业中可以调用插入Jar包中的自定义函数。 UDF Jar包的管理方式： 上传OBS管理UDF Jar包：提前将对应的Jar包上传至OBS桶中。并在此处选择对应的OBS路径。 上传DLI管理UDF Jar包：提前将对应的Jar包上传至OBS桶中，并在DLI管理控制台的“数据管理>程序包管理”中创建程序包。 Flink1.15版本不再支持DLI管理UDF Jar包。
	UdfJarUrl *string `json:"udf_jar_url,omitempty"`

	// 用户为作业选择的管理单元（jobmanager）CU数量，默认值为“1”。
	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	// 每个taskmanager的CU数，默认值为“1”。
	TmCus *int32 `json:"tm_cus,omitempty"`

	// 每个taskmanager的slot数，默认值为“(parallel_number*tm_cus)/(cu_number-manager_cu_number)”
	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 异常重启是否从checkpoint恢复。
	ResumeCheckpoint *bool `json:"resume_checkpoint,omitempty"`

	// 异常重试最大次数，单位：次/小时。取值范围：-1或大于0。默认值为“-1”，表示无限次数。
	ResumeMaxNum *int32 `json:"resume_max_num,omitempty"`

	// Flink作业运行时自定义优化参数。
	RuntimeConfig *string `json:"runtime_config,omitempty"`

	// Flink版本。当前只支持1.10和1.12。
	FlinkVersion *string `json:"flink_version,omitempty"`

	// 授权给DLI的委托名。Flink1.15版本时支持配置该参数。
	ExecutionAgencyUrn *string `json:"execution_agency_urn,omitempty"`

	// 资源配置版本。可选值 \"v1\" ,\"v2\".默认为“v1”。v2版本对比于v1模版不支持设置CU数量，支持直接设置Job Manager Memory和Task Manager Memory。v1：适用于Flink 1.12、Flink 1.13、Flink 1.15。v2：适用于Flink 1.13、Flink 1.15、Flink 1.17。优先推荐使用V2版本的参数设置。
	ResourceConfigVersion *string `json:"resource_config_version,omitempty"`

	ResourceConfig *ResourceConfig `json:"resource_config,omitempty"`
}

func (o CreateFlinkSqlJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobRequestBody", string(data)}, " ")
}
