package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmPolicyIpReputationMapRequest Request Object
type ConfirmPolicyIpReputationMapRequest struct {

	// **参数解释：** 语言 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Lang string `json:"lang"`

	// **参数解释：** 语言 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Type string `json:"type"`
}

func (o ConfirmPolicyIpReputationMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmPolicyIpReputationMapRequest struct{}"
	}

	return strings.Join([]string{"ConfirmPolicyIpReputationMapRequest", string(data)}, " ")
}
