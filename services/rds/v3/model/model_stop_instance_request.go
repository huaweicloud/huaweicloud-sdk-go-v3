package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type StopInstanceRequest struct {
	// 语言

	XLanguage *StopInstanceRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o StopInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopInstanceRequest", string(data)}, " ")
}

type StopInstanceRequestXLanguage struct {
	value string
}

type StopInstanceRequestXLanguageEnum struct {
	ZH_CN StopInstanceRequestXLanguage
	EN_US StopInstanceRequestXLanguage
}

func GetStopInstanceRequestXLanguageEnum() StopInstanceRequestXLanguageEnum {
	return StopInstanceRequestXLanguageEnum{
		ZH_CN: StopInstanceRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StopInstanceRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StopInstanceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopInstanceRequestXLanguage) UnmarshalJSON(b []byte) error {
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
