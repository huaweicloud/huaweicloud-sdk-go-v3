package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScalingPolicyBySession 基于会话的弹性策略
type ScalingPolicyBySession struct {

	// 分组的总会话使用率(达到改阈值后扩容)
	SessionUsageThreshold int32 `json:"session_usage_threshold"`

	// 给定时间内无会话连接的的实例进行释放
	ShrinkAfterSessionIdleMinutes int32 `json:"shrink_after_session_idle_minutes"`
}

func (o ScalingPolicyBySession) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPolicyBySession struct{}"
	}

	return strings.Join([]string{"ScalingPolicyBySession", string(data)}, " ")
}
