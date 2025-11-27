package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInspectionReportRequest Request Object
type ListInspectionReportRequest struct {

	// 开始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`

	// 数据库类型。支持MySQL、TaurusDB、GaussDB、MariaDB。
	DatastoreType string `json:"datastore_type"`

	// 健康等级
	HealthRank *string `json:"health_rank,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序顺序（true:正序, false:逆序）
	Asc *bool `json:"asc,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInspectionReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInspectionReportRequest struct{}"
	}

	return strings.Join([]string{"ListInspectionReportRequest", string(data)}, " ")
}
