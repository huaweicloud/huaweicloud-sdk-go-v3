package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSparkJobRequestBody 启动Spark作业的请求参数， job_config：作业配置参数，根据作业类型自动选择对应的参数结构（SparkJarParameter、SparkPyParameter或SparkSqlScriptParameter）。 resource_config：Spark作业的资源配置。 image：Spark镜像配置信息。 restore_strategy：Spark作业的运行及恢复策略。 logging_config：Spark作业日志配置。
type RunSparkJobRequestBody struct {

	// **参数解释**：Spark作业名称，用于标识作业。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：端点名称，用于指定Spark作业执行环境。 **约束限制**：不涉及。 **取值范围**：只能由英文小写字母、数字及中划线组成，以英文小写字母开头，以英文小写字母或数字结尾，且长度为1~63个字符。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`

	JobConfig *SparkJobArtifact `json:"job_config"`

	// **参数解释**：Catalog名称，用于指定作业使用的数据目录。 **约束限制**：不涉及。 **取值范围**：长度不超过128个字符。 **默认取值**：不涉及。
	CatalogName *string `json:"catalog_name,omitempty"`

	// **参数解释**：自定义委托名称，用于作业操作OBS对象、转储日志、访问DLI元数据等。 **约束限制**：不涉及。 **取值范围**：长度为1~64个字符。 **默认取值**：不涉及。
	JobAgency *string `json:"job_agency,omitempty"`

	ResourceConfig *SparkResourceConfig `json:"resource_config,omitempty"`

	// **参数解释**：用户自定义Spark参数配置，用于优化Spark作业性能。 **约束限制**：最多支持100个参数。
	SparkConfig map[string]string `json:"spark_config,omitempty"`

	Image *SparkJobImageConfig `json:"image,omitempty"`

	LoggingConfig *SparkLoggingConfig `json:"logging_config,omitempty"`

	RestoreStrategy *SparkRestoreStrategy `json:"restore_strategy,omitempty"`

	// **参数解释**：作业标签列表，用于标识和分类作业。 **约束限制**：标签数量不能超过16条。
	Labels *[]SparkJobLabel `json:"labels,omitempty"`

	// **参数解释**：Spark作业描述信息，用于说明作业用途。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`
}

func (o RunSparkJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSparkJobRequestBody struct{}"
	}

	return strings.Join([]string{"RunSparkJobRequestBody", string(data)}, " ")
}
