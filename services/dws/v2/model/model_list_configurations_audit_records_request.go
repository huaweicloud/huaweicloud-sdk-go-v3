package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsAuditRecordsRequest Request Object
type ListConfigurationsAuditRecordsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 任务时间
	ActionTime *int64 `json:"action_time,omitempty"`

	// 过滤配置信息
	FilterBy *string `json:"filter_by,omitempty"`

	// 过滤条件
	Filter *string `json:"filter,omitempty"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListConfigurationsAuditRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsAuditRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationsAuditRecordsRequest", string(data)}, " ")
}
