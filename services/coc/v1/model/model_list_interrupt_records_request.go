package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInterruptRecordsRequest Request Object
type ListInterruptRecordsRequest struct {

	// SLO的ID
	SloId string `json:"slo_id"`

	// 分页指针
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 资源ID
	SourceId *string `json:"source_id,omitempty"`

	// regionId
	RegionId *string `json:"region_id,omitempty"`

	// 时间范围 - 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 时间范围 - 结束时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ListInterruptRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInterruptRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListInterruptRecordsRequest", string(data)}, " ")
}
