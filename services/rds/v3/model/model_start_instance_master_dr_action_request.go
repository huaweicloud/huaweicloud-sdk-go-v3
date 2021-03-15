package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type StartInstanceMasterDrActionRequest struct {
	XLanguage  *StartInstanceMasterDrActionRequestXLanguage `json:"X-Language,omitempty"`
	InstanceId string                                       `json:"instance_id"`
	Body       *BuildMasterDrRelation                       `json:"body,omitempty"`
}

func (o StartInstanceMasterDrActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceMasterDrActionRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceMasterDrActionRequest", string(data)}, " ")
}

type StartInstanceMasterDrActionRequestXLanguage struct {
	value string
}

type StartInstanceMasterDrActionRequestXLanguageEnum struct {
	ZH_CN StartInstanceMasterDrActionRequestXLanguage
	EN_US StartInstanceMasterDrActionRequestXLanguage
}

func GetStartInstanceMasterDrActionRequestXLanguageEnum() StartInstanceMasterDrActionRequestXLanguageEnum {
	return StartInstanceMasterDrActionRequestXLanguageEnum{
		ZH_CN: StartInstanceMasterDrActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartInstanceMasterDrActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartInstanceMasterDrActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *StartInstanceMasterDrActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
