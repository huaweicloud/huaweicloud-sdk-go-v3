package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourceMobileMqParameters 事件流移动云rockectMq事件源参数
type SourceMobileMqParameters struct {

	// 消费组id
	GroupId string `json:"group_id"`

	// 实例id
	InstanceId string `json:"instance_id"`

	// topic
	Topic string `json:"topic"`

	// 标签
	Tag *string `json:"tag,omitempty"`

	// 鉴权认证
	AuthenticationRequired *bool `json:"authentication_required,omitempty"`

	// 保存消息轨迹
	MsgTraceSwitch *bool `json:"msg_trace_switch,omitempty"`

	// AccessKey
	AccessKey *string `json:"access_key,omitempty"`

	// SecretKey
	SecretKey *string `json:"secret_key,omitempty"`

	// 订阅方式
	MessageModel SourceMobileMqParametersMessageModel `json:"message_model"`

	// 接入点类型
	AddrType SourceMobileMqParametersAddrType `json:"addr_type"`

	// 地址
	Addr string `json:"addr"`

	// 依赖SDK
	SdkUrl string `json:"sdk_url"`

	// 消费超时时间
	ConsumeTimeout int32 `json:"consume_timeout"`

	// 消息类型
	MessageType SourceMobileMqParametersMessageType `json:"message_type"`

	// 失败重试的等待时间
	SuspendTime *int32 `json:"suspend_time,omitempty"`

	// 最大重试次数
	MaxReconsumerTimes *int32 `json:"max_reconsumer_times,omitempty"`

	// 消费线程数
	ConsumerThreadNums *int32 `json:"consumer_thread_nums,omitempty"`

	// 批量消费最大消息数
	ConsumerBatchMaxSize *int32 `json:"consumer_batch_max_size,omitempty"`

	// 批量消费最大等待时长，单位：秒
	ConsumerMaxWait *int32 `json:"consumer_max_wait,omitempty"`

	// 虚拟私有云
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o SourceMobileMqParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceMobileMqParameters struct{}"
	}

	return strings.Join([]string{"SourceMobileMqParameters", string(data)}, " ")
}

type SourceMobileMqParametersMessageModel struct {
	value string
}

type SourceMobileMqParametersMessageModelEnum struct {
	CLUSTERING   SourceMobileMqParametersMessageModel
	BROADCASTING SourceMobileMqParametersMessageModel
}

func GetSourceMobileMqParametersMessageModelEnum() SourceMobileMqParametersMessageModelEnum {
	return SourceMobileMqParametersMessageModelEnum{
		CLUSTERING: SourceMobileMqParametersMessageModel{
			value: "CLUSTERING",
		},
		BROADCASTING: SourceMobileMqParametersMessageModel{
			value: "BROADCASTING",
		},
	}
}

func (c SourceMobileMqParametersMessageModel) Value() string {
	return c.value
}

func (c SourceMobileMqParametersMessageModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceMobileMqParametersMessageModel) UnmarshalJSON(b []byte) error {
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

type SourceMobileMqParametersAddrType struct {
	value string
}

type SourceMobileMqParametersAddrTypeEnum struct {
	PUBLIC  SourceMobileMqParametersAddrType
	PRIVATE SourceMobileMqParametersAddrType
}

func GetSourceMobileMqParametersAddrTypeEnum() SourceMobileMqParametersAddrTypeEnum {
	return SourceMobileMqParametersAddrTypeEnum{
		PUBLIC: SourceMobileMqParametersAddrType{
			value: "PUBLIC",
		},
		PRIVATE: SourceMobileMqParametersAddrType{
			value: "PRIVATE",
		},
	}
}

func (c SourceMobileMqParametersAddrType) Value() string {
	return c.value
}

func (c SourceMobileMqParametersAddrType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceMobileMqParametersAddrType) UnmarshalJSON(b []byte) error {
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

type SourceMobileMqParametersMessageType struct {
	value string
}

type SourceMobileMqParametersMessageTypeEnum struct {
	NORMAL SourceMobileMqParametersMessageType
}

func GetSourceMobileMqParametersMessageTypeEnum() SourceMobileMqParametersMessageTypeEnum {
	return SourceMobileMqParametersMessageTypeEnum{
		NORMAL: SourceMobileMqParametersMessageType{
			value: "NORMAL",
		},
	}
}

func (c SourceMobileMqParametersMessageType) Value() string {
	return c.value
}

func (c SourceMobileMqParametersMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceMobileMqParametersMessageType) UnmarshalJSON(b []byte) error {
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
