package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTaskHistoryResponse Response Object
type ListScheduledTaskHistoryResponse struct {

	// 定时运维历史记录总数量
	Count *int64 `json:"count,omitempty"`

	// 分页标记
	NextMaker *string `json:"next_maker,omitempty"`

	// 定时运维历史记录列表
	ScheduledTaskHistoryList *[]ScheduledTaskHistory `json:"scheduled_task_history_list,omitempty"`
	HttpStatusCode           int                     `json:"-"`
}

func (o ListScheduledTaskHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTaskHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledTaskHistoryResponse", string(data)}, " ")
}
