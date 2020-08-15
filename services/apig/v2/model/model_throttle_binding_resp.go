/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

type ThrottleBindingResp struct {
	// API的发布记录编号
	PublishId string `json:"publish_id,omitempty"`
	// 策略作用域，取值如下： - 1：整个API - 2： 单个用户 - 3：单个APP  目前只支持1
	Scope int32 `json:"scope,omitempty"`
	// 流控策略的ID
	StrategyId string `json:"strategy_id,omitempty"`
	// 绑定时间
	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty"`
	// 绑定关系的ID
	Id string `json:"id,omitempty"`
}

func (o ThrottleBindingResp) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ThrottleBindingResp", string(data)}, " ")
}
