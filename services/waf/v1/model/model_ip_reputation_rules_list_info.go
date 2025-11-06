package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpReputationRulesListInfo struct {

	// **参数解释：** 规则名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** Rule ID. **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** Policy ID. **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Policyid *string `json:"policyid,omitempty"`

	// **参数解释：** 信誉类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释：** 规则描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 信誉类型的IDC数据中心 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tags *[]string `json:"tags,omitempty"`

	// **参数解释：** 规则状态，0：关闭，1：开启 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	Action *IpReputationRulesListInfoAction `json:"action,omitempty"`
}

func (o IpReputationRulesListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpReputationRulesListInfo struct{}"
	}

	return strings.Join([]string{"IpReputationRulesListInfo", string(data)}, " ")
}
