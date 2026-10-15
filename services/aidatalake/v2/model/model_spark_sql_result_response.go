package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkSqlResultResponse **参数解释**：SparkSql作业结果信息，包含查询结果的格式和路径等。
type SparkSqlResultResponse struct {

	// **参数解释**：查询类SQL结果格式，用于指定结果的存储格式。 **取值范围**：csv、arrow。
	ResultFormat *string `json:"result_format,omitempty"`

	// **参数解释**：查询类SQL结果的OBS路径，用于存储SQL查询的结果数据。 **取值范围**：采用OBS路径格式，例如：obs://bucket/aidatalake/workspace_xxx/spark/endpoint_xxx/jobs/sql_result/{statement_id}/result.csv。
	ResultPath *string `json:"result_path,omitempty"`

	// **参数解释**：查询结果行数，表示SQL查询返回的数据行数。
	ResultRecords *int64 `json:"result_records,omitempty"`

	// **参数解释**：查询结果字节数，表示SQL查询返回的数据大小。
	ResultBytes *int64 `json:"result_bytes,omitempty"`

	// **参数解释**：作业执行计划的存储路径，用于存储作业的执行计划信息。 **取值范围**：采用OBS路径格式，例如：obs://bucket/aidatalake/workspace_xxx/spark/endpoint_xxx/jobs/sql_profile/{statement_id}/。
	ExecutionProfilePath *string `json:"execution_profile_path,omitempty"`
}

func (o SparkSqlResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlResultResponse struct{}"
	}

	return strings.Join([]string{"SparkSqlResultResponse", string(data)}, " ")
}
