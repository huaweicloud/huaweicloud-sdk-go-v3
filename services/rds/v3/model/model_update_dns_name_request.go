package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type UpdateDnsNameRequest struct {
	// 语言

	XLanguage *UpdateDnsNameRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *ModifyDnsNameRequestBody `json:"body,omitempty"`
}

func (o UpdateDnsNameRequest) String() string {
	data, err := utils.Marshal(o)
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
	return utils.Marshal(c.value)
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
