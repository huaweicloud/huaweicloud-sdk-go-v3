package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmSubResponse Response Object
type UpdateAlarmSubResponse struct {

	// 告警订阅ID
	Id *string `json:"id,omitempty"`

	// 告警订阅名称
	Name *string `json:"name,omitempty"`

	// 是否开启订阅
	Enable *int32 `json:"enable,omitempty"`

	// 告警级别
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// 租户凭证ID
	ProjectId *string `json:"project_id,omitempty"`

	// 所属服务，支持DWS,DLI,DGC,CloudTable,CDM,GES,CSS
	NameSpace *string `json:"name_space,omitempty"`

	// 消息主题地址
	NotificationTarget *string `json:"notification_target,omitempty"`

	// 消息主题名称
	NotificationTargetName *string `json:"notification_target_name,omitempty"`

	// 消息主题类型
	NotificationTargetType *string `json:"notification_target_type,omitempty"`

	// 语言
	Language *string `json:"language,omitempty"`

	// 时区
	TimeZone       *string `json:"time_zone,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAlarmSubResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmSubResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmSubResponse", string(data)}, " ")
}
