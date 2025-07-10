package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartChatJobResponse Response Object
type ListSmartChatJobResponse struct {

	// 数字人智能交互对话任务总数。
	Count *int32 `json:"count,omitempty"`

	// 数字人智能交互对话任务列表。
	SmartChatJobs *[]SmartChatJobBaseInfo `json:"smart_chat_jobs,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSmartChatJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartChatJobResponse struct{}"
	}

	return strings.Join([]string{"ListSmartChatJobResponse", string(data)}, " ")
}
