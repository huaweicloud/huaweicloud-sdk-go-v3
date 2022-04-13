package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type StartupInstanceRequest struct {
	// 语言

	XLanguage *StartupInstanceRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o StartupInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartupInstanceRequest struct{}"
	}

	return strings.Join([]string{"StartupInstanceRequest", string(data)}, " ")
}

type StartupInstanceRequestXLanguage struct {
	value string
}

type StartupInstanceRequestXLanguageEnum struct {
	ZH_CN StartupInstanceRequestXLanguage
	EN_US StartupInstanceRequestXLanguage
}

func GetStartupInstanceRequestXLanguageEnum() StartupInstanceRequestXLanguageEnum {
	return StartupInstanceRequestXLanguageEnum{
		ZH_CN: StartupInstanceRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartupInstanceRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartupInstanceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartupInstanceRequestXLanguage) UnmarshalJSON(b []byte) error {
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
