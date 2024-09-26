package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpPolicyAction 操作
type HttpPolicyAction struct {

	// 防护等级
	Category *string `json:"category,omitempty"`

	// 攻击惩罚规则ID
	FollowedActionId *string `json:"followed_action_id,omitempty"`
}

func (o HttpPolicyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpPolicyAction struct{}"
	}

	return strings.Join([]string{"HttpPolicyAction", string(data)}, " ")
}
