package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventSubResponse Response Object
type CreateEventSubResponse struct {

	// 订阅ID
	Id *string `json:"id,omitempty"`

	// 订阅名称
	Name *string `json:"name,omitempty"`

	// 事件源类型
	SourceType *string `json:"source_type,omitempty"`

	// 事件源ID
	SourceId *string `json:"source_id,omitempty"`

	// 事件类别
	Category *string `json:"category,omitempty"`

	// 事件级别
	Severity *string `json:"severity,omitempty"`

	// 事件标签
	Tag *string `json:"tag,omitempty"`

	// 是否开启订阅 1为开启，0为关闭
	Enable *int32 `json:"enable,omitempty"`

	// 租户凭证ID
	ProjectId *string `json:"project_id,omitempty"`

	// 所属服务
	NameSpace *string `json:"name_space,omitempty"`

	// 消息通知主题地址
	NotificationTarget *string `json:"notification_target,omitempty"`

	// 消息通知主题名称
	NotificationTargetName *string `json:"notification_target_name,omitempty"`

	// 消息通知类型
	NotificationTargetType *string `json:"notification_target_type,omitempty"`

	// 语言
	Language *string `json:"language,omitempty"`

	// 时区
	TimeZone       *string `json:"time_zone,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEventSubResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSubResponse struct{}"
	}

	return strings.Join([]string{"CreateEventSubResponse", string(data)}, " ")
}
