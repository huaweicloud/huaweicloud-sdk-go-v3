package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkJarJobRequestBody 新建Flink Jar作业的请求参数。
type CreateFlinkJarJobRequestBody struct {

	// 作业名称。长度限制：0-57个字符。
	Name string `json:"name"`

	// 作业描述。长度限制：0-512个字符。
	Desc *string `json:"desc,omitempty"`

	// 队列名称。长度限制：1-128个字符。
	QueueName *string `json:"queue_name,omitempty"`

	// 用户为作业选择的CU数量，默认值为2。
	CuNumber *int32 `json:"cu_number,omitempty"`

	// 用户为作业选择的管理节点CU数量，对应为flink jobmanager数量，默认值为1。
	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	// 用户为作业选择的并发量，默认值为1.
	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	// 是否开启作业日志。 开启：true 关闭：false 默认：false
	LogEnabled *bool `json:"log_enabled,omitempty"`

	// 当“log_enabled”为“true”时, 用户授权保存作业日志的OBS桶名。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// 当作业异常时，向该SMN主题推送告警信息。
	SmnTopic *string `json:"smn_topic,omitempty"`

	// 作业入口类。
	MainClass *string `json:"main_class,omitempty"`

	// 作业入口类参数，多个参数之间空格分隔。
	EntrypointArgs *string `json:"entrypoint_args,omitempty"`

	// 是否开启异常重启功能，默认值为“false”。
	RestartWhenException *bool `json:"restart_when_exception,omitempty"`

	// 选择Jar作业程序包。 Jar包的管理方式： 上传OBS管理程序包：提前将对应的Jar包上传至OBS桶中。并在此处选择对应的OBS路径。 上传DLI管理程序包：提前将对应的Jar包上传至OBS桶中，并在DLI管理控制台的“数据管理>程序包管理”中创建程序包。 Flink1.15版本不推荐DLI管理程序包，Flink1.15版本以上不再支持DLI管理程序包。
	Entrypoint *string `json:"entrypoint,omitempty"`

	// 用户自定义的依赖程序包。依赖的相关程序包将会被放置到集群classpath下。 依赖程序包程序包的管理方式： 上传OBS管理依赖程序包：提前将对应的Jar包上传至OBS桶中。并在此处选择对应的OBS路径。 上传DLI管理依赖程序包：提前将对应的Jar包上传至OBS桶中，并在DLI管理控制台的“数据管理>程序包管理”中创建程序包 Flink1.15版本不推荐DLI管理依赖程序包，Flink1.15版本以上不再支持DLI管理依赖程序包。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// 用户自定义的依赖文件。 依赖文件的管理方式： 上传OBS管理依赖文件：提前将对应的依赖文件上传至OBS桶中。并在此处选择对应的OBS路径。 上传DLI管理依赖文件：提前将对应的依赖文件上传至OBS桶中，并在DLI管理控制台的“数据管理>程序包管理”中创建程序。 Flink1.15版本不推荐DLI管理依赖文件，Flink1.15版本以上不再支持DLI管理依赖文件。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	// Flink版本。当用户设置“feature”为“basic”时，该参数生效。用户可通过与“feature”参数配合使用，指定作业运行使用的DLI基础Flink镜像的版本。
	FlinkVersion *string `json:"flink_version,omitempty"`

	// 自定义镜像。格式为：组织名/镜像名:镜像版本。当用户设置“feature”为“custom”时，该参数生效。用户可通过与“feature”参数配合使用，指定作业运行使用自定义的Flink镜像。关于如何使用自定义镜像，请参考《数据湖探索用户指南》
	Image *string `json:"image,omitempty"`

	// 每个taskmanager的slot数，默认值为“(parallel_number*tm_cus)/(cu_number-manager_cu_number)”。
	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	// 每个taskmanager的CU数，默认值为“1”。
	TmCus *int32 `json:"tm_cus,omitempty"`

	// 作业特性。表示用户作业使用的Flink镜像类型。basic：表示使用DLI提供的基础Flink镜像。custom：表示使用用户自定义的Flink镜像。
	Feature *string `json:"feature,omitempty"`

	// 异常重启是否从checkpoint恢复。
	ResumeCheckpoint *bool `json:"resume_checkpoint,omitempty"`

	// 异常重试最大次数，单位：次/小时。取值范围：-1或大于0。默认值为“-1”，表示无限次数。
	ResumeMaxNum *int32 `json:"resume_max_num,omitempty"`

	// 用户Jar中checkpoint的储存地址，不同作业路径需要保持不同。
	CheckpointPath *string `json:"checkpoint_path,omitempty"`

	// Flink作业运行时自定义优化参数。
	RuntimeConfig *string `json:"runtime_config,omitempty"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`

	ResourceConfig *ResourceConfig `json:"resource_config,omitempty"`

	// 资源配置版本。可选值 \"v1\" ,\"v2\".默认为“v1”。
	ResourceConfigVersion *string `json:"resource_config_version,omitempty"`
}

func (o CreateFlinkJarJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkJarJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlinkJarJobRequestBody", string(data)}, " ")
}
