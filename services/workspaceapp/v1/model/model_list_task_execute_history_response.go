package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskExecuteHistoryResponse Response Object
type ListTaskExecuteHistoryResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 定时任务执行记录列表，返回列表条目数量上限为分页的最大上限值。
	Items          *[]ScheduleTaskExecuteHistory `json:"items,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListTaskExecuteHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskExecuteHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListTaskExecuteHistoryResponse", string(data)}, " ")
}
