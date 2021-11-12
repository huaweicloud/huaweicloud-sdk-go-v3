package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThrottleBindingResp struct {
	// API的发布记录编号

	PublishId *string `json:"publish_id,omitempty"`
	// 策略作用域，取值如下： - 1：整个API - 2： 单个用户 - 3：单个APP  目前只支持1

	Scope *int32 `json:"scope,omitempty"`
	// 流控策略的ID

	StrategyId *string `json:"strategy_id,omitempty"`
	// 绑定时间

	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty"`
	// 绑定关系的ID

	Id *string `json:"id,omitempty"`
}

func (o ThrottleBindingResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleBindingResp struct{}"
	}

	return strings.Join([]string{"ThrottleBindingResp", string(data)}, " ")
}
