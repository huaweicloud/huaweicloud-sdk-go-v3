package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type StartInstanceSlaveDrActionRequest struct {
	XLanguage  *StartInstanceSlaveDrActionRequestXLanguage `json:"X-Language,omitempty"`
	InstanceId string                                      `json:"instance_id"`
	Body       *BuildSlaveDrRelation                       `json:"body,omitempty"`
}

func (o StartInstanceSlaveDrActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceSlaveDrActionRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceSlaveDrActionRequest", string(data)}, " ")
}

type StartInstanceSlaveDrActionRequestXLanguage struct {
	value string
}

type StartInstanceSlaveDrActionRequestXLanguageEnum struct {
	ZH_CN StartInstanceSlaveDrActionRequestXLanguage
	EN_US StartInstanceSlaveDrActionRequestXLanguage
}

func GetStartInstanceSlaveDrActionRequestXLanguageEnum() StartInstanceSlaveDrActionRequestXLanguageEnum {
	return StartInstanceSlaveDrActionRequestXLanguageEnum{
		ZH_CN: StartInstanceSlaveDrActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartInstanceSlaveDrActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartInstanceSlaveDrActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *StartInstanceSlaveDrActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
