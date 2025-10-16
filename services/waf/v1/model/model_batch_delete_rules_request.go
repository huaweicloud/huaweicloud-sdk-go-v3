package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRulesRequest Request Object
type BatchDeleteRulesRequest struct {

	// **参数解释：** 需要删除的规则类型,目前支持cc,custom,whiteblackip,geoip,ip-reputation,antitamper,antileakage,ignore,privacy **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleType string `json:"rule_type"`

	Body *PolicyRuleIdRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteRulesRequest", string(data)}, " ")
}
