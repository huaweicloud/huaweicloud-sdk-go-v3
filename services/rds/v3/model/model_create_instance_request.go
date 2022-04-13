package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type CreateInstanceRequest struct {
	// 语言

	XLanguage *CreateInstanceRequestXLanguage `json:"X-Language,omitempty"`

	Body *InstanceRequest `json:"body,omitempty"`
}

func (o CreateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequest", string(data)}, " ")
}

type CreateInstanceRequestXLanguage struct {
	value string
}

type CreateInstanceRequestXLanguageEnum struct {
	ZH_CN CreateInstanceRequestXLanguage
	EN_US CreateInstanceRequestXLanguage
}

func GetCreateInstanceRequestXLanguageEnum() CreateInstanceRequestXLanguageEnum {
	return CreateInstanceRequestXLanguageEnum{
		ZH_CN: CreateInstanceRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateInstanceRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateInstanceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceRequestXLanguage) UnmarshalJSON(b []byte) error {
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
