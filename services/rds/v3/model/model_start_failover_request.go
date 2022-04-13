package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type StartFailoverRequest struct {
	// 语言

	XLanguage *StartFailoverRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o StartFailoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartFailoverRequest struct{}"
	}

	return strings.Join([]string{"StartFailoverRequest", string(data)}, " ")
}

type StartFailoverRequestXLanguage struct {
	value string
}

type StartFailoverRequestXLanguageEnum struct {
	ZH_CN StartFailoverRequestXLanguage
	EN_US StartFailoverRequestXLanguage
}

func GetStartFailoverRequestXLanguageEnum() StartFailoverRequestXLanguageEnum {
	return StartFailoverRequestXLanguageEnum{
		ZH_CN: StartFailoverRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartFailoverRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartFailoverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartFailoverRequestXLanguage) UnmarshalJSON(b []byte) error {
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
