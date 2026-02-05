package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScalingPolicyByResource 基于资源的弹性策略。
type ScalingPolicyByResource struct {

	// 分组的总cpu使用率(达到阈值后扩容)。
	CpuUsageThreshold *int32 `json:"cpu_usage_threshold,omitempty"`

	// 分组的总mem使用率(达到改阈值后扩容)。
	MemUsageThreshold *int32 `json:"mem_usage_threshold,omitempty"`

	// 分组的总显存使用率(达到改阈值后扩容)。
	GpuUsageThreshold *int32 `json:"gpu_usage_threshold,omitempty"`
}

func (o ScalingPolicyByResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPolicyByResource struct{}"
	}

	return strings.Join([]string{"ScalingPolicyByResource", string(data)}, " ")
}
