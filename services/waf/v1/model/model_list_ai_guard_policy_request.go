package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiGuardPolicyRequest Request Object
type ListAiGuardPolicyRequest struct {

	// **参数解释：** 最大数量 **约束限制：** 不涉及 **取值范围：** 1-100 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 0-10000 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 策略名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`
}

func (o ListAiGuardPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiGuardPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListAiGuardPolicyRequest", string(data)}, " ")
}
