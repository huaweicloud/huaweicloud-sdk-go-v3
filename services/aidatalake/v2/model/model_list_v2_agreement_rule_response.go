package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2AgreementRuleResponse Response Object
type ListV2AgreementRuleResponse struct {

	// **参数解释**：系统协议列表，每个元素为一个V2AgreementRule对象，包含协议名称、展示名称、版本、链接等信息。
	AgreementRules *[]V2AgreementRule `json:"agreement_rules,omitempty"`

	// **参数解释**：系统协议总数。 **取值范围**：0~1000。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListV2AgreementRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2AgreementRuleResponse struct{}"
	}

	return strings.Join([]string{"ListV2AgreementRuleResponse", string(data)}, " ")
}
