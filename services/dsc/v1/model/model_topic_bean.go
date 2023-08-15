package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopicBean struct {

	// 消息通知主题名称
	Name *string `json:"name,omitempty"`

	// 消息通知主题的唯一资源标识符
	TopicUrn *string `json:"topic_urn,omitempty"`
}

func (o TopicBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicBean struct{}"
	}

	return strings.Join([]string{"TopicBean", string(data)}, " ")
}
