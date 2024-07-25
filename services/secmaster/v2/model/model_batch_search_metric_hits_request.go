package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSearchMetricHitsRequest Request Object
type BatchSearchMetricHitsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 查询指标的时间范围，ISO8601格式，例如：2007-03-01T13:00:00Z/2008-05-11T15:30:00Z或2007-03-01T13:00:00Z/P1Y2M10DT2H30M或P1Y2M10DT2H30M/2008-05-11T15:30:00Z
	Timespan *string `json:"timespan,omitempty"`

	// 是否启用缓存，默认true, 禁用缓存 false
	Cache *bool `json:"cache,omitempty"`

	Body *BatchSearchMetricHitsRequestBody `json:"body,omitempty"`
}

func (o BatchSearchMetricHitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSearchMetricHitsRequest struct{}"
	}

	return strings.Join([]string{"BatchSearchMetricHitsRequest", string(data)}, " ")
}
