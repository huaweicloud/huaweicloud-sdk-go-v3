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

	// 用户已上传到DLI资源管理系统的资源包名，用户自定义作业主类所在的jar包。
	Entrypoint *string `json:"entrypoint,omitempty"`

	// 用户已上传到DLI资源管理系统的资源包名，用户自定义作业的其他依赖包。示例：[Group/test.jar,myGroup/test1.jar]
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// 用户已上传到DLI资源管理系统的资源包名，用户自定义作业的依赖文件，示例：[myGroup/test.cvs,myGroup/test1.csv]
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
}

func (o CreateFlinkJarJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkJarJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlinkJarJobRequestBody", string(data)}, " ")
}
