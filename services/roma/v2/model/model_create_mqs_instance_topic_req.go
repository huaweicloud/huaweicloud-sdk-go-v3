package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateMqsInstanceTopicReq struct {
	// topic名称，以字母开头，仅能包含数字,字母,下划线(_)，中划线（-）,长度3-200字符。

	Name string `json:"name"`
	// 集成应用key。

	AppId string `json:"app_id"`
	// topic分区数，设置消费的并发数。  取值范围：1-50。  默认值：3。

	Partition *int32 `json:"partition,omitempty"`
	// 副本数，配置数据的可靠性。  取值范围：1-3。  默认值：3。  > 体验版实例的副本数只能为1。

	Replication *int32 `json:"replication,omitempty"`
	// 权限类型。   - all：发布+订阅   - pub：发布   - sub：订阅

	AccessPolicy CreateMqsInstanceTopicReqAccessPolicy `json:"access_policy"`
	// 是否使用同步落盘。默认值为false。同步落盘会导致性能降低。

	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`
	// 是否开启同步复制，开启后，客户端生产消息时相应的也要设置acks=-1，否则不生效,默认关闭。

	SyncReplication *bool `json:"sync_replication,omitempty"`
	// 消息老化时间。默认值为72。取值范围1~720，单位小时。

	RetentionTime *int32 `json:"retention_time,omitempty"`
	// 权限类型对应的标签。  当权限类型是all时，发布和订阅的标签用符号“&”隔开。  当有多个标签时，标签用符号“||”隔开。

	Tag *string `json:"tag,omitempty"`
	// 描述。长度0-1000字符。

	Description *string `json:"description,omitempty"`
	// 敏感字段。  当有多个敏感字段时，敏感字段用符号“||”隔开。

	SensitiveWord *string `json:"sensitive_word,omitempty"`
}

func (o CreateMqsInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMqsInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"CreateMqsInstanceTopicReq", string(data)}, " ")
}

type CreateMqsInstanceTopicReqAccessPolicy struct {
	value string
}

type CreateMqsInstanceTopicReqAccessPolicyEnum struct {
	ALL CreateMqsInstanceTopicReqAccessPolicy
	PUB CreateMqsInstanceTopicReqAccessPolicy
	SUB CreateMqsInstanceTopicReqAccessPolicy
}

func GetCreateMqsInstanceTopicReqAccessPolicyEnum() CreateMqsInstanceTopicReqAccessPolicyEnum {
	return CreateMqsInstanceTopicReqAccessPolicyEnum{
		ALL: CreateMqsInstanceTopicReqAccessPolicy{
			value: "all",
		},
		PUB: CreateMqsInstanceTopicReqAccessPolicy{
			value: "pub",
		},
		SUB: CreateMqsInstanceTopicReqAccessPolicy{
			value: "sub",
		},
	}
}

func (c CreateMqsInstanceTopicReqAccessPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMqsInstanceTopicReqAccessPolicy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
