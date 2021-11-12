package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateWhiteblackipRuleResponse struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 黑白名单地址

	Addr *string `json:"addr,omitempty"`
	// 黑白名单规则描述

	Description *string `json:"description,omitempty"`
	// 设置的ip地址类型，1放行，0拦截，2仅记录

	White          *int32 `json:"white,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateWhiteblackipRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhiteblackipRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateWhiteblackipRuleResponse", string(data)}, " ")
}
