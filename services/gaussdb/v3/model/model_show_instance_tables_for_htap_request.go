package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTablesForHtapRequest Request Object
type ShowInstanceTablesForHtapRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage string `json:"X-Language"`

	// **参数解释**：  HTAP标准版实例ID，严格匹配UUID规则。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，且长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  查询记录数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  100。
	Limit *string `json:"limit,omitempty"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  0。
	Offset *string `json:"offset,omitempty"`

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认取值**：  application/json。
	ContentType string `json:"Content-Type"`

	Body *QueryTableRequestV3 `json:"body,omitempty"`
}

func (o ShowInstanceTablesForHtapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTablesForHtapRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceTablesForHtapRequest", string(data)}, " ")
}
