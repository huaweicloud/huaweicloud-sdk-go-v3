package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFullRuleRequest Request Object
type UpdateFullRuleRequest struct {

	// **参数解释：** 加速域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	Body *FullUpdateRulesRequest `json:"body,omitempty"`
}

func (o UpdateFullRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFullRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateFullRuleRequest", string(data)}, " ")
}
