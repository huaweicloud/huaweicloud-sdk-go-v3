package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type StartInstanceRestartActionRequest struct {
	// 语言

	XLanguage *StartInstanceRestartActionRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *InstanceRestartRequsetBody `json:"body,omitempty"`
}

func (o StartInstanceRestartActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceRestartActionRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceRestartActionRequest", string(data)}, " ")
}

type StartInstanceRestartActionRequestXLanguage struct {
	value string
}

type StartInstanceRestartActionRequestXLanguageEnum struct {
	ZH_CN StartInstanceRestartActionRequestXLanguage
	EN_US StartInstanceRestartActionRequestXLanguage
}

func GetStartInstanceRestartActionRequestXLanguageEnum() StartInstanceRestartActionRequestXLanguageEnum {
	return StartInstanceRestartActionRequestXLanguageEnum{
		ZH_CN: StartInstanceRestartActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartInstanceRestartActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartInstanceRestartActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartInstanceRestartActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
