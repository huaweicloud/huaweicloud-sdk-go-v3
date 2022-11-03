package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新健康检查的详细信息。
type UpdateHealthCheckOption struct {
	Protocol *HealthCheckProtocol `json:"protocol,omitempty"`

	// 健康检查的端口。
	Port *int32 `json:"port,omitempty"`

	// 健康检查的时间间隔，单位为秒。
	Interval *int32 `json:"interval,omitempty"`

	// 健康检查的超时时间，单位为秒。建议该值小于interval的值。
	Timeout *int32 `json:"timeout,omitempty"`

	// 最大重试次数。将终端节点的状态从“健康”设置为“不健康”或从“不健康”设置为“健康”所需的连续健康检查次数。
	MaxRetries *int32 `json:"max_retries,omitempty"`

	// 是否开启健康检查。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o UpdateHealthCheckOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthCheckOption struct{}"
	}

	return strings.Join([]string{"UpdateHealthCheckOption", string(data)}, " ")
}
