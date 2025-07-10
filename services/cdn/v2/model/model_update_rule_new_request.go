package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleNewRequest Request Object
type UpdateRuleNewRequest struct {

	// **参数解释：** 加速域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释：** 规则ID，可以通过查询规则引擎列表接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`

	Body *UpdateRuleRequest `json:"body,omitempty"`
}

func (o UpdateRuleNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleNewRequest struct{}"
	}

	return strings.Join([]string{"UpdateRuleNewRequest", string(data)}, " ")
}
