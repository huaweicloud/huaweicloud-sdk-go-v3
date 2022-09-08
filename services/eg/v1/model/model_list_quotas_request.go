package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
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
