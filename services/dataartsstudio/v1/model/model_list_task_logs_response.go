package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskLogsResponse Response Object
type ListTaskLogsResponse struct {

	// 请求是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息。
	Message *string `json:"message,omitempty"`

	// 作业日志内容列表。
	TaskLogsContentList *[]TaskLogsContent `json:"task_logs_content_list,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o ListTaskLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskLogsResponse struct{}"
	}

	return strings.Join([]string{"ListTaskLogsResponse", string(data)}, " ")
}
