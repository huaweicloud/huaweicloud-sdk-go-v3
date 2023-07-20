package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlowLog 流日志详情
type FlowLog struct {

	// 流日志ID
	Id string `json:"id"`

	// 流日志名称
	Name string `json:"name"`

	// 流日志描述
	Description *string `json:"description,omitempty"`

	// 流日志任务创建者项目ID
	ProjectId string `json:"project_id"`

	// 采集的资源类型:attachment
	ResourceType string `json:"resource_type"`

	// 采集的资源ID
	ResourceId string `json:"resource_id"`

	// 日志组ID
	LogGroupId string `json:"log_group_id"`

	// 日志流ID
	LogStreamId string `json:"log_stream_id"`

	// 流日志存储类型
	LogStoreType string `json:"log_store_type"`

	// 日志聚合时间，单位s，取值范围：60-600
	LogAggregationInterval *int32 `json:"log_aggregation_interval,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 日志资源状态:pending|available|modifying|deleting|deleted|failed
	State string `json:"state"`

	// 日志开关:true|false
	Enabled bool `json:"enabled"`
}

func (o FlowLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowLog struct{}"
	}

	return strings.Join([]string{"FlowLog", string(data)}, " ")
}
