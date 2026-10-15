package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkSqlTimeout **参数解释**：SparkSql作业超时配置信息，包含排队超时和运行超时时间。
type SparkSqlTimeout struct {

	// **参数解释**：作业排队超时时间，单位为分钟。如果超过此时间作业仍未开始运行，则作业会被标记为排队超时并失败。 **取值范围**：10~180分钟。
	QueuedTimeout *int32 `json:"queued_timeout,omitempty"`

	// **参数解释**：作业运行超时时间，单位为分钟。如果超过此时间作业仍未运行结束，则作业会被取消并标记为运行超时。 **取值范围**：10~720分钟。
	RunningTimeout *int32 `json:"running_timeout,omitempty"`
}

func (o SparkSqlTimeout) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlTimeout struct{}"
	}

	return strings.Join([]string{"SparkSqlTimeout", string(data)}, " ")
}
