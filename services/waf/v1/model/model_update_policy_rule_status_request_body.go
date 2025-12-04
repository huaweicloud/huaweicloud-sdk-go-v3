package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePolicyRuleStatusRequestBody struct {

	// **参数解释：** 功能状态标识，用于指定对应功能的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`
}

func (o UpdatePolicyRuleStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRuleStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRuleStatusRequestBody", string(data)}, " ")
}
