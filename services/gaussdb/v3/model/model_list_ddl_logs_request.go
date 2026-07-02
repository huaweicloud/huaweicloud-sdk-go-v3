package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDdlLogsRequest Request Object
type ListDdlLogsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  租户在某一project下的实例ID  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  实例下的节点ID  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  ≥0  **默认取值**：  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  每页显示条数。  **约束限制**：  不涉及。  **取值范围**：  1-100  **默认取值**：  10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  开始日期。格式为\"yyyy-mm-ddThh:mm:ssZ\"。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：  结束时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。只能查询当前时间前一个月内的慢日志。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListDdlLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDdlLogsRequest struct{}"
	}

	return strings.Join([]string{"ListDdlLogsRequest", string(data)}, " ")
}
