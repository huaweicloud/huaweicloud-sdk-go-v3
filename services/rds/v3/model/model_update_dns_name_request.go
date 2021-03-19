package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateDnsNameRequest struct {
	XLanguage *UpdateDnsNameRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *ModifyDnsNameRequestBody `json:"body,omitempty"`
}

func (o UpdateDnsNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDnsNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateDnsNameRequest", string(data)}, " ")
}

type UpdateDnsNameRequestXLanguage struct {
	value string
}

type UpdateDnsNameRequestXLanguageEnum struct {
	ZH_CN UpdateDnsNameRequestXLanguage
	EN_US UpdateDnsNameRequestXLanguage
}

func GetUpdateDnsNameRequestXLanguageEnum() UpdateDnsNameRequestXLanguageEnum {
	return UpdateDnsNameRequestXLanguageEnum{
		ZH_CN: UpdateDnsNameRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateDnsNameRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateDnsNameRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateDnsNameRequestXLanguage) UnmarshalJSON(b []byte) error {
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
