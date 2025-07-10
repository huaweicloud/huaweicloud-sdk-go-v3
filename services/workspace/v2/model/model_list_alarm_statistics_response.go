package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmStatisticsResponse Response Object
type ListAlarmStatisticsResponse struct {

	// 紧急告警记录列表总数。
	CriticalCount *int32 `json:"critical_count,omitempty"`

	// 重要告警记录列表总数。
	MajorCount *int32 `json:"major_count,omitempty"`

	// 次要告警记录列表总数。
	MinorCount *int32 `json:"minor_count,omitempty"`

	// 提示告警记录列表总数。
	InfoCount *int32 `json:"info_count,omitempty"`

	// 告警记录列表总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmStatisticsResponse", string(data)}, " ")
}
