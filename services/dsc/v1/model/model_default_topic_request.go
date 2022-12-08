package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DefaultTopicRequest struct {

	// 告警主题id
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 告警通知状态
	Status *int32 `json:"status,omitempty"`

	// 消息通知主题的唯一资源标识符
	TopicUrn *string `json:"topic_urn,omitempty"`
}

func (o DefaultTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefaultTopicRequest struct{}"
	}

	return strings.Join([]string{"DefaultTopicRequest", string(data)}, " ")
}
