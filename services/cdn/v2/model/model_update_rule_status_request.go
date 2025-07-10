package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleStatusRequest **参数解释：** 规则配置 **约束限制：** 不涉及
type UpdateRuleStatusRequest struct {

	// **参数解释：** 规则ID，可以通过查询规则引擎列表接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`

	// **参数解释：** 是否开启规则 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭 **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释：** 此条规则的优先级，数值越大，优先级越高 **约束限制：** 优先级不能相同 **取值范围：** 1-100 **默认取值：** 不涉及
	Priority *int32 `json:"priority,omitempty"`
}

func (o UpdateRuleStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateRuleStatusRequest", string(data)}, " ")
}
