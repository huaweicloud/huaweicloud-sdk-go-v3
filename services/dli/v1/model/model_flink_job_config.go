package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkJobConfig 作业配置，当“show_detail”为“true”时才有该参数配置。
type FlinkJobConfig struct {

	// 参数解释:  是否开启作业自动快照功能 示例: false 约束限制:  无 取值范围: true,false 默认取值: false
	CheckpointEnabled *bool `json:"checkpoint_enabled,omitempty"`

	// 参数解释:  快照模式 示例: exactly_once 约束限制:  无 取值范围: exactly_once（数据只被消费一次） at_least_once（数据至少被消费一次） 默认取值: 无
	CheckpointMode *FlinkJobConfigCheckpointMode `json:"checkpoint_mode,omitempty"`

	// 参数解释:  快照时间间隔, 单位为秒 示例: 10 约束限制:  无 取值范围: 无 默认取值: 10
	CheckpointInterval *int32 `json:"checkpoint_interval,omitempty"`

	// 参数解释:  是否启用日志存储 示例: false 约束限制:  无 取值范围: true,false 默认取值: false
	LogEnabled *bool `json:"log_enabled,omitempty"`

	// 参数解释:  OBS桶名 示例: obs-demo 约束限制:  无 取值范围: 无 默认取值: 无
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// 参数解释:  当作业异常时，向该SMN主题推送告警信息 示例: cs_job_exception 约束限制:  无 取值范围: 无 默认取值: 无
	SmnTopic *string `json:"smn_topic,omitempty"`

	// 参数解释:  边缘计算组ID列表 示例: 62de1e1c-066e-48a8-a79d-f461a31b2ee1,2eb00f85-99f2-4144-bcb7-d39ff47f9002 约束限制:  无 取值范围: 无 默认取值: 无
	EdgeGroupIds *[]string `json:"edge_group_ids,omitempty"`

	// 参数解释:  父作业ID 示例: 11 约束限制:  无 取值范围: 无 默认取值: 无
	RootId *int64 `json:"root_id,omitempty"`

	// 参数解释:  管理单元CU数 示例: 1 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	// 参数解释:  用户为作业选择的CU数量 示例: 2 约束限制:  无 取值范围: [2,400] 默认取值: 2
	CuNumber *int32 `json:"cu_number,omitempty"`

	// 参数解释:  用户设置的作业并行数， “show_detail”为“true”时独有 示例: 1 约束限制:  无 取值范围: [1,2000] 默认取值: 1
	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	// 参数解释:  是否开启异常重启功能 示例: false 约束限制:  无 取值范围: true,false 默认取值: 无
	RestartWhenException *bool `json:"restart_when_exception,omitempty"`

	// 参数解释:  空闲状态过期周期 示例: 3600 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	IdleStateRetention *int32 `json:"idle_state_retention,omitempty"`

	// 参数解释:  用户已上传到DLI资源管理系统的资源包名，用户sql作业的udf jar通过该参数传入 示例: obs://cs-append/jar/udf/flink-1.13-demo-1.0-udf.jar obs://onlyci-7/flink/udx.jar 约束限制:  无 取值范围: 无 默认取值: 无
	UdfJarUrl *string `json:"udf_jar_url,omitempty"`

	// 参数解释:  作业脏数据策略 示例: 0 约束限制:  无 取值范围: 0（忽略） 1（抛出异常） 2:obsDir（保存，obsDir表示脏数据存储路径） 默认取值: 无
	DirtyDataStrategy *string `json:"dirty_data_strategy,omitempty"`

	// 参数解释:  用户已上传到DLI资源管理系统的资源包名，用户自定义作业主类所在的jar包 示例: obs://test/demo/WindowJoin.jar 约束限制:  无 取值范围: 无 默认取值: 无
	Entrypoint *string `json:"entrypoint,omitempty"`

	// 参数解释:  用户已上传到DLI资源管理系统的资源包名，用户自定义作业的其他依赖包 示例: [\"zsdas/wordcount.jar\",\"ygj/flink-dis-to-kafka-v5.jar\"] 约束限制:  无 取值范围: 无 默认取值: 无
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// 参数解释:  用户已上传到DLI资源管理系统的资源包名，用户自定义作业的依赖文件 示例: [\"zsdas/wordcount.jar\",\"ygj/flink-dis-to-kafka-v5.jar\"] 约束限制:  无 取值范围: 无 默认取值: 无
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	// 参数解释:  作业使用计算节点个数 示例: 0 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	ExecutorNumber *int32 `json:"executor_number,omitempty"`

	// 参数解释:  计算节点cu数 示例: 0 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	ExecutorCuNumber *int32 `json:"executor_cu_number,omitempty"`

	// 参数解释:  授权给DLI的委托名。Flink1.15版本时支持配置该参数。 示例: agency 约束限制:  无 取值范围: 无 默认取值: 无
	ExecutionAgencyUrn *string `json:"execution_agency_urn,omitempty"`

	// 参数解释:  异常自动重启时，是否从最新checkpoint恢复，默认false 示例: 0 约束限制:  无 取值范围: true,false 默认取值: false
	ResumeCheckpoint *bool `json:"resume_checkpoint,omitempty"`

	// 参数解释: Flink作业运行时自定义优化参数 示例: [{\\\"key\\\":\\\"high-availability\\\",\\\"value\\\":\\\"org.apache.flink.kubernetes.highavailability.KubernetesHaServicesFactory\\\" },{ \\\"key\\\":\\\"kubernetes.jobmanager.replicas\\\",\\\"value\\\":\\\"2\\\" },{ \\\"key\\\":\\\"high-availability.storageDir\\\",\\\"value\\\":\\\"obs://fz-test/test\\\"}] 约束限制:  无 取值范围: 无 默认取值: 无
	RuntimeConfig *string `json:"runtime_config,omitempty"`

	// 参数解释: 流图编辑开关 示例: false 约束限制:  无 取值范围: 无 默认取值: false
	GraphEditorEnabled *bool `json:"graph_editor_enabled,omitempty"`

	// 参数解释: 流图编辑数据 约束限制:  无 取值范围: 无 默认取值: 无
	GraphEditorData *string `json:"graph_editor_data,omitempty"`

	// 参数解释: 异常重试最大次数。-1代表无限 示例: -1 约束限制:  无 取值范围: 无 默认取值: 无
	ResumeMaxNum *int64 `json:"resume_max_num,omitempty"`

	// 参数解释: 检查点保存路径 示例: obs://cwk/checkpoint/ 约束限制:  无 取值范围: 无 默认取值: 无
	CheckpointPath *string `json:"checkpoint_path,omitempty"`

	// 参数解释: 用户上传的config包OBS路径。 示例: obs://bucket-62099355b894428e8916573ae635f1f9/di-flink/lib/client.jks,obs://bucket-62099355b894428e8916573ae635f1f9/di-flink/lib/client.jks 约束限制:  无 取值范围: 无 默认取值: 无
	ConfigUrl *string `json:"config_url,omitempty"`

	// 参数解释: 单TM所占CU数 示例: 1 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	TmCus *int32 `json:"tm_cus,omitempty"`

	// 参数解释: 单TM Slot数 示例: 1 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	// 参数解释: 自定义镜像。格式为：组织名/镜像名:镜像版本。当用户设置“feature”为“custom”时，该参数生效。用户可通过与“feature”参数配合使用，指定作业运行使用自定义的Flink镜像 示例: dli/spark:2.4.8 约束限制:  无 取值范围: 无 默认取值: 无
	Image *string `json:"image,omitempty"`

	// 参数解释: 自定义作业特性。表示用户作业使用的Flink镜像类型。basic：表示使用DLI提供的基础Flink镜像。custom：表示使用用户自定义的Flink镜像 示例: basic 约束限制:  无 取值范围: basic：表示使用DLI提供的基础Flink镜像 custom：表示使用用户自定义的Flink镜像 默认取值: 无
	Feature *string `json:"feature,omitempty"`

	// 参数解释: Flink版本。当用户设置“feature”为“basic”时，该参数生效。用户可通过与“feature”参数配合使用，指定作业运行使用的DLI基础Flink镜像的版本 示例: 1.17 约束限制:  无 取值范围: 无 默认取值: 无
	FlinkVersion *string `json:"flink_version,omitempty"`

	// 参数解释: 各算子并行度参数，以json的形式展示各算子id和并行度 示例: '{\"operator_list\":[{\"id\":\"0a448493b4782967b150582570326227\",\"parallelism\":1},{\"id\":\"6d2677a0ecc3fd8df0b72ec675edf8f4\",\"parallelism\":1},{\"id\":\"ea632d67b7d595e5b851708ae9ad79d6\",\"parallelism\":1},{\"id\":\"ddb598ad156ed281023ba4eebbe487e3\",\"parallelism\":1},{\"id\":\"bc764cd8ddf7a0cff126f51c16239658\",\"parallelism\":1}]}' 约束限制:  符合Json格式的字符串 取值范围: 无 默认取值: 无
	OperatorConfig *string `json:"operator_config,omitempty"`

	// 参数解释: 静态流图资源预估参数，以json的形式展示 示例: '{\"operator_list\":[{\"id\":\"0a448493b4782967b150582570326227\",\"parallelism\":1},{\"id\":\"6d2677a0ecc3fd8df0b72ec675edf8f4\",\"parallelism\":1},{\"id\":\"ea632d67b7d595e5b851708ae9ad79d6\",\"parallelism\":1},{\"id\":\"ddb598ad156ed281023ba4eebbe487e3\",\"parallelism\":1},{\"id\":\"bc764cd8ddf7a0cff126f51c16239658\",\"parallelism\":1}]}' 约束限制:  符合Json格式的字符串 取值范围: 无 默认取值: 无
	StaticEstimatorConfig *string `json:"static_estimator_config,omitempty"`

	// 参数解释: 实际使用的CU数 示例: 0 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 0
	RealCuNumber *int32 `json:"real_cu_number,omitempty"`

	ResourceConfig *ResourceConfig `json:"resource_config,omitempty"`

	// 资源配置版本。可选值 \"v1\" ,\"v2\".默认为“v1”。
	ResourceConfigVersion *string `json:"resource_config_version,omitempty"`
}

func (o FlinkJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobConfig struct{}"
	}

	return strings.Join([]string{"FlinkJobConfig", string(data)}, " ")
}

type FlinkJobConfigCheckpointMode struct {
	value string
}

type FlinkJobConfigCheckpointModeEnum struct {
	EXACTLY_ONCE  FlinkJobConfigCheckpointMode
	AT_LEAST_ONCE FlinkJobConfigCheckpointMode
}

func GetFlinkJobConfigCheckpointModeEnum() FlinkJobConfigCheckpointModeEnum {
	return FlinkJobConfigCheckpointModeEnum{
		EXACTLY_ONCE: FlinkJobConfigCheckpointMode{
			value: "exactly_once",
		},
		AT_LEAST_ONCE: FlinkJobConfigCheckpointMode{
			value: "at_least_once",
		},
	}
}

func (c FlinkJobConfigCheckpointMode) Value() string {
	return c.value
}

func (c FlinkJobConfigCheckpointMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkJobConfigCheckpointMode) UnmarshalJSON(b []byte) error {
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
