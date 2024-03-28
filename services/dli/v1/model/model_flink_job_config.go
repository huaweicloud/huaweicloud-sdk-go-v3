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
	RootId *int32 `json:"root_id,omitempty"`

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
}

func (o FlinkJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobConfig struct{}"
	}

	return strings.Join([]string{"FlinkJobConfig", string(data)}, " ")
}
