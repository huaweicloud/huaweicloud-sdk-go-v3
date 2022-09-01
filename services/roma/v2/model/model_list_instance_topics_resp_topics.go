package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInstanceTopicsRespTopics struct {

	// 是否只更新策略。
	PoliciesOnly *bool `json:"policiesOnly,omitempty" xml:"policiesOnly"`

	// topic名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 副本数，配置数据的可靠性。
	Replication *int32 `json:"replication,omitempty" xml:"replication"`

	// topic分区数，设置消费的并发数。
	Partition *int32 `json:"partition,omitempty" xml:"partition"`

	// 消息老化时间。
	RetentionTime *int32 `json:"retention_time,omitempty" xml:"retention_time"`

	// 是否使用同步落盘。默认值为false。同步落盘会导致性能降低。
	SyncMessageFlush *bool `json:"sync_message_flush,omitempty" xml:"sync_message_flush"`

	// 是否开启同步复制，开启后，客户端生产消息时相应的也要设置acks=-1，否则不生效,默认关闭。
	SyncReplication *bool `json:"sync_replication,omitempty" xml:"sync_replication"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 允许操作的权限。
	Permissions *[]string `json:"permissions,omitempty" xml:"permissions"`

	// 其他配置。
	ExternalConfigs *interface{} `json:"external_configs,omitempty" xml:"external_configs"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 敏感字段。
	SensitiveWord *string `json:"sensitive_word,omitempty" xml:"sensitive_word"`

	// topic类型。
	TopicType *int32 `json:"topic_type,omitempty" xml:"topic_type"`
}

func (o ListInstanceTopicsRespTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTopicsRespTopics struct{}"
	}

	return strings.Join([]string{"ListInstanceTopicsRespTopics", string(data)}, " ")
}
