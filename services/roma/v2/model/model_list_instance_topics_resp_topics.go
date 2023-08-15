package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInstanceTopicsRespTopics struct {

	// 是否只更新策略。
	PoliciesOnly *bool `json:"policiesOnly,omitempty"`

	// topic名称。
	Name *string `json:"name,omitempty"`

	// 副本数，配置数据的可靠性。
	Replication *int32 `json:"replication,omitempty"`

	// topic分区数，设置消费的并发数。
	Partition *int32 `json:"partition,omitempty"`

	// 消息老化时间。
	RetentionTime *int32 `json:"retention_time,omitempty"`

	// 是否使用同步落盘。默认值为false。同步落盘会导致性能降低。
	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`

	// 是否开启同步复制，开启后，客户端生产消息时相应的也要设置acks=-1，否则不生效,默认关闭。
	SyncReplication *bool `json:"sync_replication,omitempty"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 允许操作的权限。
	Permissions *[]string `json:"permissions,omitempty"`

	// 其他配置。
	ExternalConfigs *interface{} `json:"external_configs,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 敏感字段。
	SensitiveWord *string `json:"sensitive_word,omitempty"`

	// topic类型。
	TopicType *int32 `json:"topic_type,omitempty"`
}

func (o ListInstanceTopicsRespTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTopicsRespTopics struct{}"
	}

	return strings.Join([]string{"ListInstanceTopicsRespTopics", string(data)}, " ")
}
