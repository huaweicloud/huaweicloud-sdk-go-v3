package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowMysqlProjectQuotasRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// '功能说明：根据type过滤查询指定类型的配额' 取值范围：instance

	Type *ShowMysqlProjectQuotasRequestType `json:"type,omitempty"`
}

func (o ShowMysqlProjectQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlProjectQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowMysqlProjectQuotasRequest", string(data)}, " ")
}

type ShowMysqlProjectQuotasRequestType struct {
	value string
}

type ShowMysqlProjectQuotasRequestTypeEnum struct {
	INSTANCE ShowMysqlProjectQuotasRequestType
}

func GetShowMysqlProjectQuotasRequestTypeEnum() ShowMysqlProjectQuotasRequestTypeEnum {
	return ShowMysqlProjectQuotasRequestTypeEnum{
		INSTANCE: ShowMysqlProjectQuotasRequestType{
			value: "instance",
		},
	}
}

func (c ShowMysqlProjectQuotasRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowMysqlProjectQuotasRequestType) UnmarshalJSON(b []byte) error {
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
