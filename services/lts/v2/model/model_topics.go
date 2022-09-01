package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Topics struct {

	// 主题名称
	Name string `json:"name" xml:"name"`

	// Topic的唯一的资源标识。
	TopicUrn string `json:"topic_urn" xml:"topic_urn"`

	// Topic的显示名，推送邮件消息时，作为邮件发件人显示
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 消息推送的策略
	PushPolicy *int32 `json:"push_policy,omitempty" xml:"push_policy"`
}

func (o Topics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Topics struct{}"
	}

	return strings.Join([]string{"Topics", string(data)}, " ")
}
