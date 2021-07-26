package model

import (
	"encoding/json"

	"strings"
)

// 创建黑白名单规则body
type UpdateWhiteBlackIpRuleRequestBody struct {
	// 黑白名单地址

	Addr string `json:"addr"`
	// 黑白名单规则描述

	Description *string `json:"description,omitempty"`
	// 设置的ip地址类型，1放行，0拦截，2仅记录

	White *int32 `json:"white,omitempty"`
}

func (o UpdateWhiteBlackIpRuleRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateWhiteBlackIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWhiteBlackIpRuleRequestBody", string(data)}, " ")
}
