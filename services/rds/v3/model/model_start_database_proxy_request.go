package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type StartDatabaseProxyRequest struct {
	// 语言

	XLanguage *StartDatabaseProxyRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *OpenProxyRequest `json:"body,omitempty"`
}

func (o StartDatabaseProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDatabaseProxyRequest struct{}"
	}

	return strings.Join([]string{"StartDatabaseProxyRequest", string(data)}, " ")
}

type StartDatabaseProxyRequestXLanguage struct {
	value string
}

type StartDatabaseProxyRequestXLanguageEnum struct {
	ZH_CN StartDatabaseProxyRequestXLanguage
	EN_US StartDatabaseProxyRequestXLanguage
}

func GetStartDatabaseProxyRequestXLanguageEnum() StartDatabaseProxyRequestXLanguageEnum {
	return StartDatabaseProxyRequestXLanguageEnum{
		ZH_CN: StartDatabaseProxyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartDatabaseProxyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartDatabaseProxyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartDatabaseProxyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
