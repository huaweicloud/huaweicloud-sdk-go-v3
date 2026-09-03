package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmnTopicInfo SMN主题信息
type SmnTopicInfo struct {

	// Topic的唯一的资源标识
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 创建topic的名字
	Name *string `json:"name,omitempty"`

	// Topic的显示名，推送邮件消息时，作为邮件发件人显示
	DisplayName *string `json:"display_name,omitempty"`

	// 消息推送的策略，0表示发送失败保留到失败队列，1表示直接丢弃发送失败的消息
	PushPolicy *string `json:"push_policy,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 主题ID
	TopicId *string `json:"topic_id,omitempty"`
}

func (o SmnTopicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnTopicInfo struct{}"
	}

	return strings.Join([]string{"SmnTopicInfo", string(data)}, " ")
}
