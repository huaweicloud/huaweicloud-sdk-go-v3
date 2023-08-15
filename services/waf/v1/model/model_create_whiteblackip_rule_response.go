package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWhiteblackipRuleResponse Response Object
type CreateWhiteblackipRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 黑白名单规则名称
	Name *string `json:"name,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 黑白名单ip地址，需要输入标准的ip地址或地址段，例如：42.123.120.66或42.123.120.0/16
	Addr *string `json:"addr,omitempty"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White *int32 `json:"white,omitempty"`

	IpGroup *IpGroup `json:"ip_group,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 创建规则的时间戳,13位毫秒时间戳
	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateWhiteblackipRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhiteblackipRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateWhiteblackipRuleResponse", string(data)}, " ")
}
