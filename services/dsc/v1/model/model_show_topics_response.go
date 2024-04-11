package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopicsResponse Response Object
type ShowTopicsResponse struct {

	// DSC告警主题ID（非消息通知服务主题ID）
	Id *string `json:"id,omitempty"`

	// 默认消息通知主题的唯一资源标识符
	DefaultTopicUrn *string `json:"default_topic_urn,omitempty"`

	// 已确认的消息通知主题数量
	TopicCount *int32 `json:"topic_count,omitempty"`

	// 已确认的消息通知主题列表
	Topics         *[]TopicBean `json:"topics,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicsResponse struct{}"
	}

	return strings.Join([]string{"ShowTopicsResponse", string(data)}, " ")
}
