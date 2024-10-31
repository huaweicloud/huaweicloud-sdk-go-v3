package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpBlockTrustIpRuleResponse Response Object
type ShowHttpBlockTrustIpRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则所在策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 规则所在策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则开关： - 0：关闭 - 1：开启
	Status *int32 `json:"status,omitempty"`

	// ip地址/地址段
	Addr *string `json:"addr,omitempty"`

	// - 0：拦截 - 1：放行 - 2：仅记录
	White *int32 `json:"white,omitempty"`

	// 攻击惩罚规则id
	FollowedActionId *string `json:"followed_action_id,omitempty"`

	IpGroup        *HttpIpGroup `json:"ip_group,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowHttpBlockTrustIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpBlockTrustIpRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpBlockTrustIpRuleResponse", string(data)}, " ")
}
