package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoutePolicyDetail 路由策略。
type RoutePolicyDetail struct {

	// 单台服务器最大的连接会话数。
	MaxSession *int32 `json:"max_session,omitempty"`

	// cpu使用率阈值，单位为%。
	CpuThreshold *int32 `json:"cpu_threshold,omitempty"`

	// gpu显存使用率阈值，单位为%。
	GpuThreshold *int32 `json:"gpu_threshold,omitempty"`

	// 内存使用率阈值，单位为%。
	MemThreshold *int32 `json:"mem_threshold,omitempty"`
}

func (o RoutePolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutePolicyDetail struct{}"
	}

	return strings.Join([]string{"RoutePolicyDetail", string(data)}, " ")
}
