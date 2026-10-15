package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkSqlErrorDto **参数解释**：SparkSql作业错误详情，包含错误描述、错误码等信息。
type SparkSqlErrorDto struct {

	// **参数解释**：SparkSql错误描述，用于说明错误的具体原因。 **取值范围**：长度为1~1024个字符，例如：表不存在。
	Message *string `json:"message,omitempty"`

	// **参数解释**：SparkSql错误码，用于标识错误的类型。 **取值范围**：采用标准SQL错误码格式，例如：42P01。
	SqlState *string `json:"sql_state,omitempty"`

	// **参数解释**：SparkSql错误类型，用于标识错误的分类。 **取值范围**：长度为1~128个字符，例如：SCHEMA_ALREADY_EXISTS。
	ErrorClass *string `json:"error_class,omitempty"`

	// **参数解释**：报错行号，用于定位SQL语句中的错误位置。 **取值范围**：大于0的整数，例如：2。
	Line *int32 `json:"line,omitempty"`

	// **参数解释**：报错起始位置，用于定位SQL语句中错误的起始字符位置。 **取值范围**：大于0的整数，例如：2。
	StartPosition *int32 `json:"start_position,omitempty"`
}

func (o SparkSqlErrorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlErrorDto struct{}"
	}

	return strings.Join([]string{"SparkSqlErrorDto", string(data)}, " ")
}
