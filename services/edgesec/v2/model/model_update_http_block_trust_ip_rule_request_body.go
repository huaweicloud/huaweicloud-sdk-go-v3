package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHttpBlockTrustIpRuleRequestBody struct {

	// 规则名称
	Name string `json:"name"`

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 规则开关： - 0：关闭 - 1：开启
	Status *int32 `json:"status,omitempty"`

	// ip地址/地址段；ip地址/地址段或者ip地址组id至少有一个
	Addr *string `json:"addr,omitempty"`

	// - 0：拦截 - 1：放行 - 2：仅记录
	White int32 `json:"white"`

	// 攻击惩罚规则id
	FollowedActionId *string `json:"followed_action_id,omitempty"`

	// ip地址组id；ip地址/地址段或者ip地址组id至少有一个
	IpGroupId *string `json:"ip_group_id,omitempty"`
}

func (o UpdateHttpBlockTrustIpRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpBlockTrustIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpBlockTrustIpRuleRequestBody", string(data)}, " ")
}
