package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNotificationTopicsResponse struct {

	// request_id
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// topics数量
	TopicCount *int32 `json:"topic_count,omitempty" xml:"topic_count"`

	// 主题信息
	Topics         *[]Topics `json:"topics,omitempty" xml:"topics"`
	HttpStatusCode int       `json:"-"`
}

func (o ListNotificationTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationTopicsResponse", string(data)}, " ")
}
