package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SourceCommunityMqParameters struct {

	// 实例名称，仅dms的rockectMq需要该字段
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例ID，仅dms的rockectMq需要该字段
	InstanceId *string `json:"instance_id,omitempty"`

	// rockectMq连接地址
	Addr string `json:"addr"`

	// 消费组
	Group string `json:"group"`

	// topic名称
	Topic string `json:"topic"`

	// 标签
	Tag *string `json:"tag,omitempty"`

	// 虚拟云id
	VpcId string `json:"vpc_id"`

	// 子网id
	SubnetId string `json:"subnet_id"`

	// 开启SSL
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// ACL访问控制
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// 用户名
	AccessKey *string `json:"access_key,omitempty"`

	// 密码
	SecretKey *string `json:"secret_key,omitempty"`

	// 消息类型
	MessageType *SourceCommunityMqParametersMessageType `json:"message_type,omitempty"`

	// 消费超时时间
	ConsumeTimeout *int32 `json:"consume_timeout,omitempty"`

	// 线程消费数
	ConsumerThreadNums *int32 `json:"consumer_thread_nums,omitempty"`

	// 批量消费最大消息数
	ConsumerBatchMaxSize *int32 `json:"consumer_batch_max_size,omitempty"`

	// 最大重试次数，-1表示一直重试
	MaxReconsumeTimes *int32 `json:"max_reconsume_times,omitempty"`

	// 重试间隔，单位ms
	SuspendCurrentQueueTimeMillis *int32 `json:"suspend_current_queue_time_millis,omitempty"`
}

func (o SourceCommunityMqParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceCommunityMqParameters struct{}"
	}

	return strings.Join([]string{"SourceCommunityMqParameters", string(data)}, " ")
}

type SourceCommunityMqParametersMessageType struct {
	value string
}

type SourceCommunityMqParametersMessageTypeEnum struct {
	NORMAL SourceCommunityMqParametersMessageType
}

func GetSourceCommunityMqParametersMessageTypeEnum() SourceCommunityMqParametersMessageTypeEnum {
	return SourceCommunityMqParametersMessageTypeEnum{
		NORMAL: SourceCommunityMqParametersMessageType{
			value: "normal",
		},
	}
}

func (c SourceCommunityMqParametersMessageType) Value() string {
	return c.value
}

func (c SourceCommunityMqParametersMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceCommunityMqParametersMessageType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
