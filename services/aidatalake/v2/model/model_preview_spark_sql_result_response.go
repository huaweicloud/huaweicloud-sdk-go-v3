package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreviewSparkSqlResultResponse Response Object
type PreviewSparkSqlResultResponse struct {

	// **参数解释**：查询类SQL结果格式，用于指定结果的存储格式。 **取值范围**：csv，arrow。
	ResultFormat *string `json:"result_format,omitempty"`

	// **参数解释**：查询类SQL结果的OBS路径，用于存储SQL查询的结果数据。 **取值范围**：采用OBS路径格式，例如：obs://bucket/aidatalake/workspace_xxx/spark/endpoint_xxx/jobs/sql_result/{statement_id}/result.csv。
	ResultPath *string `json:"result_path,omitempty"`

	// **参数解释**：结果数据的列结构定义，包含列名和列类型信息。数组中的每个元素为SparkSqlSchemaItem对象。
	Schema *[]SparkSqlSchemaItem `json:"schema,omitempty"`

	// **参数解释**：作业结果集，包含查询返回的实际数据行。每行为一个数组，对应schema中定义的列顺序。
	Rows           *[][]string `json:"rows,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o PreviewSparkSqlResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewSparkSqlResultResponse struct{}"
	}

	return strings.Join([]string{"PreviewSparkSqlResultResponse", string(data)}, " ")
}
