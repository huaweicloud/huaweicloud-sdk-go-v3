package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestoreTimeResponse Response Object
type ListRestoreTimeResponse struct {

	// 实例可恢复时间段总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 实例可恢复的时间段
	RestorableTimePeriods *[]RestorableTime `json:"restorable_time_periods,omitempty"`
	HttpStatusCode        int               `json:"-"`
}

func (o ListRestoreTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreTimeResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreTimeResponse", string(data)}, " ")
}
