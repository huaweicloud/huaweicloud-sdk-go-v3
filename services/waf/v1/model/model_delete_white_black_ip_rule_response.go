package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteWhiteBlackIpRuleResponse struct {
	// 黑白名单规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 黑白名单规则名称

	Name *string `json:"name,omitempty"`
	// 删除规则时间，13位毫秒时间戳

	Timestamp *int64 `json:"timestamp,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 规则状态，0：关闭，1：开启

	Status *int32 `json:"status,omitempty"`
	// 黑白名单ip地址，标准的ip地址或地址段，例如：42.123.120.66或42.123.120.0/16

	Addr *string `json:"addr,omitempty"`
	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录

	White          *int32 `json:"white,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteWhiteBlackIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWhiteBlackIpRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteWhiteBlackIpRuleResponse", string(data)}, " ")
}
