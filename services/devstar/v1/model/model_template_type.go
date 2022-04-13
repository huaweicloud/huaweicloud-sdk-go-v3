package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type TemplateType struct {
	value string
}

type TemplateTypeEnum struct {
	MOBILE       TemplateType
	MICROSERVICE TemplateType
	WEB          TemplateType
	FUNCTION     TemplateType
	IOT          TemplateType
	AI           TemplateType
	OTHERS       TemplateType
	NONE         TemplateType
}

func GetTemplateTypeEnum() TemplateTypeEnum {
	return TemplateTypeEnum{
		MOBILE: TemplateType{
			value: "mobile",
		},
		MICROSERVICE: TemplateType{
			value: "microservice",
		},
		WEB: TemplateType{
			value: "web",
		},
		FUNCTION: TemplateType{
			value: "function",
		},
		IOT: TemplateType{
			value: "iot",
		},
		AI: TemplateType{
			value: "ai",
		},
		OTHERS: TemplateType{
			value: "others",
		},
		NONE: TemplateType{
			value: "none",
		},
	}
}

func (c TemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateType) UnmarshalJSON(b []byte) error {
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
