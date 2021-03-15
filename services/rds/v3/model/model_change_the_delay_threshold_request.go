package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ChangeTheDelayThresholdRequest struct {
	XLanguage  *ChangeTheDelayThresholdRequestXLanguage `json:"X-Language,omitempty"`
	InstanceId string                                   `json:"instance_id"`
	Body       *ChangingTheDelayThresholdRequestBody    `json:"body,omitempty"`
}

func (o ChangeTheDelayThresholdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeTheDelayThresholdRequest struct{}"
	}

	return strings.Join([]string{"ChangeTheDelayThresholdRequest", string(data)}, " ")
}

type ChangeTheDelayThresholdRequestXLanguage struct {
	value string
}

type ChangeTheDelayThresholdRequestXLanguageEnum struct {
	ZH_CN ChangeTheDelayThresholdRequestXLanguage
	EN_US ChangeTheDelayThresholdRequestXLanguage
}

func GetChangeTheDelayThresholdRequestXLanguageEnum() ChangeTheDelayThresholdRequestXLanguageEnum {
	return ChangeTheDelayThresholdRequestXLanguageEnum{
		ZH_CN: ChangeTheDelayThresholdRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeTheDelayThresholdRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeTheDelayThresholdRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ChangeTheDelayThresholdRequestXLanguage) UnmarshalJSON(b []byte) error {
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
