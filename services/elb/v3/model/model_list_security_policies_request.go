package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityPoliciesRequest Request Object
type ListSecurityPoliciesRequest struct {

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：自定义安全策略的ID。 支持多值查询，查询条件格式：*id=xxx&id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Id *[]string `json:"id,omitempty"`

	// **参数解释**：自定义安全策略的名称。 支持多值查询，查询条件格式：*name=xxx&name=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *[]string `json:"name,omitempty"`

	// **参数解释**：自定义安全策略的描述信息。 支持多值查询，查询条件格式：*description=xxx&description=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Description *[]string `json:"description,omitempty"`

	// **参数解释**：空格分隔的自定义安全策略的TLS协议。 支持多值查询，查询条件格式：*protocols=xxx&protocols=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Protocols *[]string `json:"protocols,omitempty"`

	// **参数解释**：冒号分隔的自定义安全策略的加密套件。 支持多值查询，查询条件格式：*ciphers=xxx&ciphers=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Ciphers *[]string `json:"ciphers,omitempty"`
}

func (o ListSecurityPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityPoliciesRequest", string(data)}, " ")
}
