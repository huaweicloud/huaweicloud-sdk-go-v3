package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleNewRequest Request Object
type DeleteRuleNewRequest struct {

	// **参数解释：** 加速域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释：** 规则ID，可以通过查询规则引擎列表接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`
}

func (o DeleteRuleNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleNewRequest struct{}"
	}

	return strings.Join([]string{"DeleteRuleNewRequest", string(data)}, " ")
}
