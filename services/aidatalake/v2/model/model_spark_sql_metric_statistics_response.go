package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkSqlMetricStatisticsResponse **参数解释**：SparkSql作业的指标统计信息，包含读取数据量、输出数据量、计算资源使用情况等。
type SparkSqlMetricStatisticsResponse struct {

	// **参数解释**：脏数据行数，用于标识处理过程中发现的错误数据行数。 **取值范围**：大于等于0的整数。
	BadRecords *int64 `json:"bad_records,omitempty"`

	// **参数解释**：读取数据字节数，用于标识作业读取的数据量。 **取值范围**：大于等于0的整数，单位为字节。
	InputBytes *int64 `json:"input_bytes,omitempty"`

	// **参数解释**：读取数据行数，用于标识作业读取的数据记录数。 **取值范围**：大于等于0的整数。
	InputRecords *int64 `json:"input_records,omitempty"`

	// **参数解释**：输出数据字节数，用于标识作业输出的数据量。 **取值范围**：大于等于0的整数，单位为字节。
	OutputBytes *int64 `json:"output_bytes,omitempty"`

	// **参数解释**：输出数据行数，用于标识作业输出的数据记录数。 **取值范围**：大于等于0的整数。
	OutputRecords *int64 `json:"output_records,omitempty"`

	// **参数解释**：计算使用CPU秒数，用于标识作业消耗的计算资源。 **取值范围**：大于等于0的整数，单位为秒。
	CpuTime *int64 `json:"cpu_time,omitempty"`
}

func (o SparkSqlMetricStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlMetricStatisticsResponse struct{}"
	}

	return strings.Join([]string{"SparkSqlMetricStatisticsResponse", string(data)}, " ")
}
