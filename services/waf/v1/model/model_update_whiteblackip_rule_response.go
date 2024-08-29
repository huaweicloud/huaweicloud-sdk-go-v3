package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWhiteblackipRuleResponse Response Object
type UpdateWhiteblackipRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 黑白名单规则名称
	Name *string `json:"name,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 黑白名单Ip/IP段
	Addr *string `json:"addr,omitempty"`

	// 黑白名单规则描述
	Description *string `json:"description,omitempty"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White *int32 `json:"white,omitempty"`

	// 生效模式，默认为permanent（立即生效）
	TimeMode *string `json:"time_mode,omitempty"`

	IpGroup        *IpGroup `json:"ip_group,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o UpdateWhiteblackipRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhiteblackipRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateWhiteblackipRuleResponse", string(data)}, " ")
}
