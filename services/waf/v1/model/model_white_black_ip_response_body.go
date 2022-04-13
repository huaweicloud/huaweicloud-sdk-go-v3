package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WhiteBlackIpResponseBody struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 策略id

	Policyid *string `json:"policyid,omitempty"`
	// 创建规则的时间戳

	Timestamp *int64 `json:"timestamp,omitempty"`
	// 规则描述

	Description *string `json:"description,omitempty"`
	// 规则状态，0：关闭，1：开启

	Status *int32 `json:"status,omitempty"`
	// 黑白名单

	Addr *string `json:"addr,omitempty"`
	// 防护动作：  - 0拦截  - 1放行  - 2仅记录

	White *int32 `json:"white,omitempty"`
}

func (o WhiteBlackIpResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhiteBlackIpResponseBody struct{}"
	}

	return strings.Join([]string{"WhiteBlackIpResponseBody", string(data)}, " ")
}
