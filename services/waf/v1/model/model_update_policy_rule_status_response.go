package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyRuleStatusResponse Response Object
type UpdatePolicyRuleStatusResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 规则创建时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 功能状态标识，用于指定对应功能的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status         *int32 `json:"status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdatePolicyRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRuleStatusResponse", string(data)}, " ")
}
