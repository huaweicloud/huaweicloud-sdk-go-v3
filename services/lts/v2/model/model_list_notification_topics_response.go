package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNotificationTopicsResponse struct {
	// request_id

	RequestId *string `json:"request_id,omitempty"`
	// topics数量

	TopicCount *int32 `json:"topic_count,omitempty"`
	// 主题信息

	Topics         *[]Topics `json:"topics,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListNotificationTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationTopicsResponse", string(data)}, " ")
}
