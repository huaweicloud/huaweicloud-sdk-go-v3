package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowProjectQuotasRequest Request Object
type ShowProjectQuotasRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// '功能说明：根据type过滤查询指定类型的配额' 取值范围：instance
	Type *ShowProjectQuotasRequestType `json:"type,omitempty"`
}

func (o ShowProjectQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectQuotasRequest", string(data)}, " ")
}

type ShowProjectQuotasRequestType struct {
	value string
}

type ShowProjectQuotasRequestTypeEnum struct {
	INSTANCE ShowProjectQuotasRequestType
}

func GetShowProjectQuotasRequestTypeEnum() ShowProjectQuotasRequestTypeEnum {
	return ShowProjectQuotasRequestTypeEnum{
		INSTANCE: ShowProjectQuotasRequestType{
			value: "instance",
		},
	}
}

func (c ShowProjectQuotasRequestType) Value() string {
	return c.value
}

func (c ShowProjectQuotasRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProjectQuotasRequestType) UnmarshalJSON(b []byte) error {
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
