package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmnTopicsResponse Response Object
type ListSmnTopicsResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 请求的唯一标识ID
	RequestId *string `json:"request_id,omitempty"`

	// 返回的Topic个数
	TopicCount *int32 `json:"topic_count,omitempty"`

	// 主题列表
	Topics         *[]SmnTopicInfo `json:"topics,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListSmnTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmnTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListSmnTopicsResponse", string(data)}, " ")
}
