package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type CreateDnsNameRequest struct {
	// 语言

	XLanguage *CreateDnsNameRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *CreateDnsNameRequestBody `json:"body,omitempty"`
}

func (o CreateDnsNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDnsNameRequest struct{}"
	}

	return strings.Join([]string{"CreateDnsNameRequest", string(data)}, " ")
}

type CreateDnsNameRequestXLanguage struct {
	value string
}

type CreateDnsNameRequestXLanguageEnum struct {
	ZH_CN CreateDnsNameRequestXLanguage
	EN_US CreateDnsNameRequestXLanguage
}

func GetCreateDnsNameRequestXLanguageEnum() CreateDnsNameRequestXLanguageEnum {
	return CreateDnsNameRequestXLanguageEnum{
		ZH_CN: CreateDnsNameRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateDnsNameRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateDnsNameRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDnsNameRequestXLanguage) UnmarshalJSON(b []byte) error {
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
