package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditLogsRequest Request Object
type ListAuditLogsRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  查询开始时间。  **约束限制**：  格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**：  查询结束时间。  **约束限制**： 格式为“yyyy-mm-ddThh:mm:ssZ”， 大于查询开始时间，时间跨度不超过30天，  其中，T指某个时间的开始，Z指时区偏移量，例如北京时间偏移显示为+0800。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**：    索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。    **约束限制**：    必须为整数，不能为负数。    **取值范围**：    ≥0  **默认取值**：   0
	Offset int32 `json:"offset"`

	// **参数解释**：  查询记录数。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  1-100  **默认取值**：  10
	Limit int32 `json:"limit"`
}

func (o ListAuditLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditLogsRequest struct{}"
	}

	return strings.Join([]string{"ListAuditLogsRequest", string(data)}, " ")
}
