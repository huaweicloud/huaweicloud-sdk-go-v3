package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobConfig 作业配置，当“show_detail”为“true”时才有该参数配置。
type FlinkJobConfig struct {

	// 是否开启作业自动快照功能。 开启：true； 关闭：false； 默认：false。
	CheckpointEnabled *bool `json:"checkpoint_enabled,omitempty"`

	// 快照模式： exactly_once：数据只被消费一次。 at_least_once：数据至少被消费一次。 默认值为exactly_once。
	CheckpointMode *string `json:"checkpoint_mode,omitempty"`

	// 快照时间间隔, 单位为秒，默认值为10。
	CheckpointInterval *int32 `json:"checkpoint_interval,omitempty"`

	// 是否启用日志存储。默认为false。
	LogEnabled *bool `json:"log_enabled,omitempty"`

	// OBS桶名。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// 当作业异常时，向该SMN主题推送告警信息。
	SmnTopic *string `json:"smn_topic,omitempty"`

	// 边缘计算组ID列表。
	EdgeGroupIds *[]string `json:"edge_group_ids,omitempty"`

	// 父作业ID。
	RootId *int64 `json:"root_id,omitempty"`

	// 管理单元CU数。默认为1。
	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	// 用户为作业选择的CU数量, “show_detail”。默认为2。
	CuNumber *int32 `json:"cu_number,omitempty"`

	// 用户设置的作业并行数， “show_detail”为“true”时独有。默认值为1。 最小值：1，最大值：2000。
	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	// 是否开启异常重启功能。
	RestartWhenException *bool `json:"restart_when_exception,omitempty"`

	// 空闲状态过期周期。
	IdleStateRetention *int32 `json:"idle_state_retention,omitempty"`

	// 用户已上传到DLI资源管理系统的资源包名，用户sql作业的udf jar通过该参数传入。
	UdfJarUrl *string `json:"udf_jar_url,omitempty"`

	// 作业脏数据策略。 “2:obs-wan-wulan3/jobs”：保存 “1”：抛出异常 “0”：忽略
	DirtyDataStrategy *string `json:"dirty_data_strategy,omitempty"`

	// 用户已上传到DLI资源管理系统的资源包名，用户自定义作业主类所在的jar包.
	Entrypoint *string `json:"entrypoint,omitempty"`

	// 用户已上传到DLI资源管理系统的资源包名，用户自定义作业的其他依赖包
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// 用户已上传到DLI资源管理系统的资源包名，用户自定义作业的依赖文件.
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	// 作业使用计算节点个数
	ExecutorNumber *int32 `json:"executor_number,omitempty"`

	// 计算节点cu数
	ExecutorCuNumber *int32 `json:"executor_cu_number,omitempty"`

	// 异常自动重启时，是否从最新checkpoint恢复，默认false
	ResumeCheckpoint *bool `json:"resume_checkpoint,omitempty"`

	// Flink作业运行时自定义优化参数。
	RuntimeConfig *string `json:"runtime_config,omitempty"`

	// 流图编辑开关。默认为“false。
	GraphEditorEnabled *bool `json:"graph_editor_enabled,omitempty"`

	// 流图编辑数据。默认为null。
	GraphEditorData *string `json:"graph_editor_data,omitempty"`

	// 异常重试最大次数。-1代表无限。
	ResumeMaxNum *int64 `json:"resume_max_num,omitempty"`

	// 检查点保存路径。
	CheckpointPath *string `json:"checkpoint_path,omitempty"`

	// 用户上传的config包OBS路径。
	ConfigUrl *string `json:"config_url,omitempty"`

	// 单TM所占CU数。
	TmCus *int32 `json:"tm_cus,omitempty"`

	// 单TM Slot数。
	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	// 自定义镜像。格式为：组织名/镜像名:镜像版本。当用户设置“feature”为“custom”时，该参数生效。用户可通过与“feature”参数配合使用，指定作业运行使用自定义的Flink镜像。
	Image *string `json:"image,omitempty"`

	// 自定义作业特性。表示用户作业使用的Flink镜像类型。basic：表示使用DLI提供的基础Flink镜像。custom：表示使用用户自定义的Flink镜像。
	Feature *string `json:"feature,omitempty"`

	// Flink版本。当用户设置“feature”为“basic”时，该参数生效。用户可通过与“feature”参数配合使用，指定作业运行使用的DLI基础Flink镜像的版本。
	FlinkVersion *string `json:"flink_version,omitempty"`

	// 各算子并行度参数，以json的形式展示各算子id和并行度。
	OperatorConfig *string `json:"operator_config,omitempty"`

	// 静态流图资源预估参数，以json的形式展示。
	StaticEstimatorConfig *string `json:"static_estimator_config,omitempty"`

	RealCuNumber *int32 `json:"real_cu_number,omitempty"`
}

func (o FlinkJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobConfig struct{}"
	}

	return strings.Join([]string{"FlinkJobConfig", string(data)}, " ")
}
