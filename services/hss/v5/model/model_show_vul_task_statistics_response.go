package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulTaskStatisticsResponse Response Object
type ShowVulTaskStatisticsResponse struct {

	// 未读的已完成的任务
	UnreadTaskNum  *int32 `json:"unread_task_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVulTaskStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulTaskStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowVulTaskStatisticsResponse", string(data)}, " ")
}
