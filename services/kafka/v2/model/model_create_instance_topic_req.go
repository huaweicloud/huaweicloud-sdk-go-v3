package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstanceTopicReq struct {

	// Topic名称，长度为3-200，以字母开头且只支持大小写字母、中横线、下划线、点以及数字。
	Id string `json:"id"`

	// 副本数，配置数据的可靠性。 取值范围：1-3。
	Replication *int32 `json:"replication,omitempty"`

	// 是否使用同步落盘。默认值为false。同步落盘会导致性能降低。
	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`

	// Topic分区数，设置消费的并发数。 取值范围：1-200。
	Partition *int32 `json:"partition,omitempty"`

	// 是否开启同步复制，开启后，客户端生产消息时相应的也要设置acks=-1，否则不生效，默认关闭。
	SyncReplication *bool `json:"sync_replication,omitempty"`

	// 消息老化时间。默认值为72。 取值范围1-720，单位小时。
	RetentionTime *int32 `json:"retention_time,omitempty"`

	// Topic配置
	TopicOtherConfigs *[]CreateInstanceTopicReqTopicOtherConfigs `json:"topic_other_configs,omitempty"`

	// Topic描述
	TopicDesc *string `json:"topic_desc,omitempty"`
}

func (o CreateInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceTopicReq", string(data)}, " ")
}
