package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTaskResponse Response Object
type ListScheduledTaskResponse struct {

	// 定时运维总数量
	Count *int64 `json:"count,omitempty"`

	// 下次起始点
	NextMaker *string `json:"next_maker,omitempty"`

	// 定时运维列表
	ScheduledTasks *[]ScheduledTaskBasicInfo `json:"scheduled_tasks,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListScheduledTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTaskResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledTaskResponse", string(data)}, " ")
}
