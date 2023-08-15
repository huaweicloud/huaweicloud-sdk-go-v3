package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthCheckDetail 健康检查实例。
type HealthCheckDetail struct {

	// 健康检查ID。
	Id *string `json:"id,omitempty"`

	// 终端节点组ID。
	EndpointGroupId *string `json:"endpoint_group_id,omitempty"`

	Protocol *HealthCheckProtocol `json:"protocol,omitempty"`

	Status *ConfigStatus `json:"status,omitempty"`

	// 健康检查的端口。
	Port *int32 `json:"port,omitempty"`

	// 健康检查的时间间隔，单位为秒。
	Interval *int32 `json:"interval,omitempty"`

	// 健康检查的超时时间，单位为秒。建议该值小于interval的值。
	Timeout *int32 `json:"timeout,omitempty"`

	// 最大重试次数。将终端节点的状态从“健康”设置为“不健康”或从“不健康”设置为“健康”所需的连续健康检查次数。
	MaxRetries *int32 `json:"max_retries,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 是否开启健康检查。
	Enabled *bool `json:"enabled,omitempty"`

	FrozenInfo *FrozenInfo `json:"frozen_info,omitempty"`
}

func (o HealthCheckDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCheckDetail struct{}"
	}

	return strings.Join([]string{"HealthCheckDetail", string(data)}, " ")
}
