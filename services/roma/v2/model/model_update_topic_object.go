package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTopicObject struct {

	// topic名称。
	Name string `json:"name" xml:"name"`

	// 消息老化时间。默认值为72。取值范围1~720，单位小时。
	RetentionTime *int32 `json:"retention_time,omitempty" xml:"retention_time"`

	// 是否开启同步复制。
	SyncReplication *bool `json:"sync_replication,omitempty" xml:"sync_replication"`

	// 是否使用同步落盘。
	SyncMessageFlush *bool `json:"sync_message_flush,omitempty" xml:"sync_message_flush"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 敏感字段。
	SensitiveWord *string `json:"sensitive_word,omitempty" xml:"sensitive_word"`
}

func (o UpdateTopicObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicObject struct{}"
	}

	return strings.Join([]string{"UpdateTopicObject", string(data)}, " ")
}
