package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNosqlTaskListResponse Response Object
type ListNosqlTaskListResponse struct {

	// 记录总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 任务详情。
	Schedules      *[]ScheduleDetailInfo `json:"schedules,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListNosqlTaskListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNosqlTaskListResponse struct{}"
	}

	return strings.Join([]string{"ListNosqlTaskListResponse", string(data)}, " ")
}
