package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubscriptionTarget struct {

	// 订阅目标ID，需保证全局唯一，由小写字母、数字、中划线组成，必须字母或数字开头。 更新订阅场景时，指定ID的订阅目标存在时则进行更新，否则进行创建； 创建订阅目标场景时，指定ID作为待创建的订阅目标对象ID，未指定时由系统自动生成。 更新订阅目标时，此字段忽略；
	Id *string `json:"id,omitempty"`

	// 订阅的事件目标名称
	Name string `json:"name"`

	// 订阅的事件目标的提供方类型
	ProviderType SubscriptionTargetProviderType `json:"provider_type"`

	// 订阅的事件目标使用的目标链接ID
	ConnectionId *string `json:"connection_id,omitempty"`

	// 订阅的事件目标参数列表，该字段序列化后总长度不超过1024字节，函数(urn、invoke_type、agency_name)、函数流（workflow_id、agency_name）、webhook（url）订阅目标必填，其中函数、函数流委托名称必填
	Detail *interface{} `json:"detail,omitempty"`

	KafkaDetail *KafkaTargetDetail `json:"kafka_detail,omitempty"`

	SmnDetail *SmnTargetDetail `json:"smn_detail,omitempty"`

	EgDetail *EgTargetDetail `json:"eg_detail,omitempty"`

	Transform *TransForm `json:"transform"`

	DeadLetterQueue *DeadLetterQueue `json:"dead_letter_queue,omitempty"`
}

func (o SubscriptionTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionTarget struct{}"
	}

	return strings.Join([]string{"SubscriptionTarget", string(data)}, " ")
}

type SubscriptionTargetProviderType struct {
	value string
}

type SubscriptionTargetProviderTypeEnum struct {
	CUSTOM   SubscriptionTargetProviderType
	OFFICIAL SubscriptionTargetProviderType
}

func GetSubscriptionTargetProviderTypeEnum() SubscriptionTargetProviderTypeEnum {
	return SubscriptionTargetProviderTypeEnum{
		CUSTOM: SubscriptionTargetProviderType{
			value: "CUSTOM",
		},
		OFFICIAL: SubscriptionTargetProviderType{
			value: "OFFICIAL",
		},
	}
}

func (c SubscriptionTargetProviderType) Value() string {
	return c.value
}

func (c SubscriptionTargetProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubscriptionTargetProviderType) UnmarshalJSON(b []byte) error {
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
