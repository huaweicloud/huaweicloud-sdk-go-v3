package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoutePolicy 路由策略
type RoutePolicy struct {

	// 单台服务器最大的链接会话数。
	MaxSession *int32 `json:"max_session,omitempty"`

	// cpu使用率阈值，单位为%。
	CpuThreshold *int32 `json:"cpu_threshold,omitempty"`

	// 内存使用率阈值，单位为%。
	MemThreshold *int32 `json:"mem_threshold,omitempty"`
}

func (o RoutePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutePolicy struct{}"
	}

	return strings.Join([]string{"RoutePolicy", string(data)}, " ")
}
