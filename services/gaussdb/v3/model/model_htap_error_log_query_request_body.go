package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HtapErrorLogQueryRequestBody **参数解释**：  查看错误日志请求体。  **约束限制**：  不涉及。
type HtapErrorLogQueryRequestBody struct {

	// **参数解释**： HTAP标准版实例节点ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**： 日志开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **取值范围**： 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**： 日志结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**： 日志级别。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **取值范围**： 不涉及。
	Level string `json:"level"`

	// **参数解释**： 查询记录数。  **约束限制**：  不涉及。  **取值范围**： 0-100。  **默认取值**： 不涉及。
	Limit int32 `json:"limit"`

	// **参数解释**： 日志单行序列号，第一次查询时不需要此参数，后续分页查询时需要使用，可从上次查询的返回信息中获取。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	LineNum *string `json:"line_num,omitempty"`
}

func (o HtapErrorLogQueryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapErrorLogQueryRequestBody struct{}"
	}

	return strings.Join([]string{"HtapErrorLogQueryRequestBody", string(data)}, " ")
}
