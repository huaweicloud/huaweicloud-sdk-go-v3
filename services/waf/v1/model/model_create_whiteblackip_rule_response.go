package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateWhiteblackipRuleResponse struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 黑白名单

	Ip *string `json:"ip,omitempty"`
	// 类型，0拦截，1放行

	White *int32 `json:"white,omitempty"`
	// 创建规则的时间戳

	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateWhiteblackipRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWhiteblackipRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateWhiteblackipRuleResponse", string(data)}, " ")
}
