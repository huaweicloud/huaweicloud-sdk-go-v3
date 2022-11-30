package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// smn主题相关信息
type SmnTopics struct {

	// Topic的显示名，推送邮件消息时，作为邮件发件人显示。显示名的长度为192byte或64个中文。默认值为空。
	DisplayName *string `json:"display_name,omitempty"`

	// 创建topic的名字。Topic名称只能包含大写字母、小写字母、数字、-和_，且必须由大写字母、小写字母或数字开头，长度为1到255个字符。
	Name string `json:"name"`

	// SMN消息推送策略。取值为0或1
	PushPolicy int32 `json:"push_policy"`

	// topic中订阅者的状态。0:主题已删除或主题下订阅列表为空。1:主题下的订阅列表存在状态为“已订阅”的订阅信息。2:主题下的订阅信息状态处于“未订阅”或“已取消”。
	Status *int32 `json:"status,omitempty"`

	// Topic的唯一的资源标识。
	TopicUrn string `json:"topic_urn"`
}

func (o SmnTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnTopics struct{}"
	}

	return strings.Join([]string{"SmnTopics", string(data)}, " ")
}
