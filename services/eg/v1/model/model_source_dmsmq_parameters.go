package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SourceDmsmqParameters struct {

	// 实例名称，仅dms的rockectMq需要该字段
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例ID，仅dms的rockectMq需要该字段
	InstanceId string `json:"instance_id"`

	// 消费组
	Group string `json:"group"`

	// topic名称
	Topic string `json:"topic"`

	// 标签
	Tag *string `json:"tag,omitempty"`

	// 开启SSL
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// ACL访问控制
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// 用户名
	AccessKey *string `json:"access_key,omitempty"`

	// 密码
	SecretKey *string `json:"secret_key,omitempty"`

	// 消费方式，针对不同生产顺序消息类型，选择消费方式会导致不同结果，请严格按照需求选择消费方式。1、生产顺序为：设置消息组，保证消息顺序发送。消费方式为：顺序消费，实际消息处理结果：按照消息组粒度，严格保证消息顺序。 同一消息组内的消息的消费顺序和发送顺序完全一致。2、生产顺序为：设置消息组，保证消息顺序发送。消费方式为：并发消费，实际消息处理结果：并发消费，尽可能按时间顺序处理。3、生产顺序为：未设置消息组，消息乱序发送。消费方式为：顺序消费，实际消息处理结果：按队列存储粒度，严格顺序。 基于 Apache RocketMQ 本身队列的属性，消费顺序和队列存储的顺序一致，但不保证和发送顺序一致。4、生产顺序为：未设置消息组，消息乱序发送。消费方式为：并发消费，实际消息处理结果：并发消费，尽可能按照时间顺序处理。
	MessageType *SourceDmsmqParametersMessageType `json:"message_type,omitempty"`

	// mq实例版本
	EngineVersion *SourceDmsmqParametersEngineVersion `json:"engine_version,omitempty"`

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

func (o SourceDmsmqParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceDmsmqParameters struct{}"
	}

	return strings.Join([]string{"SourceDmsmqParameters", string(data)}, " ")
}

type SourceDmsmqParametersMessageType struct {
	value string
}

type SourceDmsmqParametersMessageTypeEnum struct {
	NORMAL SourceDmsmqParametersMessageType
	ORDER  SourceDmsmqParametersMessageType
}

func GetSourceDmsmqParametersMessageTypeEnum() SourceDmsmqParametersMessageTypeEnum {
	return SourceDmsmqParametersMessageTypeEnum{
		NORMAL: SourceDmsmqParametersMessageType{
			value: "NORMAL",
		},
		ORDER: SourceDmsmqParametersMessageType{
			value: "ORDER",
		},
	}
}

func (c SourceDmsmqParametersMessageType) Value() string {
	return c.value
}

func (c SourceDmsmqParametersMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceDmsmqParametersMessageType) UnmarshalJSON(b []byte) error {
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

type SourceDmsmqParametersEngineVersion struct {
	value string
}

type SourceDmsmqParametersEngineVersionEnum struct {
	E_4_X SourceDmsmqParametersEngineVersion
	E_5_X SourceDmsmqParametersEngineVersion
}

func GetSourceDmsmqParametersEngineVersionEnum() SourceDmsmqParametersEngineVersionEnum {
	return SourceDmsmqParametersEngineVersionEnum{
		E_4_X: SourceDmsmqParametersEngineVersion{
			value: "4.x",
		},
		E_5_X: SourceDmsmqParametersEngineVersion{
			value: "5.x",
		},
	}
}

func (c SourceDmsmqParametersEngineVersion) Value() string {
	return c.value
}

func (c SourceDmsmqParametersEngineVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceDmsmqParametersEngineVersion) UnmarshalJSON(b []byte) error {
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
