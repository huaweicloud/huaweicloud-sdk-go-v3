package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新黑白名单规则body
type UpdateWhiteBlackIpRuleRequestBody struct {
	// 黑白名单规则名称

	Name string `json:"name"`
	// 黑白名单ip地址，需要输入标准的ip地址或地址段，例如：42.123.120.66或42.123.120.0/16

	Addr string `json:"addr"`
	// 黑白名单规则描述

	Description *string `json:"description,omitempty"`
	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录

	White int32 `json:"white"`
}

func (o UpdateWhiteBlackIpRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhiteBlackIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWhiteBlackIpRuleRequestBody", string(data)}, " ")
}
