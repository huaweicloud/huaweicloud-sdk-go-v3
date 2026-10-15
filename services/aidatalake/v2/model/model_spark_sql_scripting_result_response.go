package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkSqlScriptingResultResponse **参数解释**：SQL脚本查询结果信息，包含结果格式和存储路径。
type SparkSqlScriptingResultResponse struct {

	// **参数解释**：查询类SQL结果格式，用于指定结果数据的输出格式。 **取值范围**：支持csv格式。
	ResultFormat *string `json:"result_format,omitempty"`

	// **参数解释**：查询类SQL结果OBS路径，用于指定结果数据的存储位置。 **取值范围**：OBS URL格式，长度为1~1024个字符，例如：obs://bucket/results/result.csv.
	ResultPath *string `json:"result_path,omitempty"`
}

func (o SparkSqlScriptingResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlScriptingResultResponse struct{}"
	}

	return strings.Join([]string{"SparkSqlScriptingResultResponse", string(data)}, " ")
}
