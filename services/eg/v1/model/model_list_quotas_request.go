package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListQuotasRequest Request Object
type ListQuotasRequest struct {

	// 指定查询资源类型的配额
	Type *ListQuotasRequestType `json:"type,omitempty"`
}

func (o ListQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListQuotasRequest", string(data)}, " ")
}

type ListQuotasRequestType struct {
	value string
}

type ListQuotasRequestTypeEnum struct {
	CHANNEL              ListQuotasRequestType
	CHANNEL_SUBSCRIPTION ListQuotasRequestType
	SOURCE               ListQuotasRequestType
	SUBSCRIPTION         ListQuotasRequestType
	SUBSCRIPTION_TARGET  ListQuotasRequestType
	SCHEMA               ListQuotasRequestType
	SCHEMA_VERSION       ListQuotasRequestType
	CONNECTION           ListQuotasRequestType
	PRIVATE_ENDPOINT     ListQuotasRequestType
	SOURCE_RABBITMQ      ListQuotasRequestType
	SOURCE_ROCKETMQ      ListQuotasRequestType
}

func GetListQuotasRequestTypeEnum() ListQuotasRequestTypeEnum {
	return ListQuotasRequestTypeEnum{
		CHANNEL: ListQuotasRequestType{
			value: "CHANNEL",
		},
		CHANNEL_SUBSCRIPTION: ListQuotasRequestType{
			value: "CHANNEL_SUBSCRIPTION",
		},
		SOURCE: ListQuotasRequestType{
			value: "SOURCE",
		},
		SUBSCRIPTION: ListQuotasRequestType{
			value: "SUBSCRIPTION",
		},
		SUBSCRIPTION_TARGET: ListQuotasRequestType{
			value: "SUBSCRIPTION_TARGET",
		},
		SCHEMA: ListQuotasRequestType{
			value: "SCHEMA",
		},
		SCHEMA_VERSION: ListQuotasRequestType{
			value: "SCHEMA_VERSION",
		},
		CONNECTION: ListQuotasRequestType{
			value: "CONNECTION",
		},
		PRIVATE_ENDPOINT: ListQuotasRequestType{
			value: "PRIVATE_ENDPOINT",
		},
		SOURCE_RABBITMQ: ListQuotasRequestType{
			value: "SOURCE_RABBITMQ",
		},
		SOURCE_ROCKETMQ: ListQuotasRequestType{
			value: "SOURCE_ROCKETMQ",
		},
	}
}

func (c ListQuotasRequestType) Value() string {
	return c.value
}

func (c ListQuotasRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListQuotasRequestType) UnmarshalJSON(b []byte) error {
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
