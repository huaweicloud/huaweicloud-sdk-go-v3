package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DirectoryTraversalProtectionConfig struct {
	Action *PolicyAction `json:"action,omitempty"`

	// 检测时间范围，单位秒
	TimeRange *int32 `json:"time_range,omitempty"`

	// 请求数阈值，在检测时间内请求数超过该值则触发防护
	NumRequests *int32 `json:"num_requests,omitempty"`

	// 404状态码占比阈值，单位为百分比
	NotFoundStatusPercentage *int32 `json:"not_found_status_percentage,omitempty"`

	// 不存在目录数阈值，在检测时间内探测到不存在目录的数量达到该值时触发防护
	NumNonExistentDirectories *int32 `json:"num_non_existent_directories,omitempty"`

	// 封禁周期，单位秒
	BlockingPeriod *int32 `json:"blocking_period,omitempty"`
}

func (o DirectoryTraversalProtectionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectoryTraversalProtectionConfig struct{}"
	}

	return strings.Join([]string{"DirectoryTraversalProtectionConfig", string(data)}, " ")
}
