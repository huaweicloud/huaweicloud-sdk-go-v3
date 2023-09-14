package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Record 事件通知记录。
type Record struct {

	// 凭据名称。
	EventName *string `json:"event_name,omitempty"`

	// 凭据类型  取值 ： COMMON ：通用凭据 RDS ：RDS凭据
	TriggerEventType *string `json:"trigger_event_type,omitempty"`

	// 事件通知记录的创建时间，时间戳，即从1970年1月1日至该时间的总秒数。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 凭据名称。
	SecretName *string `json:"secret_name,omitempty"`

	// 凭据类型  取值 ： COMMON ：通用凭据(默认)。用于应用系统中的各种敏感信息储存。         RDS ：RDS凭据 。专门针对RDS的凭据，用于存储RDS的账号信息。
	SecretType *string `json:"secret_type,omitempty"`

	// 事件通知的对象名称。
	NotificationTargetName *string `json:"notification_target_name,omitempty"`

	// 事件通知的对象ID。
	NotificationTargetId *string `json:"notification_target_id,omitempty"`

	// 凭据的描述信息。
	NotificationContent *string `json:"notification_content,omitempty"`

	// 凭据类型  取值 ： SUCCESS ：事件通知成功。         FAIL ：事件通知失败。         INVALID ：事件通知配置主题信息无效或不正确，无法触发通知。
	NotificationStatus *string `json:"notification_status,omitempty"`
}

func (o Record) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Record struct{}"
	}

	return strings.Join([]string{"Record", string(data)}, " ")
}
