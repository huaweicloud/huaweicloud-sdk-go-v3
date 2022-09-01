package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateWhiteblackipRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty" xml:"id"`

	// 黑白名单规则名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 黑白名单Ip/IP段
	Addr *string `json:"addr,omitempty" xml:"addr"`

	// 黑白名单规则描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White *int32 `json:"white,omitempty" xml:"white"`

	IpGroup        *IpGroup `json:"ip_group,omitempty" xml:"ip_group"`
	HttpStatusCode int      `json:"-"`
}

func (o UpdateWhiteblackipRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhiteblackipRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateWhiteblackipRuleResponse", string(data)}, " ")
}
