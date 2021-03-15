package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type StartInstanceDrToMasterActionRequest struct {
	XLanguage  *StartInstanceDrToMasterActionRequestXLanguage `json:"X-Language,omitempty"`
	InstanceId string                                         `json:"instance_id"`
	Body       *DrReplicaToMaster                             `json:"body,omitempty"`
}

func (o StartInstanceDrToMasterActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceDrToMasterActionRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceDrToMasterActionRequest", string(data)}, " ")
}

type StartInstanceDrToMasterActionRequestXLanguage struct {
	value string
}

type StartInstanceDrToMasterActionRequestXLanguageEnum struct {
	ZH_CN StartInstanceDrToMasterActionRequestXLanguage
	EN_US StartInstanceDrToMasterActionRequestXLanguage
}

func GetStartInstanceDrToMasterActionRequestXLanguageEnum() StartInstanceDrToMasterActionRequestXLanguageEnum {
	return StartInstanceDrToMasterActionRequestXLanguageEnum{
		ZH_CN: StartInstanceDrToMasterActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartInstanceDrToMasterActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartInstanceDrToMasterActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *StartInstanceDrToMasterActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
