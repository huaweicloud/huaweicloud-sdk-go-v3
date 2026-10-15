package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSparkSqlRequestBody 执行SparkSql作业的请求参数。
type RunSparkSqlRequestBody struct {

	// **参数解释**：端点名称，用于指定SparkSql作业运行的计算引擎。可在控制台的端点管理页面查看，或通过查询端点列表接口获取。 **约束限制**：不涉及。 **取值范围**：只能以英文小写字母开头，由英文小写字母、数字及中划线组成，以英文小写字母或数字结尾，且长度为1~63个字符。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`

	CatalogContext *SparkSqlCatalogContext `json:"catalog_context"`

	// **参数解释**：用户SQL语句，用于执行SparkSql作业。支持DDL、DCL、DQL、DML等多种SQL类型。 **约束限制**：不涉及。 **取值范围**：不超过500000个字符。 **默认取值**：不涉及。
	Statement string `json:"statement"`

	// **参数解释**：用户SQL语句中的占位符参数列表，用于SQL参数化执行。数组中的每个元素为SparkSqlParameter对象，包含占位符的键、值和类型信息。 **约束限制**：占位符参数数量不能超过16条。
	Parameters *[]SparkSqlParameter `json:"parameters,omitempty"`

	// **参数解释**：用户自定义Spark参数配置，用于优化作业性能。格式为key/value键值对，Key为参数名称，Value为参数值。例如：spark.executor.memory=4g。 **约束限制**：参数配置项数量不能超过100条，每个参数值的长度不超过1024个字符。
	SparkConfig map[string]string `json:"spark_config,omitempty"`

	Timeout *SparkSqlTimeout `json:"timeout,omitempty"`

	// **参数解释**：作业标签列表，用于标识和分类作业。数组中的每个元素为SparkSqlLabel对象，包含标签的键和值。 **约束限制**：标签数量不能超过16条。
	Labels *[]SparkSqlLabel `json:"labels,omitempty"`
}

func (o RunSparkSqlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSparkSqlRequestBody struct{}"
	}

	return strings.Join([]string{"RunSparkSqlRequestBody", string(data)}, " ")
}
