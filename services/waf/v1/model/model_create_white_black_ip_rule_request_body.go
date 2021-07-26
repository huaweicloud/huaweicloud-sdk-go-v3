package model

import (
	"encoding/json"

	"strings"
)

// 创建黑白名单规则body
type CreateWhiteBlackIpRuleRequestBody struct {
	// 黑白名单地址

	Addr string `json:"addr"`
	// 黑白名单规则描述

	Description *string `json:"description,omitempty"`
	// 设置ip的防护动作，1放行，0拦截，2仅记录

	White *int32 `json:"white,omitempty"`
}

func (o CreateWhiteBlackIpRuleRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWhiteBlackIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWhiteBlackIpRuleRequestBody", string(data)}, " ")
}
